// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermApplicationInsightsWebTestInvalidKindRule checks the pattern is valid
type AzurermApplicationInsightsWebTestInvalidKindRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermApplicationInsightsWebTestInvalidKindRule returns new rule with default attributes
func NewAzurermApplicationInsightsWebTestInvalidKindRule() *AzurermApplicationInsightsWebTestInvalidKindRule {
	return &AzurermApplicationInsightsWebTestInvalidKindRule{
		resourceType:  "azurerm_application_insights_web_test",
		attributeName: "kind",
		enum: []string{
			"ping",
			"multistep",
		},
	}
}

// Name returns the rule name
func (r *AzurermApplicationInsightsWebTestInvalidKindRule) Name() string {
	return "azurerm_application_insights_web_test_invalid_kind"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermApplicationInsightsWebTestInvalidKindRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermApplicationInsightsWebTestInvalidKindRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermApplicationInsightsWebTestInvalidKindRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermApplicationInsightsWebTestInvalidKindRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as kind`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
