// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermCosmosdbAccountInvalidOfferTypeRule checks the pattern is valid
type AzurermCosmosdbAccountInvalidOfferTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermCosmosdbAccountInvalidOfferTypeRule returns new rule with default attributes
func NewAzurermCosmosdbAccountInvalidOfferTypeRule() *AzurermCosmosdbAccountInvalidOfferTypeRule {
	return &AzurermCosmosdbAccountInvalidOfferTypeRule{
		resourceType:  "azurerm_cosmosdb_account",
		attributeName: "offer_type",
		enum: []string{
			"Standard",
		},
	}
}

// Name returns the rule name
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Name() string {
	return "azurerm_cosmosdb_account_invalid_offer_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermCosmosdbAccountInvalidOfferTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as offer_type`, truncateLongMessage(val)),
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
