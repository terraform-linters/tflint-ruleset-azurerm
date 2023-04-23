// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule checks the pattern is valid
type AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule returns new rule with default attributes
func NewAzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule() *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule {
	return &AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule{
		resourceType:  "azurerm_data_factory_dataset_sql_server_table",
		attributeName: "linked_service_name",
		pattern:       regexp.MustCompile(`^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`),
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Name() string {
	return "azurerm_data_factory_dataset_sql_server_table_invalid_linked_service_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryDatasetSQLServerTableInvalidLinkedServiceNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`),
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
