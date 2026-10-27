//go:build generators
// +build generators

// This generator reads a local checkout of hashicorp/terraform-provider-azurerm and generates:
//
//   - resources.go: the Azure resource provider namespaces used by each resource type, taken from the
//     generated "## API Providers" section of the resource documentation.
//   - registrations.go: the namespaces registered by each resource_provider_registrations mode, taken
//     from internal/resourceproviders/required.go.
//
// Usage (from rules/resourceproviders):
//
//	go run -tags generators ./generator -azurerm-dir /path/to/terraform-provider-azurerm
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"text/template"
)

const (
	registrationsFilename = "registrations.go"
	resourcesFilename     = "resources.go"
)

var (
	apiProviderLine = regexp.MustCompile("^\\* `([^`]+)`")
	registrationSet = map[string]string{
		"Core":     "core",
		"Extended": "extended",
		"All":      "all",
		"Legacy":   "legacy",
	}
)

func main() {
	azurermDir := flag.String("azurerm-dir", os.Getenv("AZURERM_DIR"), "Path to a local checkout of hashicorp/terraform-provider-azurerm")
	flag.Parse()

	if err := run(*azurermDir); err != nil {
		log.Fatal(err)
	}
}

func run(azurermDir string) error {
	if azurermDir == "" {
		return fmt.Errorf("-azurerm-dir (or AZURERM_DIR) is required")
	}

	resources, err := loadResources(filepath.Join(azurermDir, "website", "docs", "r"))
	if err != nil {
		return fmt.Errorf("loading resource documentation: %w", err)
	}

	registrations, err := loadRegistrations(filepath.Join(azurermDir, "internal", "resourceproviders", "required.go"))
	if err != nil {
		return fmt.Errorf("loading registration sets: %w", err)
	}

	if err := write(resourcesFilename, resourcesTemplate, resources); err != nil {
		return err
	}
	return write(registrationsFilename, registrationsTemplate, registrations)
}

// loadResources returns the API provider namespaces per resource type, skipping resources without the section.
func loadResources(docsDir string) (map[string][]string, error) {
	files, err := filepath.Glob(filepath.Join(docsDir, "*.html.markdown"))
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no resource documentation found in %s", docsDir)
	}

	resources := map[string][]string{}
	for _, file := range files {
		namespaces, err := parseAPIProviders(file)
		if err != nil {
			return nil, err
		}
		if len(namespaces) == 0 {
			continue
		}
		name := "azurerm_" + strings.TrimSuffix(filepath.Base(file), ".html.markdown")
		resources[name] = namespaces
	}

	return resources, nil
}

func parseAPIProviders(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	seen := map[string]struct{}{}
	inSection := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "## ") {
			inSection = line == "## API Providers"
			continue
		}
		if !inSection {
			continue
		}
		if match := apiProviderLine.FindStringSubmatch(line); match != nil {
			seen[match[1]] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return slices.Sorted(maps.Keys(seen)), nil
}

// loadRegistrations parses the Core, Extended, All and Legacy functions in required.go. A function body
// contains map literals with the namespaces, and may call another set function that it merges with.
func loadRegistrations(file string) (map[string][]string, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, file, nil, 0)
	if err != nil {
		return nil, err
	}

	funcs := map[string]*ast.FuncDecl{}
	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Recv == nil {
			funcs[fn.Name.Name] = fn
		}
	}

	var resolve func(name string) (map[string]struct{}, error)
	resolve = func(name string) (map[string]struct{}, error) {
		fn, ok := funcs[name]
		if !ok {
			return nil, fmt.Errorf("function %s not found in %s", name, file)
		}

		result := map[string]struct{}{}
		var inspectErr error
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.KeyValueExpr:
				if lit, ok := v.Key.(*ast.BasicLit); ok && lit.Kind == token.STRING {
					value, err := strconv.Unquote(lit.Value)
					if err != nil {
						inspectErr = err
						return false
					}
					result[value] = struct{}{}
				}
			case *ast.CallExpr:
				if ident, ok := v.Fun.(*ast.Ident); ok {
					if _, isSet := registrationSet[ident.Name]; isSet {
						nested, err := resolve(ident.Name)
						if err != nil {
							inspectErr = err
							return false
						}
						for namespace := range nested {
							result[namespace] = struct{}{}
						}
					}
				}
			}
			return true
		})

		return result, inspectErr
	}

	registrations := map[string][]string{"none": {}}
	for funcName, mode := range registrationSet {
		set, err := resolve(funcName)
		if err != nil {
			return nil, err
		}
		if len(set) == 0 {
			return nil, fmt.Errorf("no resource providers found for %s", funcName)
		}
		registrations[mode] = slices.Sorted(maps.Keys(set))
	}

	return registrations, nil
}

func write(filename string, body string, data map[string][]string) error {
	tpl, err := template.New(filename).Parse(body)
	if err != nil {
		return fmt.Errorf("parsing template for %s: %w", filename, err)
	}

	var buffer bytes.Buffer
	if err := tpl.Execute(&buffer, map[string]any{"Keys": slices.Sorted(maps.Keys(data)), "Data": data}); err != nil {
		return fmt.Errorf("executing template for %s: %w", filename, err)
	}

	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("formatting %s: %w", filename, err)
	}

	if err := os.WriteFile(filename, formatted, 0o644); err != nil {
		return fmt.Errorf("writing %s: %w", filename, err)
	}
	return nil
}

const resourcesTemplate = `// Code generated by generator/main.go; DO NOT EDIT.

package resourceproviders

// Resources maps azurerm resource types to the Azure resource provider namespaces they use.
var Resources = map[string][]string{
{{- range $key := .Keys }}
	{{ printf "%q" $key }}: { {{- range index $.Data $key }}{{ printf "%q" . }}, {{ end -}} },
{{- end }}
}
`

const registrationsTemplate = `// Code generated by generator/main.go; DO NOT EDIT.

package resourceproviders

// Registrations maps resource_provider_registrations modes to the namespaces the azurerm provider registers.
var Registrations = map[string][]string{
{{- range $key := .Keys }}
	{{ printf "%q" $key }}: {
	{{- range index $.Data $key }}
		{{ printf "%q" . }},
	{{- end }}
	},
{{- end }}
}
`
