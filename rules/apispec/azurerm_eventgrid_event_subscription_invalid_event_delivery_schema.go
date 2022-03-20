// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule checks the pattern is valid
type AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule returns new rule with default attributes
func NewAzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule() *AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule {
	return &AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule{
		resourceType:  "azurerm_eventgrid_event_subscription",
		attributeName: "event_delivery_schema",
		enum: []string{
			"EventGridSchema",
			"CustomInputSchema",
			"CloudEventSchemaV1_0",
		},
	}
}

// Name returns the rule name
func (r *AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule) Name() string {
	return "azurerm_eventgrid_event_subscription_invalid_event_delivery_schema"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermEventgridEventSubscriptionInvalidEventDeliverySchemaRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as event_delivery_schema`, truncateLongMessage(val)),
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
