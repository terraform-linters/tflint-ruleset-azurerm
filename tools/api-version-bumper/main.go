package main

import (
	"cmp"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/agext/levenshtein"
)

func main() {
	// go run ./api-version-bumper <path-to-file>
	if len(os.Args) == 2 {
		processFile(os.Args[1])
		return
	}

	// Find Terraform resource files
	files, err := filepath.Glob("api-version-bumper/terraform-provider-azurerm/internal/services/*/*_resource.go")
	if err != nil {
		panic(err)
	}

	var errs []error
	for _, file := range files {
		// As an exception, the actual name of `user_assigned_identity_resource` is `*_resource_gen`, so replace the file name.
		// https://github.com/hashicorp/terraform-provider-azurerm/tree/v4.23.0/internal/services/managedidentity
		if strings.HasSuffix(file, "user_assigned_identity_resource.go") {
			file = strings.ReplaceAll(file, "user_assigned_identity_resource.go", "user_assigned_identity_resource_gen.go")
		}
		if err := processFile(file); err != nil {
			errs = append(errs, err)
		}
	}
	if errs != nil {
		fmt.Print("\n!!!!!!!!!! Failed to parse some files !!!!!!!!!!\n\n")
		for _, err := range errs {
			fmt.Println(err)
		}
	}
}

var apiVersionRE = regexp.MustCompile(`\d{4}-\d{2}-\d{2}(-preview)?`)

func processFile(filePath string) error {
	fmt.Printf("Processing %s ...\n", filePath)

	resourceType := findResourceType(filePath)
	apiVersion := findAPIVersion(filePath)

	fmt.Printf("%s: %s\n", resourceType, apiVersion)

	if resourceType == "" || apiVersion == "" {
		return fmt.Errorf("path: %s\nresource type: %s\nAPI version: %s\n", filePath, resourceType, apiVersion)
	}

	mappingFile := fmt.Sprintf("apispec-rule-gen/mappings/%s.hcl", resourceType)
	content, err := os.ReadFile(mappingFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		panic(err)
	}
	new := apiVersionRE.ReplaceAll(content, []byte(apiVersion))
	if err := os.WriteFile(mappingFile, new, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func findResourceType(resourcePath string) string {
	resourceType, resourceFunc := extractResourceTypeOrResourceFunc(resourcePath)
	if resourceType != "" {
		return resourceType
	}

	registrationPath := filepath.Join(filepath.Dir(resourcePath), "registration.go")
	return extractResourceTypeFromResourceFunc(registrationPath, resourceFunc)
}

// Extract a resource type from `*_resource.go`.
// If the resource type is not found, returns the name of a function
// that will return the resource for further processing.
func extractResourceTypeOrResourceFunc(resourcePath string) (resourceType string, resourceFunc string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, resourcePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// A string literal beginning with "azurerm_" is likely a resource type.
		if basicLit, ok := n.(*ast.BasicLit); ok {
			str := strings.Trim(basicLit.Value, "\"")
			if strings.HasPrefix(str, "azurerm_") {
				resourceType = str
			}
		}

		// Extract the function name that returns *pluginsdk.Resource.
		// This function may be linked to the resource type in another file.
		//
		// e.g. `func resourceAnalysisServicesServer() *pluginsdk.Resource {`
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			if funcDecl.Type.Results != nil && len(funcDecl.Type.Results.List) == 1 {
				if starExpr, ok := funcDecl.Type.Results.List[0].Type.(*ast.StarExpr); ok {
					if selExpr, ok := starExpr.X.(*ast.SelectorExpr); ok {
						if selExpr.Sel.Name == "Resource" {
							resourceFunc = funcDecl.Name.Name
						}
					}
				}
			}
		}

		return true
	})

	return
}

// Extract a resource type from `registration.go`.
// This file generally defines a map that links function names with resource types.
func extractResourceTypeFromResourceFunc(registrationPath, resourceFunc string) string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, registrationPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var resourceType string
	ast.Inspect(node, func(n ast.Node) bool {
		// In registration.go, a map is defined with the resource type as the key and the function call
		// that returns the resource as the value, so we search for a key in KeyValueExpr
		// whose value matches the resourceFunc.
		//
		// e.g.
		//
		// map[string]*pluginsdk.Resource{
		// 	   "azurerm_analysis_services_server": resourceAnalysisServicesServer(),
		// }
		if kvExpr, ok := n.(*ast.KeyValueExpr); ok {
			var key string
			if basicLit, ok := kvExpr.Key.(*ast.BasicLit); ok {
				key = strings.Trim(basicLit.Value, "\"")
			}

			if callExpr, ok := kvExpr.Value.(*ast.CallExpr); ok {
				if ident, ok := callExpr.Fun.(*ast.Ident); ok {
					if ident.Name == resourceFunc {
						resourceType = key
					}
				}
			}
		}

		// e.g. `resources["azurerm_security_center_auto_provisioning"] = resourceSecurityCenterAutoProvisioning()`
		if assign, ok := n.(*ast.AssignStmt); ok {
			var key string
			if indexExpr, ok := assign.Lhs[0].(*ast.IndexExpr); ok {
				if basicLit, ok := indexExpr.Index.(*ast.BasicLit); ok {
					key = strings.Trim(basicLit.Value, "\"")
				}
			}

			if callExpr, ok := assign.Rhs[0].(*ast.CallExpr); ok {
				if ident, ok := callExpr.Fun.(*ast.Ident); ok {
					if ident.Name == resourceFunc {
						resourceType = key
					}
				}
			}
		}
		return true
	})

	return resourceType
}

