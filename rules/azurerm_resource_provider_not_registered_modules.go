package rules

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-ruleset-azurerm/rules/resourceproviders"
	"github.com/zclconf/go-cty/cty"
)

// maxModuleDepth guards against cycles between local module sources, which Terraform rejects anyway.
const maxModuleDepth = 32

var (
	childModuleSchema = &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "resource", LabelNames: []string{"type", "name"}},
			{Type: "module", LabelNames: []string{"name"}},
			{Type: "provider", LabelNames: []string{"name"}},
		},
	}
	childResourceSchema = &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{{Name: "provider"}},
	}
	moduleCallSchema = &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{{Name: "source"}, {Name: "providers"}},
	}
)

// moduleCall is a module block, with the provider aliases of the calling module it passes to the called module.
type moduleCall struct {
	name   string
	source string
	// providers maps the azurerm aliases of the called module to those of the calling module. It is nil when the
	// providers argument cannot be interpreted.
	providers map[string]string
}

// moduleResource is a resource declared in a (nested) child module.
type moduleResource struct {
	resourceType string
	// address is the module address the resource is declared in, e.g. module.base.module.network.
	address string
	// alias is the azurerm alias of the root module configuration the resource uses.
	alias string
}

// moduleLoader reads the resources of child modules from disk, the way Terraform resolves them.
type moduleLoader struct {
	rootDir string
	// manifest maps module keys, e.g. "base.network", to their directory relative to rootDir.
	manifest map[string]string
	parser   *hclparse.Parser
}

func newModuleLoader(rootDir string) (*moduleLoader, error) {
	manifest, err := loadModuleManifest(rootDir)
	if err != nil {
		return nil, err
	}
	return &moduleLoader{rootDir: rootDir, manifest: manifest, parser: hclparse.NewParser()}, nil
}

// loadModuleManifest reads the modules.json file written by terraform init. It returns an empty manifest when the
// configuration has not been initialized, in which case only local module sources can be resolved.
func loadModuleManifest(rootDir string) (map[string]string, error) {
	dataDir := os.Getenv("TF_DATA_DIR")
	if dataDir == "" {
		dataDir = ".terraform"
	}
	if !filepath.IsAbs(dataDir) {
		dataDir = filepath.Join(rootDir, dataDir)
	}

	src, err := os.ReadFile(filepath.Join(dataDir, "modules", "modules.json"))
	if errors.Is(err, fs.ErrNotExist) {
		return map[string]string{}, nil
	}
	if err != nil {
		return nil, err
	}

	var content struct {
		Modules []struct {
			Key string `json:"Key"`
			Dir string `json:"Dir"`
		} `json:"Modules"`
	}
	if err := json.Unmarshal(src, &content); err != nil {
		return nil, fmt.Errorf("failed to parse module manifest: %w", err)
	}

	manifest := make(map[string]string, len(content.Modules))
	for _, module := range content.Modules {
		manifest[module.Key] = module.Dir
	}
	return manifest, nil
}

// resources returns the azurerm resources declared in the module called by call from the module with the given key
// and directory, including nested modules. rootAliases maps the azurerm aliases of the calling module to those of the
// root module.
func (l *moduleLoader) resources(parentKey, parentDir, parentAddress string, call moduleCall, rootAliases map[string]string, depth int) []moduleResource {
	key := call.name
	if parentKey != "" {
		key = parentKey + "." + call.name
	}
	address := parentAddress + "module." + call.name

	if depth > maxModuleDepth {
		logger.Debug("skipping module, maximum depth exceeded", "module", address)
		return nil
	}
	if call.providers == nil {
		logger.Debug("skipping module, providers argument cannot be interpreted", "module", address)
		return nil
	}

	dir, ok := l.manifest[key]
	if !ok {
		if !isLocalModuleSource(call.source) {
			logger.Debug("skipping module, not installed", "module", address, "source", call.source)
			return nil
		}
		dir = filepath.Join(parentDir, call.source)
	}

	path := dir
	if !filepath.IsAbs(path) {
		path = filepath.Join(l.rootDir, dir)
	}
	content, err := l.parseModule(path)
	if err != nil {
		logger.Warn("skipping module, failed to parse", "module", address, "error", err)
		return nil
	}

	// A module with its own provider configuration does not use the configuration of the root module.
	for _, block := range content.Blocks {
		if block.Type == "provider" && block.Labels[0] == "azurerm" {
			logger.Debug("skipping module, it declares an azurerm provider configuration", "module", address)
			return nil
		}
	}

	aliases := map[string]string{}
	for childAlias, parentAlias := range call.providers {
		if rootAlias, ok := rootAliases[parentAlias]; ok {
			aliases[childAlias] = rootAlias
		}
	}

	var resources []moduleResource
	for _, block := range content.Blocks {
		switch block.Type {
		case "resource":
			resource, ok := childResource(block, address, aliases)
			if ok {
				resources = append(resources, resource)
			}
		case "module":
			nested, ok := parseModuleCall(block)
			if !ok {
				logger.Debug("skipping module, source cannot be determined", "module", address+".module."+block.Labels[0])
				continue
			}
			resources = append(resources, l.resources(key, dir, address+".", nested, aliases, depth+1)...)
		}
	}

	return resources
}

