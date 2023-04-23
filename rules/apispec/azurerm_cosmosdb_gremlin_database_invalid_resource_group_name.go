// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule checks the pattern is valid
type AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule() *AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule {
	return &AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule{
		resourceType:  "azurerm_cosmosdb_gremlin_database",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[-\w\._\(\)]+$`),
	}
}

// Name returns the rule name
func (r *AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule) Name() string {
	return "azurerm_cosmosdb_gremlin_database_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCosmosdbGremlinDatabaseInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-\w\._\(\)]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
