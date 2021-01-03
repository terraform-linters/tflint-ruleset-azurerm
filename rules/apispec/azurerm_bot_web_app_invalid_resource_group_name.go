// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermBotWebAppInvalidResourceGroupNameRule checks the pattern is valid
type AzurermBotWebAppInvalidResourceGroupNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermBotWebAppInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermBotWebAppInvalidResourceGroupNameRule() *AzurermBotWebAppInvalidResourceGroupNameRule {
	return &AzurermBotWebAppInvalidResourceGroupNameRule{
		resourceType:  "azurerm_bot_web_app",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
	}
}

// Name returns the rule name
func (r *AzurermBotWebAppInvalidResourceGroupNameRule) Name() string {
	return "azurerm_bot_web_app_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBotWebAppInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBotWebAppInvalidResourceGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBotWebAppInvalidResourceGroupNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermBotWebAppInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