func childResource(block *hcl.Block, address string, aliases map[string]string) (moduleResource, bool) {
	resource := moduleResource{resourceType: block.Labels[0], address: address}
	if _, ok := resourceproviders.Resources[resource.resourceType]; !ok {
		return resource, false
	}

	body, _, diags := block.Body.PartialContent(childResourceSchema)
	if diags.HasErrors() {
		return resource, false
	}

	alias := ""
	if attr, ok := body.Attributes["provider"]; ok {
		if alias, ok = azurermAlias(attr.Expr); !ok {
			return resource, false
		}
	}

	rootAlias, ok := aliases[alias]
	if !ok {
		logger.Debug("skipping resource, provider configuration is not passed to the module", "resource", resource.resourceType, "module", address, "alias", alias)
		return resource, false
	}
	resource.alias = rootAlias
	return resource, true
}

// parseModule returns the blocks of all Terraform files in a module directory, sorted by file name.
func (l *moduleLoader) parseModule(dir string) (*hcl.BodyContent, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	merged := &hcl.BodyContent{}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !(strings.HasSuffix(name, ".tf") || strings.HasSuffix(name, ".tf.json")) {
			continue
		}

		var file *hcl.File
		var diags hcl.Diagnostics
		if strings.HasSuffix(name, ".json") {
			file, diags = l.parser.ParseJSONFile(filepath.Join(dir, name))
		} else {
			file, diags = l.parser.ParseHCLFile(filepath.Join(dir, name))
		}
		if diags.HasErrors() {
			return nil, diags
		}

		content, _, diags := file.Body.PartialContent(childModuleSchema)
		if diags.HasErrors() {
			return nil, diags
		}
		merged.Blocks = append(merged.Blocks, content.Blocks...)
	}

	return merged, nil
}

// parseModuleCall interprets a module block. It returns false when the source is not a static string.
func parseModuleCall(block *hcl.Block) (moduleCall, bool) {
	body, _, diags := block.Body.PartialContent(moduleCallSchema)
	if diags.HasErrors() {
		return moduleCall{}, false
	}
	var source, providers hcl.Expression
	if attr, ok := body.Attributes["source"]; ok {
		source = attr.Expr
	}
	if attr, ok := body.Attributes["providers"]; ok {
		providers = attr.Expr
	}
	return newModuleCall(block.Labels[0], source, providers)
}

// newModuleCall returns a module call from the source and providers arguments, which may be nil. It returns false
// when the source is not a static string.
func newModuleCall(name string, source, providers hcl.Expression) (moduleCall, bool) {
	call := moduleCall{name: name}
	if source == nil {
		return call, false
	}
	value, diags := source.Value(nil)
	if diags.HasErrors() || value.IsNull() || !value.IsKnown() || !value.Type().Equals(cty.String) {
		return call, false
	}
	call.source = value.AsString()
	call.providers = moduleProviders(providers)
	return call, true
}

// moduleProviders maps the azurerm aliases of a called module to those of the calling module. The default configuration
// is implicitly inherited unless the providers argument maps it explicitly. It returns nil when the argument cannot be
// interpreted.
func moduleProviders(expr hcl.Expression) map[string]string {
	providers := map[string]string{"": ""}
	if expr == nil {
		return providers
	}

	pairs, diags := hcl.ExprMap(expr)
	if diags.HasErrors() {
		return nil
	}
	for _, pair := range pairs {
		childAlias, ok := azurermAlias(pair.Key)
		if !ok {
			continue
		}
		parentAlias, ok := azurermAlias(pair.Value)
		if !ok {
			continue
		}
		providers[childAlias] = parentAlias
	}
	return providers
}

// azurermAlias returns the alias in a provider reference like azurerm or azurerm.alias, with "" for the default
// configuration. It returns false when the reference is not to an azurerm provider configuration. Quoted references like
// "azurerm.alias", deprecated but still accepted by Terraform, are supported too.
func azurermAlias(expr hcl.Expression) (string, bool) {
	traversal, diags := hcl.AbsTraversalForExpr(expr)
	if diags.HasErrors() {
		value, valueDiags := expr.Value(nil)
		if valueDiags.HasErrors() || value.IsNull() || !value.IsKnown() || !value.Type().Equals(cty.String) {
			return "", false
		}
		traversal, diags = hclsyntax.ParseTraversalAbs([]byte(value.AsString()), "", hcl.InitialPos)
	}
	if diags.HasErrors() || traversal.RootName() != "azurerm" {
		return "", false
	}
	switch len(traversal) {
	case 1:
		return "", true
	case 2:
		if step, ok := traversal[1].(hcl.TraverseAttr); ok {
			return step.Name, true
		}
	}
	return "", false
}

// isLocalModuleSource reports whether a module source is a local path, which Terraform does not install.
func isLocalModuleSource(source string) bool {
	return slices.ContainsFunc([]string{"./", "../", ".\\", "..\\"}, func(prefix string) bool {
		return strings.HasPrefix(source, prefix)
	})
}
