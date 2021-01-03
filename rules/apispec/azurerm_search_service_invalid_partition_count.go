// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermSearchServiceInvalidPartitionCountRule checks the pattern is valid
type AzurermSearchServiceInvalidPartitionCountRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermSearchServiceInvalidPartitionCountRule returns new rule with default attributes
func NewAzurermSearchServiceInvalidPartitionCountRule() *AzurermSearchServiceInvalidPartitionCountRule {
	return &AzurermSearchServiceInvalidPartitionCountRule{
		resourceType:  "azurerm_search_service",
		attributeName: "partition_count",
		max:           12,
		min:           1,
	}
}

// Name returns the rule name
func (r *AzurermSearchServiceInvalidPartitionCountRule) Name() string {
	return "azurerm_search_service_invalid_partition_count"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSearchServiceInvalidPartitionCountRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSearchServiceInvalidPartitionCountRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSearchServiceInvalidPartitionCountRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermSearchServiceInvalidPartitionCountRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssueOnExpr(
					r,
					"partition_count must be 12 or less",
					attribute.Expr,
				)
			}
			if val < r.min {
				runner.EmitIssueOnExpr(
					r,
					"partition_count must be 1 or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
