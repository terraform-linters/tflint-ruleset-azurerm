package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLinuxVirtualMachineScaleSetInvalidSkuRule checks the pattern is valid
type AzurermLinuxVirtualMachineScaleSetInvalidSkuRule struct {
	resourceType  string
	attributeName string
}

// NewAzurermLinuxVirtualMachineScaleSetInvalidSkuRule returns new rule with default attributes
func NewAzurermLinuxVirtualMachineScaleSetInvalidSkuRule() *AzurermLinuxVirtualMachineScaleSetInvalidSkuRule {
	return &AzurermLinuxVirtualMachineScaleSetInvalidSkuRule{
		resourceType:  "azurerm_linux_virtual_machine_scale_set",
		attributeName: "sku",
	}
}

// Name returns the rule name
func (r *AzurermLinuxVirtualMachineScaleSetInvalidSkuRule) Name() string {
	return "azurerm_linux_virtual_machine_scale_set_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLinuxVirtualMachineScaleSetInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLinuxVirtualMachineScaleSetInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermLinuxVirtualMachineScaleSetInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermLinuxVirtualMachineScaleSetInvalidSkuRule) Check(runner tflint.Runner) error {
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
