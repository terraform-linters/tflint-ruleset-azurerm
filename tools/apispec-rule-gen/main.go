package main

import (
	"encoding/json"
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
	"github.com/serenize/snaker"
)

type mappingFile struct {
	Mappings []mapping `hcl:"mapping,block"`
}

type mapping struct {
	Resource   string         `hcl:"resource,label"`
	ImportPath string         `hcl:"import_path"`
	Attrs      hcl.Attributes `hcl:",remain"`
}

type ruleMeta struct {
	RuleName      string
	RuleNameCC    string
	ResourceType  string
	AttributeName string
	Sensitive     bool
	Max           int
	Min           int
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

func main() {
	files, err := filepath.Glob("apispec-rule-gen/mappings/*.hcl")
	if err != nil {
		panic(err)
	}

	mappingFiles := []mappingFile{}
	for _, file := range files {
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
		mappingFiles = append(mappingFiles, mf)
	}

	provider := loadProviderSchema()

	generatedRuleNames := []string{}
	generatedRuleNameCCs := []string{}
	for _, mappingFile := range mappingFiles {
		for _, mapping := range mappingFile.Mappings {
			raw, err := ioutil.ReadFile(fmt.Sprintf("apispec-rule-gen/%s", mapping.ImportPath))
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

			for attribute, value := range mapping.Attrs {
				variable := value.Expr.Variables()[0]
				props := []string{}
				for _, prop := range variable {
					if len(props) == 0 {
						props = []string{prop.(hcl.TraverseRoot).Name}
					} else {
						props = append(props, prop.(hcl.TraverseAttr).Name)
					}
				}

				var definition map[string]interface{}
				if len(props) == 1 {
					if props[0] == "any" {
						continue
					}
					if definitions[props[0]] != nil {
						definition = definitions[props[0]].(map[string]interface{})
					} else {
						definition = parameters[props[0]].(map[string]interface{})
					}
				} else {
					definition = definitions[props[0]].(map[string]interface{})["properties"].(map[string]interface{})[props[1]].(map[string]interface{})
				}

				if validMapping(definition) {
					attrSchema := extractAttrSchema(mapping.Resource, attribute, definition, provider)
					meta := generateRuleFile(mapping, attribute, definition, attrSchema)
					generatedRuleNames = append(generatedRuleNames, meta.RuleName)
					generatedRuleNameCCs = append(generatedRuleNameCCs, meta.RuleNameCC)
				}
			}
		}
	}

	sort.Strings(generatedRuleNameCCs)
	generateProviderFile(generatedRuleNameCCs)
	sort.Strings(generatedRuleNames)
	generateRulesIndexDoc(generatedRuleNames)
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

func extractAttrSchema(resource, attribute string, definition map[string]interface{}, provider provider) attribute {
	resourceSchema, ok := provider.ResourceSchemas[resource]
	if !ok {
		panic(fmt.Sprintf("resource `%s` not found in the Terraform schema", resource))
	}
	attrSchema, ok := resourceSchema.Block.Attributes[attribute]
	if !ok {
		panic(fmt.Sprintf("`%s.%s` not found in the Terraform schema", resource, attribute))
	}

	switch definition["type"].(string) {
	case "string":
		ty, ok := attrSchema.Type.(string)
		if !ok {
			panic(fmt.Sprintf("`%s.%s` is expected as string, but not (%s)", resource, attribute, attrSchema.Type))
		}
		if ty != "string" && ty != "number" {
			panic(fmt.Sprintf("`%s.%s` is expected as string, but not (%s)", resource, attribute, attrSchema.Type))
		}
	case "integer":
		ty, ok := attrSchema.Type.(string)
		if !ok {
			panic(fmt.Sprintf("`%s.%s` is expected as integer, but not (%s)", resource, attribute, attrSchema.Type))
		}
		if ty != "number" && ty != "string" {
			panic(fmt.Sprintf("`%s.%s` is expected as integer, but not (%s)", resource, attribute, attrSchema.Type))
		}
	default:
		// noop
	}

	return attrSchema
}

func generateRuleFile(mapping mapping, attribute string, definition map[string]interface{}, schema attribute) *ruleMeta {
	ruleName := fmt.Sprintf("%s_invalid_%s", mapping.Resource, attribute)

	meta := &ruleMeta{
		RuleName:      ruleName,
		RuleNameCC:    snaker.SnakeToCamel(ruleName),
		ResourceType:  mapping.Resource,
		AttributeName: attribute,
		Sensitive:     schema.Sensitive,
		Max:           fetchNumber(definition, "maximum"),
		Min:           fetchNumber(definition, "minimum"),
		Pattern:       fetchString(definition, "pattern"),
		Enum:          fetchStrings(definition, "enum"),
		ReferenceURL:  fmt.Sprintf("https://github.com/Azure/azure-rest-api-specs/tree/master%s", strings.TrimPrefix(mapping.ImportPath, "azure-rest-api-specs")),
	}

	// Testing generated regexp
	regexp.MustCompile(meta.Pattern)

	generateFile(fmt.Sprintf("../rules/apispec/%s.go", ruleName), "apispec-rule-gen/rule.go.tmpl", meta)
	generateFile(fmt.Sprintf("../docs/rules/%s.md", ruleName), "apispec-rule-gen/rule.md.tmpl", meta)

	return meta
}

func generateProviderFile(ruleNames []string) {
	meta := &providerMeta{}

	for _, ruleName := range ruleNames {
		meta.RuleNameCCList = append(meta.RuleNameCCList, ruleName)
	}

	generateFile("../rules/apispec/provider.go", "apispec-rule-gen/provider.go.tmpl", meta)
}

func generateRulesIndexDoc(ruleNames []string) {
	meta := &ruleDocIndexMeta{}

	for _, ruleName := range ruleNames {
		meta.RuleNameList = append(meta.RuleNameList, ruleName)
	}

	generateFile("../docs/README.md", "apispec-rule-gen/doc_README.md.tmpl", meta)
}

func fetchNumber(definition map[string]interface{}, key string) int {
	if v, ok := definition[key]; ok {
		return int(v.(float64))
	}
	return 0
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
