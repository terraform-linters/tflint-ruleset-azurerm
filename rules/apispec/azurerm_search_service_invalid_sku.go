// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermSearchServiceInvalidSkuRule checks the pattern is valid
type AzurermSearchServiceInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermSearchServiceInvalidSkuRule returns new rule with default attributes
func NewAzurermSearchServiceInvalidSkuRule() *AzurermSearchServiceInvalidSkuRule {
	return &AzurermSearchServiceInvalidSkuRule{
		resourceType:  "azurerm_search_service",
		attributeName: "sku",
		enum: []string{
			"free",
			"basic",
			"standard",
			"standard2",
			"standard3",
			"storage_optimized_l1",
			"storage_optimized_l2",
		},
	}
}

// Name returns the rule name
func (r *AzurermSearchServiceInvalidSkuRule) Name() string {
	return "azurerm_search_service_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSearchServiceInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSearchServiceInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSearchServiceInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermSearchServiceInvalidSkuRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as sku`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
