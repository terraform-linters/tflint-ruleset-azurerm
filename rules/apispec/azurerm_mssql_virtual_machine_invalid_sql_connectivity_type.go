// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule checks the pattern is valid
type AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule returns new rule with default attributes
func NewAzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule() *AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule {
	return &AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule{
		resourceType:  "azurerm_mssql_virtual_machine",
		attributeName: "sql_connectivity_type",
		enum: []string{
			"LOCAL",
			"PRIVATE",
			"PUBLIC",
		},
	}
}

// Name returns the rule name
func (r *AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule) Name() string {
	return "azurerm_mssql_virtual_machine_invalid_sql_connectivity_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermMssqlVirtualMachineInvalidSQLConnectivityTypeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sql_connectivity_type`, truncateLongMessage(val)),
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
