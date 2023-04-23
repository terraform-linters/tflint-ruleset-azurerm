// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermFirewallNetworkRuleCollectionInvalidPriorityRule checks the pattern is valid
type AzurermFirewallNetworkRuleCollectionInvalidPriorityRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAzurermFirewallNetworkRuleCollectionInvalidPriorityRule returns new rule with default attributes
func NewAzurermFirewallNetworkRuleCollectionInvalidPriorityRule() *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule {
	return &AzurermFirewallNetworkRuleCollectionInvalidPriorityRule{
		resourceType:  "azurerm_firewall_network_rule_collection",
		attributeName: "priority",
		max:           65000,
		min:           100,
	}
}

// Name returns the rule name
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Name() string {
	return "azurerm_firewall_network_rule_collection_invalid_priority"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermFirewallNetworkRuleCollectionInvalidPriorityRule) Check(runner tflint.Runner) error {
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
		err := runner.EvaluateExpr(attribute.Expr, func (val int) error {
			if val > r.max {
				runner.EmitIssue(
					r,
					"priority must be 65000 or less",
					attribute.Expr.Range(),
				)
			}
			if val < r.min {
				runner.EmitIssue(
					r,
					"priority must be 100 or higher",
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
