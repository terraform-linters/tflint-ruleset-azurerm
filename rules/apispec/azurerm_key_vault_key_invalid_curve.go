// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermKeyVaultKeyInvalidCurveRule checks the pattern is valid
type AzurermKeyVaultKeyInvalidCurveRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermKeyVaultKeyInvalidCurveRule returns new rule with default attributes
func NewAzurermKeyVaultKeyInvalidCurveRule() *AzurermKeyVaultKeyInvalidCurveRule {
	return &AzurermKeyVaultKeyInvalidCurveRule{
		resourceType:  "azurerm_key_vault_key",
		attributeName: "curve",
		enum: []string{
			"P-256",
			"P-384",
			"P-521",
			"P-256K",
		},
	}
}

// Name returns the rule name
func (r *AzurermKeyVaultKeyInvalidCurveRule) Name() string {
	return "azurerm_key_vault_key_invalid_curve"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermKeyVaultKeyInvalidCurveRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermKeyVaultKeyInvalidCurveRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermKeyVaultKeyInvalidCurveRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermKeyVaultKeyInvalidCurveRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as curve`, truncateLongMessage(val)),
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
