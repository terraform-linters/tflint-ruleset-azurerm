// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermPostgresqlDatabaseInvalidCollationRule checks the pattern is valid
type AzurermPostgresqlDatabaseInvalidCollationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermPostgresqlDatabaseInvalidCollationRule returns new rule with default attributes
func NewAzurermPostgresqlDatabaseInvalidCollationRule() *AzurermPostgresqlDatabaseInvalidCollationRule {
	return &AzurermPostgresqlDatabaseInvalidCollationRule{
		resourceType:  "azurerm_postgresql_database",
		attributeName: "collation",
		pattern:       regexp.MustCompile(`^[a-zA-Z\-]+([. ]|\w)*$`),
	}
}

// Name returns the rule name
func (r *AzurermPostgresqlDatabaseInvalidCollationRule) Name() string {
	return "azurerm_postgresql_database_invalid_collation"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermPostgresqlDatabaseInvalidCollationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermPostgresqlDatabaseInvalidCollationRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermPostgresqlDatabaseInvalidCollationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermPostgresqlDatabaseInvalidCollationRule) Check(runner tflint.Runner) error {
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
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z\-]+([. ]|\w)*$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
