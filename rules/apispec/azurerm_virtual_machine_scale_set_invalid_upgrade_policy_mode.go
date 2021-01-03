// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule checks the pattern is valid
type AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule returns new rule with default attributes
func NewAzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule() *AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule {
	return &AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule{
		resourceType:  "azurerm_virtual_machine_scale_set",
		attributeName: "upgrade_policy_mode",
		enum: []string{
			"Automatic",
			"Manual",
			"Rolling",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule) Name() string {
	return "azurerm_virtual_machine_scale_set_invalid_upgrade_policy_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualMachineScaleSetInvalidUpgradePolicyModeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as upgrade_policy_mode`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
