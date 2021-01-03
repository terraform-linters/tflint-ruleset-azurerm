// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule checks the pattern is valid
type AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule returns new rule with default attributes
func NewAzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule() *AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule {
	return &AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule{
		resourceType:  "azurerm_data_factory_dataset_mysql",
		attributeName: "data_factory_name",
		pattern:       regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule) Name() string {
	return "azurerm_data_factory_dataset_mysql_invalid_data_factory_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDataFactoryDatasetMysqlInvalidDataFactoryNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
