package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsVirtualMachineRetiredSizeRule checks that Windows Virtual Machine is not configured to use a retired Virtual Machine size.
type AzurermWindowsVirtualMachineRetiredSizeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewAzurermWindowsVirtualMachineRetiredSizeRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineRetiredSizeRule() *AzurermWindowsVirtualMachineRetiredSizeRule {
	return &AzurermWindowsVirtualMachineRetiredSizeRule{
		resourceType:  "azurerm_windows_virtual_machine",
		attributeName: "size",
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineRetiredSizeRule) Name() string {
	return "azurerm_windows_virtual_machine_retired_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineRetiredSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineRetiredSizeRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineRetiredSizeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks that the azurerm_windows_virtual_machine.size is not configured to use a retried Virtual Machine size.
func (r *AzurermWindowsVirtualMachineRetiredSizeRule) Check(runner tflint.Runner) error {
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
			for _, item := range retiredMachineSizes {
				if item == val {
					found = true
				}
			}
			if found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is a retired or announced-for-retirement Virtual Machine Size`, val),
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
