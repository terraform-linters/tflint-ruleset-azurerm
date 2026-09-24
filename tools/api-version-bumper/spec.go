package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var importPathRE = regexp.MustCompile(`import_path\s*=\s*"([^"]+)"`)

// "$ref": "./common.json#/definitions/Foo" -> ./common.json
var externalRefRE = regexp.MustCompile(`"\$ref"\s*:\s*"([^"#]+)#`)

type definition struct {
	Properties map[string]json.RawMessage `json:"properties"`
}

type parameter struct {
	Ref  string `json:"$ref"`
	Name string `json:"name"`
}

type apiSpec struct {
	Definitions map[string]definition                 `json:"definitions"`
	Parameters  map[string]json.RawMessage            `json:"parameters"`
	Paths       map[string]map[string]json.RawMessage `json:"paths"`
}

// resolveImportPath checks that the API spec referenced by the mapping content defines
// all references in the mapping. Spec files are often renamed or merged between API versions
// (e.g. bastionHost.json was merged into virtualNetwork.json), so if the spec does not exist
// or lacks the references, other specs in the same directory are searched instead.
// Returns false if no spec defines all references.
func resolveImportPath(content []byte) ([]byte, bool) {
	match := importPathRE.FindSubmatch(content)
	if match == nil {
		return content, true
	}
	importPath := string(match[1])

	refs, err := mappingReferences(content)
	if err != nil {
		panic(err)
	}

	if ok, _ := specDefines(specFullPath(importPath), refs); ok {
		return content, true
	}

	candidates, err := filepath.Glob(filepath.Join(filepath.Dir(specFullPath(importPath)), "*.json"))
	if err != nil {
		panic(err)
	}
	sort.Strings(candidates)
	// Prefer the spec that defines the most references by itself rather than via other files
	best, bestLocal := "", -1
	for _, candidate := range candidates {
		if ok, local := specDefines(candidate, refs); ok && local > bestLocal {
			best, bestLocal = candidate, local
		}
	}
	if best == "" {
		return content, false
	}
	newPath := filepath.ToSlash(filepath.Join(filepath.Dir(importPath), filepath.Base(best)))
	fmt.Printf("%s is not available, use %s instead\n", importPath, newPath)
	return []byte(strings.Replace(string(content), importPath, newPath, 1)), true
}

func specFullPath(importPath string) string {
	return filepath.Join("apispec-rule-gen", importPath)
}

// mappingReferences returns references in the mapping, like "Foo" or "Foo.Bar".
func mappingReferences(content []byte) ([][]string, error) {
	file, diags := hclsyntax.ParseConfig(content, "mapping.hcl", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, diags
	}

	refs := [][]string{}
	for _, block := range file.Body.(*hclsyntax.Body).Blocks {
		for _, attr := range block.Body.Attributes {
			refs = appendReferences(refs, attr.Expr)
		}
	}
	return refs, nil
}

func appendReferences(refs [][]string, expr hclsyntax.Expression) [][]string {
	switch expr := expr.(type) {
	case *hclsyntax.ScopeTraversalExpr:
		names := []string{expr.Traversal.RootName()}
		for _, traverser := range expr.Traversal[1:] {
			if attr, ok := traverser.(hcl.TraverseAttr); ok {
				names = append(names, attr.Name)
			}
		}
		if names[0] != "any" {
			refs = append(refs, names)
		}
	case *hclsyntax.ObjectConsExpr:
		for _, item := range expr.Items {
			refs = appendReferences(refs, item.ValueExpr)
		}
	}
	return refs
}

// specDefines reports whether the spec defines all references, and how many of them
// are defined in the spec itself rather than in other files referenced by the spec.
// A reference "Foo" is looked up in definitions, parameters and path parameters,
// and "Foo.Bar" is looked up in the properties of the "Foo" definition.
func specDefines(specPath string, refs [][]string) (bool, int) {
	raw, err := os.ReadFile(specPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, 0
		}
		panic(err)
	}
	var spec apiSpec
	if err := json.Unmarshal(raw, &spec); err != nil {
		panic(fmt.Errorf("%s: %w", specPath, err))
	}

	// Definitions may be defined in other files referenced by the spec (e.g. common.json)
	definitions := map[string]definition{}
	refFiles := map[string]bool{}
	for _, match := range externalRefRE.FindAllSubmatch(raw, -1) {
		refFiles[string(match[1])] = true
	}
	for refFile := range refFiles {
		refRaw, err := os.ReadFile(filepath.Join(filepath.Dir(specPath), refFile))
		if err != nil {
			continue
		}
		var refSpec apiSpec
		if json.Unmarshal(refRaw, &refSpec) != nil {
			continue
		}
		for name, def := range refSpec.Definitions {
			definitions[name] = def
		}
	}
	for name, def := range spec.Definitions {
		definitions[name] = def
	}

	params := map[string]bool{}
	for name := range spec.Parameters {
		params[name] = true
	}
	for _, pathItem := range spec.Paths {
		for _, raw := range pathItem {
			// path-level "parameters" is an array, others are operation objects
			var operation struct {
				Parameters []parameter `json:"parameters"`
			}
			if json.Unmarshal(raw, &operation.Parameters) != nil && json.Unmarshal(raw, &operation) != nil {
				continue
			}
			for _, param := range operation.Parameters {
				if param.Ref != "" {
					params[param.Ref[strings.LastIndex(param.Ref, "/")+1:]] = true
				} else {
					params[param.Name] = true
				}
			}
		}
	}

	local := 0
	for _, ref := range refs {
		if _, ok := spec.Definitions[ref[0]]; ok || (len(ref) == 1 && params[ref[0]]) {
			local++
		}

		def, defined := definitions[ref[0]]
		if len(ref) == 1 {
			if !defined && !params[ref[0]] {
				return false, 0
			}
			continue
		}
		if !defined {
			return false, 0
		}
		if _, ok := def.Properties[ref[1]]; !ok {
			return false, 0
		}
	}
	return true, local
}
