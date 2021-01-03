// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDevTestPolicyInvalidNameRule checks the pattern is valid
type AzurermDevTestPolicyInvalidNameRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDevTestPolicyInvalidNameRule returns new rule with default attributes
func NewAzurermDevTestPolicyInvalidNameRule() *AzurermDevTestPolicyInvalidNameRule {
	return &AzurermDevTestPolicyInvalidNameRule{
		resourceType:  "azurerm_dev_test_policy",
		attributeName: "name",
		enum: []string{
			"UserOwnedLabVmCount",
			"UserOwnedLabPremiumVmCount",
			"LabVmCount",
			"LabPremiumVmCount",
			"LabVmSize",
			"GalleryImage",
			"UserOwnedLabVmCountInSubnet",
			"LabTargetCost",
		},
	}
}

// Name returns the rule name
func (r *AzurermDevTestPolicyInvalidNameRule) Name() string {
	return "azurerm_dev_test_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDevTestPolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDevTestPolicyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDevTestPolicyInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDevTestPolicyInvalidNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as name`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
