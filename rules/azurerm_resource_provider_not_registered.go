package rules

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
	"github.com/terraform-linters/tflint-ruleset-azurerm/rules/resourceproviders"
	"github.com/zclconf/go-cty/cty"
)

const (
	resourceProviderRegistrationsAttributeName = "resource_provider_registrations"
	resourceProvidersToRegisterAttributeName   = "resource_providers_to_register"
	defaultResourceProviderRegistrations       = "none"
)

// AzurermResourceProviderNotRegisteredRule checks whether the resource providers used by resources are registered
type AzurermResourceProviderNotRegisteredRule struct {
	tflint.DefaultRule

	// rootDir is the directory of the root module. TFLint starts plugins in this directory.
	rootDir string
}

type azurermResourceProviderNotRegisteredRuleConfig struct {
	DefaultRegistrations string   `hclext:"default_registrations,optional"`
	IgnoredNamespaces    []string `hclext:"ignored_namespaces,optional"`
}

// providerRegistrations describes the namespaces an azurerm provider configuration registers.
type providerRegistrations struct {
	name string
	mode string
	// registered contains the registered namespaces with their exact casing, as the azurerm provider skips
	// namespaces whose casing differs from the Azure API. It is nil when they cannot be determined statically.
	registered map[string]struct{}
	// miscased contains the resource_providers_to_register entries with the wrong casing, by canonical namespace.
	miscased map[string]miscasedNamespace
}

// miscasedNamespace is a resource_providers_to_register entry whose casing differs from the Azure API.
type miscasedNamespace struct {
	value string
	rng   hcl.Range
}

// registers reports whether the provider configuration registers the canonical namespace. Miscased entries count
// as registered, as they are reported on the provider configuration instead.
func (p *providerRegistrations) registers(namespace string) bool {
	if _, ok := p.registered[namespace]; ok {
		return true
	}
	_, ok := p.miscased[namespace]
	return ok
}

// NewAzurermResourceProviderNotRegisteredRule returns a new rule
func NewAzurermResourceProviderNotRegisteredRule() *AzurermResourceProviderNotRegisteredRule {
	return &AzurermResourceProviderNotRegisteredRule{rootDir: "."}
}

