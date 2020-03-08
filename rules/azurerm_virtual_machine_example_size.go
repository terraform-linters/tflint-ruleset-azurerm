package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AzurermVirtualMachineExampleSizeRule checks whether ...
type AzurermVirtualMachineExampleSizeRule struct{}

// NewAzurermVirtualMachineExampleSizeRule returns a new rule
func NewAzurermVirtualMachineExampleSizeRule() *AzurermVirtualMachineExampleSizeRule {
	return &AzurermVirtualMachineExampleSizeRule{}
}

// Name returns the rule name
func (r *AzurermVirtualMachineExampleSizeRule) Name() string {
	return "azurerm_virtual_machine_example_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineExampleSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineExampleSizeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineExampleSizeRule) Link() string {
	return ""
}

// Check checks whether ...
func (r *AzurermVirtualMachineExampleSizeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("azurerm_virtual_machine", "vm_size", func(attribute *hcl.Attribute) error {
		var vmSize string
		err := runner.EvaluateExpr(attribute.Expr, &vmSize)

		return runner.EnsureNoError(err, func() error {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("VM size is %s", vmSize),
				attribute.Expr.Range(),
				tflint.Metadata{Expr: attribute.Expr},
			)
		})
	})
}