var exceptions = map[string]string{
	"key_vault":            "2023-02-01", // https://github.com/hashicorp/terraform-provider-azurerm/blob/v4.23.0/internal/services/keyvault/key_vault_resource.go
	"monitor_action_group": "2023-01-01", // https://github.com/hashicorp/terraform-provider-azurerm/blob/v4.23.0/internal/services/monitor/monitor_action_group_resource.go
}

func findAPIVersion(resourcePath string) string {
	// Infers the name of the resource from the resource path.
	// e.g. application_gateway_resource.go -> application_gateway
	//
	// Used to identify the most likely path among multiple import statements.
	resource := strings.ReplaceAll(filepath.Base(resourcePath), "_resource.go", "")

	// For some reason, some cases are not easily identifiable, so we apply a dictionary to them.
	if apiVersion, ok := exceptions[resource]; ok {
		return apiVersion
	}

	apiVersion := extractAPIVersion(resourcePath, resource)
	if apiVersion != "" {
		return apiVersion
	}

	// If the API version is not found in the resource path, it is likely in another file
	// that handles client-related tasks, so we will try there.
	files, err := filepath.Glob(fmt.Sprintf("%s/client/*.go", filepath.Dir(resourcePath)))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		apiVersion = extractAPIVersion(file, resource)
		if apiVersion != "" {
			return apiVersion
		}
	}
	return ""
}

// github.com/hashicorp/go-azure-sdk/resource-manager/aadb2c/2021-04-01-preview/tenants
var goAzureSDKRE = regexp.MustCompile(`github\.com/hashicorp/go-azure-sdk/resource-manager/[^/]+/([^/]+)/([^/]+)`)

// github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn
var azureSDKForGoRE = regexp.MustCompile(`github\.com/Azure/azure-sdk-for-go/services/[^/]+/mgmt/([^/]+)/([^/]+)`)

// github.com/jackofallops/kermit/sdk/synapse/2019-06-01-preview/synapse
var kermitSDKRE = regexp.MustCompile(`github\.com/jackofallops/kermit/sdk/[^/]+/([^/]+)/([^/]+)`)

type importSpec struct {
	apiVersion string
	resource   string
}

// Extract an API version from Go import statements by regexp.
// If multiple import statements are found, the most likely path
// is chosen by comparing it with the given resource name.
func extractAPIVersion(path, resource string) string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	matches := []importSpec{}
	for _, imp := range node.Imports {
		importPath := strings.Trim(imp.Path.Value, "\"")
		if match := goAzureSDKRE.FindStringSubmatch(importPath); match != nil {
			matches = append(matches, importSpec{apiVersion: match[1], resource: match[2]})
		}
		if match := azureSDKForGoRE.FindStringSubmatch(importPath); match != nil {
			matches = append(matches, importSpec{apiVersion: match[1], resource: match[2]})
		}
		if match := kermitSDKRE.FindStringSubmatch(importPath); match != nil {
			matches = append(matches, importSpec{apiVersion: match[1], resource: match[2]})
		}
	}

	if len(matches) > 0 {
		// Pattern 1: Calculate the Longest common subsequence (LCS) of two strings.
		// If there is only one path that matches with at least 4 characters, use that.
		longMatches := []importSpec{}
		for _, match := range matches {
			if len(longestCommonSubstring(match.resource, resource)) > 3 {
				longMatches = append(longMatches, match)
			}
		}
		if len(longMatches) == 1 {
			return longMatches[0].apiVersion
		}
		// Pattern 2: Calculate the Levenshtein distance between two characters and use the smaller one.
		spec := slices.MinFunc(matches, func(a, b importSpec) int {
			as := levenshtein.Distance(a.resource, resource, nil)
			bs := levenshtein.Distance(b.resource, resource, nil)
			return cmp.Compare(as, bs)
		})
		return spec.apiVersion
	}
	return ""
}

func longestCommonSubstring(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	maxLength := 0
	endPosition := 0

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endPosition = i
				}
			}
		}
	}

	if maxLength == 0 {
		return ""
	}
	return str1[endPosition-maxLength : endPosition]
}
