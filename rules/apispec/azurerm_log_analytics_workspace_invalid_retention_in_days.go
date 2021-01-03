// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule checks the pattern is valid
type AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule returns new rule with default attributes
func NewAzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule() *AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule {
	return &AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule{
		resourceType:  "azurerm_log_analytics_workspace",
		attributeName: "retention_in_days",
		max:           730,
		min:           30,
	}
}

// Name returns the rule name
func (r *AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule) Name() string {
	return "azurerm_log_analytics_workspace_invalid_retention_in_days"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermLogAnalyticsWorkspaceInvalidRetentionInDaysRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssueOnExpr(
					r,
					"retention_in_days must be 730 or less",
					attribute.Expr,
				)
			}
			if val < r.min {
				runner.EmitIssueOnExpr(
					r,
					"retention_in_days must be 30 or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
