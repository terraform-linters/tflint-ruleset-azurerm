package rules

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
	"github.com/terraform-linters/tflint-ruleset-azurerm/rules/tags"
	"github.com/zclconf/go-cty/cty"
)

// AzurermResourceMissingTagsRule checks whether resources are tagged correctly
type AzurermResourceMissingTagsRule struct {
	tflint.DefaultRule
}

type azurermResourceTagsRuleConfig struct {
	Tags    []string `hclext:"tags"`
	Exclude []string `hclext:"exclude,optional"`
}

const (
	tagsAttributeName = "tags"
)

// NewAzurermResourceMissingTagsRule returns new rules for all resources that support tags
func NewAzurermResourceMissingTagsRule() *AzurermResourceMissingTagsRule {
	return &AzurermResourceMissingTagsRule{}
}

// Name returns the rule name
func (r *AzurermResourceMissingTagsRule) Name() string {
	return "azurerm_resource_missing_tags"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermResourceMissingTagsRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *AzurermResourceMissingTagsRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

// Link returns the rule reference link
func (r *AzurermResourceMissingTagsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks resources for missing tags
func (r *AzurermResourceMissingTagsRule) Check(runner tflint.Runner) error {
	config := azurermResourceTagsRuleConfig{}

	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	for _, resourceType := range tags.Resources {
		// Skip this resource if its type is excluded in configuration
		if stringInSlice(resourceType, config.Exclude) {
			logger.Debug("excluding", "resourceType", resourceType)
			continue
		}

		resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
			Attributes: []hclext.AttributeSchema{{Name: tagsAttributeName}},
		}, nil)

		if err != nil {
			return err
		}

		for _, resource := range resources.Blocks {

			if attribute, ok := resource.Body.Attributes[tagsAttributeName]; ok {
				logger.Debug("checking", "resource type", resource.Labels[0], "resource name", resource.Labels[1], "attribute", tagsAttributeName)
				wantType := cty.Map(cty.String)
				err := runner.EvaluateExpr(attribute.Expr, func(resourceTags map[string]string) error {
					return r.emitIssue(runner, resourceTags, config, attribute.Expr.Range())
				}, &tflint.EvaluateExprOption{WantType: &wantType})
				if err != nil {
					return err
				}
			} else {
				logger.Debug("checking", "resource type", resource.Labels[0], "resource name", resource.Labels[1])
				if err := r.emitIssue(runner, map[string]string{}, config, resource.DefRange); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *AzurermResourceMissingTagsRule) emitIssue(runner tflint.Runner, tags map[string]string, config azurermResourceTagsRuleConfig, location hcl.Range) error {
	var missing []string
	for _, tag := range config.Tags {
		if _, ok := tags[tag]; !ok {
			missing = append(missing, fmt.Sprintf("\"%s\"", tag))
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		wanted := strings.Join(missing, ", ")
		issue := fmt.Sprintf("The resource is missing the following tags: %s.", wanted)
		return runner.EmitIssue(r, issue, location)
	}

	return nil
}
