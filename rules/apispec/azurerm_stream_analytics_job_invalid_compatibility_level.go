// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule checks the pattern is valid
type AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermStreamAnalyticsJobInvalidCompatibilityLevelRule returns new rule with default attributes
func NewAzurermStreamAnalyticsJobInvalidCompatibilityLevelRule() *AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule {
	return &AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule{
		resourceType:  "azurerm_stream_analytics_job",
		attributeName: "compatibility_level",
		enum: []string{
			"1.0",
			"1.2",
		},
	}
}

// Name returns the rule name
func (r *AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule) Name() string {
	return "azurerm_stream_analytics_job_invalid_compatibility_level"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStreamAnalyticsJobInvalidCompatibilityLevelRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as compatibility_level`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