// Name returns the rule name
func (r *AzurermResourceProviderNotRegisteredRule) Name() string {
	return "azurerm_resource_provider_not_registered"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermResourceProviderNotRegisteredRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *AzurermResourceProviderNotRegisteredRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *AzurermResourceProviderNotRegisteredRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether the resource providers used by azurerm resources are registered by their provider configuration
func (r *AzurermResourceProviderNotRegisteredRule) Check(runner tflint.Runner) error {
	modulePath, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	// Child modules are checked from the root module, which owns the provider configurations.
	if !modulePath.IsRoot() {
		return nil
	}

	config := azurermResourceProviderNotRegisteredRuleConfig{}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}
	if config.DefaultRegistrations == "" {
		config.DefaultRegistrations = defaultResourceProviderRegistrations
	}
	if _, ok := resourceproviders.Registrations[config.DefaultRegistrations]; !ok {
		return fmt.Errorf(
			"invalid default_registrations %q for rule %s, must be one of: %s",
			config.DefaultRegistrations, r.Name(), strings.Join(slices.Sorted(maps.Keys(resourceproviders.Registrations)), ", "),
		)
	}

	providers, err := loadProviderRegistrations(runner, config.DefaultRegistrations)
	if err != nil {
		return err
	}
	// Without an azurerm provider block, this is most likely a reusable module.
	if providers == nil {
		return nil
	}

	for _, alias := range slices.Sorted(maps.Keys(providers)) {
		provider := providers[alias]
		for _, namespace := range slices.Sorted(maps.Keys(provider.miscased)) {
			entry := provider.miscased[namespace]
			if err := runner.EmitIssue(r, miscasedMessage(provider, entry.value, namespace), entry.rng); err != nil {
				return err
			}
		}
	}

	ignored := map[string]struct{}{}
	for _, namespace := range config.IgnoredNamespaces {
		ignored[strings.ToLower(namespace)] = struct{}{}
	}
	check := &resourceCheck{rule: r, runner: runner, providers: providers, ignored: ignored}

	// Resources and modules with count = 0 or an empty for_each are still checked, as they are expected to be
	// created eventually.
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: "provider"}},
				},
			},
			{
				Type:       "module",
				LabelNames: []string{"name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: "source"}, {Name: "providers"}},
				},
			},
		},
	}, &tflint.GetModuleContentOption{ExpandMode: tflint.ExpandModeNone})
	if err != nil {
		return err
	}

	var loader *moduleLoader
	rootAliases := make(map[string]string, len(providers))
	for alias := range providers {
		rootAliases[alias] = alias
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "resource":
			alias, ok := resourceAlias(block)
			if !ok {
				continue
			}
			resource := moduleResource{resourceType: block.Labels[0], alias: alias}
			if err := check.run(resource, block.DefRange); err != nil {
				return err
			}

		case "module":
			var source, providersExpr hcl.Expression
			if attr, ok := block.Body.Attributes["source"]; ok {
				source = attr.Expr
			}
			if attr, ok := block.Body.Attributes["providers"]; ok {
				providersExpr = attr.Expr
			}
			call, ok := newModuleCall(block.Labels[0], source, providersExpr)
			if !ok {
				logger.Debug("skipping module, source cannot be determined", "module", block.Labels[0])
				continue
			}

			if loader == nil {
				if loader, err = newModuleLoader(r.rootDir); err != nil {
					return err
				}
			}
			reported := map[moduleResource]bool{}
			for _, resource := range loader.resources("", ".", "", call, rootAliases, 0) {
				if reported[resource] {
					continue
				}
				reported[resource] = true
				if err := check.run(resource, block.DefRange); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// resourceCheck checks resources against the provider configurations of the root module.
type resourceCheck struct {
	rule      *AzurermResourceProviderNotRegisteredRule
	runner    tflint.Runner
	providers map[string]*providerRegistrations
	// ignored contains the lowercase namespaces that are registered outside Terraform.
	ignored map[string]struct{}
}

// run emits an issue at rng when a namespace required by the resource is not registered by its provider.
func (c *resourceCheck) run(resource moduleResource, rng hcl.Range) error {
	required, ok := resourceproviders.Resources[resource.resourceType]
	if !ok {
		return nil
	}

	provider, ok := c.providers[resource.alias]
	if !ok {
		logger.Debug("skipping resource, provider configuration not found", "resource", resource.resourceType, "alias", resource.alias)
		return nil
	}
	if provider.registered == nil {
		logger.Debug("skipping resource, registered namespaces are unknown", "resource", resource.resourceType, "provider", provider.name)
		return nil
	}

	var missing []string
	for _, namespace := range required {
		if canonical, ok := resourceproviders.Canonical(namespace); ok {
			namespace = canonical
		}
		if _, ok := c.ignored[strings.ToLower(namespace)]; ok {
			continue
		}
		if slices.Contains(resourceproviders.RegisteredByDefault, namespace) || provider.registers(namespace) {
			continue
		}
		missing = append(missing, namespace)
	}
	if len(missing) == 0 {
		return nil
	}

	return c.runner.EmitIssue(c.rule, notRegisteredMessage(resource, provider, missing), rng)
}

// loadProviderRegistrations returns the registrations per provider alias, with "" for the default configuration.
// It returns nil when the module has no azurerm provider configuration.
func loadProviderRegistrations(runner tflint.Runner, defaultMode string) (map[string]*providerRegistrations, error) {
	content, err := runner.GetProviderContent("azurerm", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "alias"},
			{Name: resourceProviderRegistrationsAttributeName},
			{Name: resourceProvidersToRegisterAttributeName},
		},
	}, nil)
	if err != nil {
		return nil, err
	}
	if len(content.Blocks) == 0 {
		return nil, nil
	}

	// Resources without a provider argument use the default configuration, which is implicit when only aliased
	// configurations are declared.
	providers := map[string]*providerRegistrations{
		"": {name: "azurerm", mode: defaultMode, registered: registeredNamespaces(defaultMode, nil)},
	}

	for _, block := range content.Blocks {
		provider, alias, err := loadProviderRegistration(runner, block, defaultMode)
		if err != nil {
			return nil, err
		}
		providers[alias] = provider
	}

	return providers, nil
}

