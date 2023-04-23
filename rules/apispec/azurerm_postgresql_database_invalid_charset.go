// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermPostgresqlDatabaseInvalidCharsetRule checks the pattern is valid
type AzurermPostgresqlDatabaseInvalidCharsetRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermPostgresqlDatabaseInvalidCharsetRule returns new rule with default attributes
func NewAzurermPostgresqlDatabaseInvalidCharsetRule() *AzurermPostgresqlDatabaseInvalidCharsetRule {
	return &AzurermPostgresqlDatabaseInvalidCharsetRule{
		resourceType:  "azurerm_postgresql_database",
		attributeName: "charset",
		pattern:       regexp.MustCompile(`^[a-zA-Z]+\w*$`),
	}
}

// Name returns the rule name
func (r *AzurermPostgresqlDatabaseInvalidCharsetRule) Name() string {
	return "azurerm_postgresql_database_invalid_charset"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermPostgresqlDatabaseInvalidCharsetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermPostgresqlDatabaseInvalidCharsetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermPostgresqlDatabaseInvalidCharsetRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermPostgresqlDatabaseInvalidCharsetRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z]+\w*$`),
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
