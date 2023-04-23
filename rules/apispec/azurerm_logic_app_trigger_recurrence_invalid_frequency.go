// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule checks the pattern is valid
type AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermLogicAppTriggerRecurrenceInvalidFrequencyRule returns new rule with default attributes
func NewAzurermLogicAppTriggerRecurrenceInvalidFrequencyRule() *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule {
	return &AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule{
		resourceType:  "azurerm_logic_app_trigger_recurrence",
		attributeName: "frequency",
		enum: []string{
			"NotSpecified",
			"Second",
			"Minute",
			"Hour",
			"Day",
			"Week",
			"Month",
			"Year",
		},
	}
}

// Name returns the rule name
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Name() string {
	return "azurerm_logic_app_trigger_recurrence_invalid_frequency"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermLogicAppTriggerRecurrenceInvalidFrequencyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as frequency`, truncateLongMessage(val)),
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
