package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty/gocty"
)

type mappingFile struct {
	Mappings []mapping `hcl:"mapping,block"`
}

type mapping struct {
	Resource   string         `hcl:"resource,label"`
	ImportPath string         `hcl:"import_path"`
	Attrs      hcl.Attributes `hcl:",remain"`
}

type apiSpec struct {
	definitions map[string]interface{}
	parameters  map[string]interface{}
}

type attributeRef struct {
	resource  string
	block     *string
	attribute string
	value     hcl.Expression
}

func (r *attributeRef) String() string {
	if r.block != nil {
		return fmt.Sprintf("%s.%s.%s", r.resource, *r.block, r.attribute)
	}
	return fmt.Sprintf("%s.%s", r.resource, r.attribute)
}

func (r *attributeRef) RuleName() string {
	if r.block != nil {
		return fmt.Sprintf("%s_%s_invalid_%s", r.resource, *r.block, r.attribute)
	}
	return fmt.Sprintf("%s_invalid_%s", r.resource, r.attribute)
}

func (r *attributeRef) RuleTemplate() string {
	if r.block != nil {
		return getFullPath("rule_block.go.tmpl")
	}
	return getFullPath("rule.go.tmpl")
}

type ruleMeta struct {
	RuleName      string
	RuleNameCC    string
	ResourceType  string
	BlockType     string
	AttributeName string
	Sensitive     bool
	Max           int
	SetMax        bool
	Min           int
	SetMin        bool
	Pattern       string
	Enum          []string
	ReferenceURL  string
}

type providerMeta struct {
	RuleNameCCList []string
}

type ruleDocIndexMeta struct {
	RuleNameList []string
}

var BasePath string
var RulesPath string
var DocsPath string

func getFullPath(path string) string {
	return fmt.Sprintf("%s/%s", BasePath, path)
}

func parseFlags() {
	flag.StringVar(&BasePath, "base-path", "apispec-rule-gen", "a string var")
	flag.StringVar(&RulesPath, "rules-path", "../rules", "a string var")
	flag.StringVar(&DocsPath, "docs-path", "../docs", "a string var")
	flag.Parse()
}

var terraformSchema provider
var generatedRuleNames []string = []string{}
var generatedRuleNameCCs []string = []string{}

func main() {
	parseFlags()
	terraformSchema = loadProviderSchema()

	files, err := filepath.Glob(getFullPath("mappings/*.hcl"))
	if err != nil {
		panic(err)
	}

	mappingFiles := make([]mappingFile, len(files))
	for idx, file := range files {
		parser := hclparse.NewParser()
		f, diags := parser.ParseHCLFile(file)
		if diags.HasErrors() {
			panic(diags)
		}

		var mf mappingFile
		diags = gohcl.DecodeBody(f.Body, nil, &mf)
		if diags.HasErrors() {
			panic(diags)
		}
		mappingFiles[idx] = mf
	}

	for _, mappingFile := range mappingFiles {
		for _, mapping := range mappingFile.Mappings {
			raw, err := ioutil.ReadFile(fmt.Sprintf(getFullPath("%s"), mapping.ImportPath))
			if err != nil {
				panic(err)
			}

			var spec map[string]interface{}
			err = json.Unmarshal(raw, &spec)
			if err != nil {
				panic(err)
			}
			definitions := spec["definitions"].(map[string]interface{})
			parameters := map[string]interface{}{}
			if params, exists := spec["parameters"]; exists {
				parameters = params.(map[string]interface{})
			}
			apiSpec := apiSpec{definitions: definitions, parameters: parameters}

			for attribute, value := range mapping.Attrs {
				processAttribute(apiSpec, mapping, attributeRef{resource: mapping.Resource, attribute: attribute, value: value.Expr})
			}
		}
	}

	sort.Strings(generatedRuleNameCCs)
	generateProviderFile(generatedRuleNameCCs)
	sort.Strings(generatedRuleNames)
	generateRulesIndexDoc(generatedRuleNames)
}

func processAttribute(apiSpec apiSpec, mapping mapping, ref attributeRef) {
	switch expr := ref.value.(type) {
	// attribute = Foo.Bar
	case *hclsyntax.ScopeTraversalExpr:
		fmt.Printf("Generating rule for `%s`\n", ref.String())

		// expr always consist of variables
		variable := expr.Variables()[0]
		props := make([]string, len(variable))
		for idx, prop := range variable {
			switch p := prop.(type) {
			case hcl.TraverseRoot:
				props[idx] = p.Name
			case hcl.TraverseAttr:
				props[idx] = p.Name
			}
		}

		var definition map[string]interface{}
		if len(props) == 1 {
			// attribute = Foo ("definitions" or "parameters" fields)
			if props[0] == "any" {
				return
			}
			if apiSpec.definitions[props[0]] != nil {
				definition = apiSpec.definitions[props[0]].(map[string]interface{})
			} else {
				definition = apiSpec.parameters[props[0]].(map[string]interface{})
			}
		} else {
			// attribute = Foo.Bar ("properties" fields)
			definition = apiSpec.definitions[props[0]].(map[string]interface{})["properties"].(map[string]interface{})[props[1]].(map[string]interface{})
		}

		if validMapping(definition) {
			attrSchema := extractAttrSchema(ref, definition)
			meta := generateRuleFile(mapping, ref, definition, attrSchema)
			generatedRuleNames = append(generatedRuleNames, meta.RuleName)
			generatedRuleNameCCs = append(generatedRuleNameCCs, meta.RuleNameCC)
		}
	// attribute = {
	//   child = Foo.Bar
	// }
	case *hclsyntax.ObjectConsExpr:
		for _, item := range expr.Items {
			childRef := attributeRef{resource: ref.resource, block: &ref.attribute, value: item.ValueExpr}

			val, diags := item.KeyExpr.Value(nil)
			if diags.HasErrors() {
				panic(diags)
			}

			err := gocty.FromCtyValue(val, &childRef.attribute)
			if err != nil {
				panic(err)
			}

			processAttribute(apiSpec, mapping, childRef)
		}
	}
}

