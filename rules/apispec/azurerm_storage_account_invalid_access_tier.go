// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStorageAccountInvalidAccessTierRule checks the pattern is valid
type AzurermStorageAccountInvalidAccessTierRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermStorageAccountInvalidAccessTierRule returns new rule with default attributes
func NewAzurermStorageAccountInvalidAccessTierRule() *AzurermStorageAccountInvalidAccessTierRule {
	return &AzurermStorageAccountInvalidAccessTierRule{
		resourceType:  "azurerm_storage_account",
		attributeName: "access_tier",
		enum: []string{
			"Hot",
			"Cool",
			"Premium",
			"Cold",
		},
	}
}

// Name returns the rule name
func (r *AzurermStorageAccountInvalidAccessTierRule) Name() string {
	return "azurerm_storage_account_invalid_access_tier"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStorageAccountInvalidAccessTierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStorageAccountInvalidAccessTierRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStorageAccountInvalidAccessTierRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStorageAccountInvalidAccessTierRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as access_tier`, truncateLongMessage(val)),
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
