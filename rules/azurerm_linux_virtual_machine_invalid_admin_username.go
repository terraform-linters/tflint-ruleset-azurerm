package rules

import (
	"fmt"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLinuxVirtualMachineInvalidAdminUserNameRule ensures that invalid usernames are not set.
type AzurermLinuxVirtualMachineInvalidAdminUserNameRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// Please see: https://docs.microsoft.com/en-us/azure/virtual-machines/linux/faq#what-are-the-username-requirements-when-creating-a-vm
// NewAzurermLinuxVirtualMachineInvalidSizeRule returns new rule with default attributes
func NewAzurermLinuxVirtualMachineInvalidAdminUserNameRule() *AzurermLinuxVirtualMachineInvalidAdminUserNameRule {
	return &AzurermLinuxVirtualMachineInvalidAdminUserNameRule{
		resourceType:  "azurerm_linux_virtual_machine",
		attributeName: "admin_username",
		enum: []string{
			"azureuser",
			"1",
			"123",
			"a",
			"actuser",
			"adm",
			"admin",
			"admin1",
			"admin2",
			"administrator",
			"aspnet",
			"backup",
			"console",
			"david",
			"guest",
			"john",
			"owner",
			"root",
			"server",
			"sql",
			"support_388945a0",
			"support",
			"sys",
			"test",
			"test1",
			"test2",
			"test3",
			"user",
			"user1",
			"user2",
			"user3",
			"user4",
			"user5",
			"video",
		},
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
			found := false
			val = strings.ToLower(val)
			for _, item := range r.enum {
				if strings.ToLower(item) == val {
					found = true
				}
			}
			if found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is not a valid Linux VM Admin username`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}