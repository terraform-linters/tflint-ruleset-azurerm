package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsVirtualMachineScaleSetInvalidSkuRule checks the pattern is valid
type AzurermWindowsVirtualMachineScaleSetInvalidSkuRule struct {
	resourceType  string
	attributeName string
}

// NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule() *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule {
	return &AzurermWindowsVirtualMachineScaleSetInvalidSkuRule{
		resourceType:  "azurerm_windows_virtual_machine_scale_set",
		attributeName: "sku",
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Name() string {
	return "azurerm_windows_virtual_machine_scale_set_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range validMachineSizes {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
