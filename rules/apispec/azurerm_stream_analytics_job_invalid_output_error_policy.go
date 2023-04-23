// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule checks the pattern is valid
type AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule returns new rule with default attributes
func NewAzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule() *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule {
	return &AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule{
		resourceType:  "azurerm_stream_analytics_job",
		attributeName: "output_error_policy",
		enum: []string{
			"Stop",
			"Drop",
		},
	}
}

// Name returns the rule name
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Name() string {
	return "azurerm_stream_analytics_job_invalid_output_error_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStreamAnalyticsJobInvalidOutputErrorPolicyRule) Check(runner tflint.Runner) error {
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
		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as output_error_policy`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