func validMapping(definition map[string]interface{}) bool {
	switch definition["type"].(string) {
	case "string":
		if _, ok := definition["enum"]; ok {
			return true
		}
		if _, ok := definition["pattern"]; ok {
			return true
		}
		return false
	case "integer":
		if _, ok := definition["maximum"]; ok {
			return true
		}
		if _, ok := definition["minimum"]; ok {
			return true
		}
		return false
	default:
		// Unsupported types
		return false
	}
}

func extractAttrSchema(ref attributeRef, definition map[string]interface{}) attribute {
	resourceSchema, ok := terraformSchema.ResourceSchemas[ref.resource]
	if !ok {
		panic(fmt.Sprintf("resource `%s` not found in the Terraform schema", ref.resource))
	}
	if ref.block != nil {
		resourceSchema, ok = resourceSchema.Block.BlockTypes[*ref.block]
		if !ok {
			panic(fmt.Sprintf("block `%s.%s` not found in the Terraform schema", ref.resource, *ref.block))
		}
	}
	attrSchema, ok := resourceSchema.Block.Attributes[ref.attribute]
	if !ok {
		panic(fmt.Sprintf("`%s` not found in the Terraform schema", ref.String()))
	}

	switch definition["type"].(string) {
	case "string":
		ty, ok := attrSchema.Type.(string)
		if !ok {
			panic(fmt.Sprintf("`%s` is expected as string, but not (%s)", ref.String(), attrSchema.Type))
		}
		if ty != "string" && ty != "number" {
			panic(fmt.Sprintf("`%s` is expected as string, but not (%s)", ref.String(), attrSchema.Type))
		}
	case "integer":
		ty, ok := attrSchema.Type.(string)
		if !ok {
			panic(fmt.Sprintf("`%s` is expected as integer, but not (%s)", ref.String(), attrSchema.Type))
		}
		if ty != "number" && ty != "string" {
			panic(fmt.Sprintf("`%s` is expected as integer, but not (%s)", ref.String(), attrSchema.Type))
		}
	default:
		// noop
	}

	return attrSchema
}

func generateRuleFile(mapping mapping, ref attributeRef, definition map[string]interface{}, schema attribute) *ruleMeta {
	ruleName := ref.RuleName()
	var blockType string
	if ref.block != nil {
		blockType = *ref.block
	}

	meta := &ruleMeta{
		RuleName:      ruleName,
		RuleNameCC:    toCamelCase(ruleName),
		BlockType:     blockType,
		ResourceType:  mapping.Resource,
		AttributeName: ref.attribute,
		Sensitive:     schema.Sensitive,
		Max:           fetchNumber(definition, "maximum"),
		SetMax:        numberExists(definition, "maximum"),
		Min:           fetchNumber(definition, "minimum"),
		SetMin:        numberExists(definition, "minimum"),
		Pattern:       fetchString(definition, "pattern"),
		Enum:          fetchStrings(definition, "enum"),
		ReferenceURL:  fmt.Sprintf("https://github.com/Azure/azure-rest-api-specs/tree/master%s", strings.TrimPrefix(mapping.ImportPath, "azure-rest-api-specs")),
	}

	// Testing generated regexp
	regexp.MustCompile(meta.Pattern)

	generateFile(fmt.Sprintf("%s/apispec/%s.go", RulesPath, ruleName), ref.RuleTemplate(), meta)
	generateFile(fmt.Sprintf("%s/rules/%s.md", DocsPath, ruleName), getFullPath("rule.md.tmpl"), meta)

	return meta
}

func generateProviderFile(ruleNames []string) {
	meta := &providerMeta{RuleNameCCList: ruleNames}
	generateFile(fmt.Sprintf("%s/apispec/provider.go", RulesPath), getFullPath("provider.go.tmpl"), meta)
}

func generateRulesIndexDoc(ruleNames []string) {
	meta := &ruleDocIndexMeta{RuleNameList: ruleNames}
	generateFile(fmt.Sprintf("%s/README.md", DocsPath), getFullPath("doc_README.md.tmpl"), meta)
}

func fetchNumber(definition map[string]interface{}, key string) int {
	if v, ok := definition[key]; ok {
		return int(v.(float64))
	}
	return 0
}

func numberExists(definition map[string]interface{}, key string) bool {
	_, ok := definition[key]
	return ok
}

func fetchString(definition map[string]interface{}, key string) string {
	if v, ok := definition[key]; ok {
		return v.(string)
	}
	return ""
}

func fetchStrings(definition map[string]interface{}, key string) []string {
	if raw, ok := definition[key]; ok {
		list := raw.([]interface{})
		ret := make([]string, len(list))
		for i, v := range list {
			ret[i] = v.(string)
		}
		return ret
	}
	return []string{}
}

func generateFile(fileName string, tmplName string, meta interface{}) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles(tmplName))
	err = tmpl.Execute(file, meta)
	if err != nil {
		panic(err)
	}
}

var heading = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

func toCamelCase(str string) string {
	exceptions := map[string]string{
		"ip":  "IP",
		"sql": "SQL",
		"vm":  "VM",
		"os":  "OS",
		"id":  "ID",
		"tls": "TLS",
	}
	parts := strings.Split(str, "_")
	replaced := make([]string, len(parts))
	for i, s := range parts {
		conv, ok := exceptions[s]
		if ok {
			replaced[i] = conv
		} else {
			replaced[i] = s
		}
	}
	str = strings.Join(replaced, "_")

	return heading.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}