// loadProviderRegistration returns the registrations of a provider block and its alias.
func loadProviderRegistration(runner tflint.Runner, block *hclext.Block, defaultMode string) (*providerRegistrations, string, error) {
	alias := ""
	if attr, ok := block.Body.Attributes["alias"]; ok {
		if err := runner.EvaluateExpr(attr.Expr, &alias, nil); err != nil {
			return nil, "", err
		}
	}

	provider := &providerRegistrations{name: "azurerm", mode: defaultMode}
	if alias != "" {
		provider.name = "azurerm." + alias
	}

	if attr, ok := block.Body.Attributes[resourceProviderRegistrationsAttributeName]; ok {
		value, err := evaluate(runner, attr.Expr, cty.String)
		if err != nil {
			return nil, "", err
		}
		if !value.IsWhollyKnown() {
			return provider, alias, nil
		}
		if !value.IsNull() {
			provider.mode = value.AsString()
		}
	}

	var extra []string
	if attr, ok := block.Body.Attributes[resourceProvidersToRegisterAttributeName]; ok {
		value, err := evaluate(runner, attr.Expr, cty.List(cty.String))
		if err != nil {
			return nil, "", err
		}
		if !value.IsWhollyKnown() {
			return provider, alias, nil
		}
		if !value.IsNull() {
			for it := value.ElementIterator(); it.Next(); {
				_, namespace := it.Element()
				extra = append(extra, namespace.AsString())
			}
		}
		provider.miscased = miscasedNamespaces(extra, attr.Expr.Range())
	}

	provider.registered = registeredNamespaces(provider.mode, extra)
	return provider, alias, nil
}

// registeredNamespaces returns the namespaces the azurerm provider registers for the given mode and additional
// namespaces, with their exact casing. It returns nil for an unknown mode, which the azurerm provider reports itself.
func registeredNamespaces(mode string, extra []string) map[string]struct{} {
	namespaces, ok := resourceproviders.Registrations[mode]
	if !ok {
		return nil
	}
	result := make(map[string]struct{}, len(namespaces)+len(extra))
	for _, namespace := range slices.Concat(namespaces, extra) {
		result[namespace] = struct{}{}
	}
	return result
}

// miscasedNamespaces returns the known namespaces whose casing differs from the Azure API, by canonical namespace.
func miscasedNamespaces(namespaces []string, rng hcl.Range) map[string]miscasedNamespace {
	result := map[string]miscasedNamespace{}
	for _, namespace := range namespaces {
		if canonical, ok := resourceproviders.Canonical(namespace); ok && canonical != namespace {
			result[canonical] = miscasedNamespace{value: namespace, rng: rng}
		}
	}
	return result
}

// evaluate evaluates an expression to a value of the given type, which may be null or unknown, e.g. when it refers
// to a module output or a variable without a value.
func evaluate(runner tflint.Runner, expr hcl.Expression, wantType cty.Type) (cty.Value, error) {
	var value cty.Value
	err := runner.EvaluateExpr(expr, &value, &tflint.EvaluateExprOption{WantType: &wantType})
	return value, err
}

// resourceAlias returns the azurerm provider alias used by a resource, with "" for the default configuration.
// It returns false when the resource uses a provider configuration with another local name.
func resourceAlias(resource *hclext.Block) (string, bool) {
	attr, ok := resource.Body.Attributes["provider"]
	if !ok {
		return "", true
	}
	return azurermAlias(attr.Expr)
}

func notRegisteredMessage(resource moduleResource, provider *providerRegistrations, missing []string) string {
	quoted := make([]string, len(missing))
	for i, namespace := range missing {
		quoted[i] = fmt.Sprintf("%q", namespace)
	}
	slices.Sort(quoted)

	subject := fmt.Sprintf("The %s resource provider required by %s is", quoted[0], resourceDescription(resource))
	pronoun := "it"
	if len(quoted) > 1 {
		subject = fmt.Sprintf("The %s resource providers required by %s are", strings.Join(quoted, ", "), resourceDescription(resource))
		pronoun = "them"
	}

	return fmt.Sprintf(
		"%s not registered by provider %q (resource_provider_registrations = %q). Add %s to resource_providers_to_register.",
		subject, provider.name, provider.mode, pronoun,
	)
}

func miscasedMessage(provider *providerRegistrations, value, canonical string) string {
	return fmt.Sprintf(
		"The %q resource provider in resource_providers_to_register of provider %q does not match the casing Azure uses, so the azurerm provider skips it. Use %q instead.",
		value, provider.name, canonical,
	)
}

func resourceDescription(resource moduleResource) string {
	if resource.address == "" {
		return fmt.Sprintf("%q", resource.resourceType)
	}
	return fmt.Sprintf("%q in %s", resource.resourceType, resource.address)
}
