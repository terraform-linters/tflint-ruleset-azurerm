package rules

// Please see: https://docs.microsoft.com/en-us/azure/virtual-machines/linux/faq#what-are-the-username-requirements-when-creating-a-vm

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualMachineInvalidAdminUserNameRule ensures that invalid usernames are not set.
type AzurermVirtualMachineInvalidAdminUserNameRule struct {
	resourceType  string
	attributeName string
}

// NewAzurermVirtualMachineInvalidAdminUserNameRule returns new rule with default attributes
func NewAzurermVirtualMachineInvalidAdminUserNameRule() *AzurermVirtualMachineInvalidAdminUserNameRule {
	return &AzurermVirtualMachineInvalidAdminUserNameRule{
		resourceType:  "azurerm_virtual_machine",
		attributeName: "os_profile",
	}
}

// Name returns the rule name
func (r *AzurermVirtualMachineInvalidAdminUserNameRule) Name() string {
	return "azurerm_virtual_machine_invalid_admin_username"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineInvalidAdminUserNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineInvalidAdminUserNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineInvalidAdminUserNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualMachineInvalidAdminUserNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceBlocks("azurerm_virtual_machine", "os_profile", func(block *hcl.Block) error {
		if err := runner.EmitIssue(r, "`os_profile` block found", block.DefRange); err != nil {
			return err
		}

		content, _, diags := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{Name: "admin_username"},
			},
		})
		if diags.HasErrors() {
			return diags
		}

		if attr, exists := content.Attributes["admin_username"]; exists {
			if err := runner.EmitIssueOnExpr(r, "`admin_username` attribute found", attr.Expr); err != nil {
				return err
			}

			var val string
			err := runner.EvaluateExpr(attr.Expr, &val)
			valid, err := isVlidVMAdminUserNames(val)
			if err != nil {
				panic(err)
			}
			if valid {
				return nil
			}

			return runner.EmitIssueOnExpr(
				r,
				fmt.Sprintf(`"%s" is not a valid VM Admin username`, val),
				attr.Expr,
			)
		}
		return nil
	})
}
