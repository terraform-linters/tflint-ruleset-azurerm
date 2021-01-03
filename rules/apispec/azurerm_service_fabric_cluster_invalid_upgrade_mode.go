// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermServiceFabricClusterInvalidUpgradeModeRule checks the pattern is valid
type AzurermServiceFabricClusterInvalidUpgradeModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermServiceFabricClusterInvalidUpgradeModeRule returns new rule with default attributes
func NewAzurermServiceFabricClusterInvalidUpgradeModeRule() *AzurermServiceFabricClusterInvalidUpgradeModeRule {
	return &AzurermServiceFabricClusterInvalidUpgradeModeRule{
		resourceType:  "azurerm_service_fabric_cluster",
		attributeName: "upgrade_mode",
		enum: []string{
			"Automatic",
			"Manual",
		},
	}
}

// Name returns the rule name
func (r *AzurermServiceFabricClusterInvalidUpgradeModeRule) Name() string {
	return "azurerm_service_fabric_cluster_invalid_upgrade_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermServiceFabricClusterInvalidUpgradeModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermServiceFabricClusterInvalidUpgradeModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermServiceFabricClusterInvalidUpgradeModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermServiceFabricClusterInvalidUpgradeModeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as upgrade_mode`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
