package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermMysqlServerInvalidVersionRule checks the pattern is valid
type AzurermMysqlServerInvalidVersionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermMysqlServerInvalidVersionRule returns new rule with default attributes
func NewAzurermMysqlServerInvalidVersionRule() *AzurermMysqlServerInvalidVersionRule {
	return &AzurermMysqlServerInvalidVersionRule{
		resourceType:  "azurerm_mysql_server",
		attributeName: "version",
		enum: []string{
			"5.7",
			"8.0",
		},
	}
}

// Name returns the rule name
func (r *AzurermMysqlServerInvalidVersionRule) Name() string {
	return "azurerm_mysql_server_invalid_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermMysqlServerInvalidVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermMysqlServerInvalidVersionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermMysqlServerInvalidVersionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermMysqlServerInvalidVersionRule) Check(runner tflint.Runner) error {
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
		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as version`, val),
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
