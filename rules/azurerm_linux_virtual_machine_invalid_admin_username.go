package rules

// Please see: https://docs.microsoft.com/en-us/azure/virtual-machines/linux/faq#what-are-the-username-requirements-when-creating-a-vm

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLinuxVirtualMachineInvalidAdminUserNameRule ensures that invalid usernames are not set.
type AzurermLinuxVirtualMachineInvalidAdminUserNameRule struct {
	resourceType  string
	attributeName string
}

// NewAzurermLinuxVirtualMachineInvalidAdminUserNameRule returns new rule with default attributes
func NewAzurermLinuxVirtualMachineInvalidAdminUserNameRule() *AzurermLinuxVirtualMachineInvalidAdminUserNameRule {
	return &AzurermLinuxVirtualMachineInvalidAdminUserNameRule{
		resourceType:  "azurerm_linux_virtual_machine",
		attributeName: "admin_username",
	}
}

// Name returns the rule name
func (r *AzurermLinuxVirtualMachineInvalidAdminUserNameRule) Name() string {
	return "azurerm_linux_virtual_machine_invalid_admin_username"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLinuxVirtualMachineInvalidAdminUserNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLinuxVirtualMachineInvalidAdminUserNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermLinuxVirtualMachineInvalidAdminUserNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermLinuxVirtualMachineInvalidAdminUserNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			valid, err := isVlidVMAdminUserNames(val)
			if err != nil {
				panic(err)
			}

			if valid {
				return nil
			}

			return runner.EmitIssueOnExpr(
				r,
				fmt.Sprintf(`"%s" is not a valid Linux VM Admin username`, val),
				attribute.Expr,
			)
		})
	})
}
