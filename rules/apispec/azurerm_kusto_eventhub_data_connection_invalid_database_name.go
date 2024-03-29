// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule checks the pattern is valid
type AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermKustoEventhubDataConnectionInvalidDatabaseNameRule returns new rule with default attributes
func NewAzurermKustoEventhubDataConnectionInvalidDatabaseNameRule() *AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule {
	return &AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule{
		resourceType:  "azurerm_kusto_eventhub_data_connection",
		attributeName: "database_name",
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule) Name() string {
	return "azurerm_kusto_eventhub_data_connection_invalid_database_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermKustoEventhubDataConnectionInvalidDatabaseNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*$`),
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
