// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermContainerRegistryWebhookInvalidNameRule checks the pattern is valid
type AzurermContainerRegistryWebhookInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermContainerRegistryWebhookInvalidNameRule returns new rule with default attributes
func NewAzurermContainerRegistryWebhookInvalidNameRule() *AzurermContainerRegistryWebhookInvalidNameRule {
	return &AzurermContainerRegistryWebhookInvalidNameRule{
		resourceType:  "azurerm_container_registry_webhook",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9]*$`),
	}
}

// Name returns the rule name
func (r *AzurermContainerRegistryWebhookInvalidNameRule) Name() string {
	return "azurerm_container_registry_webhook_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermContainerRegistryWebhookInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermContainerRegistryWebhookInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermContainerRegistryWebhookInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermContainerRegistryWebhookInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
