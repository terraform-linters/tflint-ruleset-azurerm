// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCdnEndpointInvalidResourceGroupNameRule checks the pattern is valid
type AzurermCdnEndpointInvalidResourceGroupNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermCdnEndpointInvalidResourceGroupNameRule returns new rule with default attributes
func NewAzurermCdnEndpointInvalidResourceGroupNameRule() *AzurermCdnEndpointInvalidResourceGroupNameRule {
	return &AzurermCdnEndpointInvalidResourceGroupNameRule{
		resourceType:  "azurerm_cdn_endpoint",
		attributeName: "resource_group_name",
		pattern:       regexp.MustCompile(`^[-\w\._\(\)]+$`),
	}
}

// Name returns the rule name
func (r *AzurermCdnEndpointInvalidResourceGroupNameRule) Name() string {
	return "azurerm_cdn_endpoint_invalid_resource_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCdnEndpointInvalidResourceGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCdnEndpointInvalidResourceGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCdnEndpointInvalidResourceGroupNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCdnEndpointInvalidResourceGroupNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[-\w\._\(\)]+$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
