package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualMachineInvalidVMSizeRule checks the pattern is valid
type AzurermVirtualMachineInvalidVMSizeRule struct {
	resourceType  string
	attributeName string
}

// NewAzurermVirtualMachineInvalidVMSizeRule returns new rule with default attributes
func NewAzurermVirtualMachineInvalidVMSizeRule() *AzurermVirtualMachineInvalidVMSizeRule {
	return &AzurermVirtualMachineInvalidVMSizeRule{
		resourceType:  "azurerm_virtual_machine",
		attributeName: "vm_size",
	}
}

// Name returns the rule name
func (r *AzurermVirtualMachineInvalidVMSizeRule) Name() string {
	return "azurerm_virtual_machine_invalid_vm_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineInvalidVMSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineInvalidVMSizeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineInvalidVMSizeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualMachineInvalidVMSizeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as vm_size`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
