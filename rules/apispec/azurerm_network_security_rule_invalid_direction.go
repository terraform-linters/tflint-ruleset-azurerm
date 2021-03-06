// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermNetworkSecurityRuleInvalidDirectionRule checks the pattern is valid
type AzurermNetworkSecurityRuleInvalidDirectionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermNetworkSecurityRuleInvalidDirectionRule returns new rule with default attributes
func NewAzurermNetworkSecurityRuleInvalidDirectionRule() *AzurermNetworkSecurityRuleInvalidDirectionRule {
	return &AzurermNetworkSecurityRuleInvalidDirectionRule{
		resourceType:  "azurerm_network_security_rule",
		attributeName: "direction",
		enum: []string{
			"Inbound",
			"Outbound",
		},
	}
}

// Name returns the rule name
func (r *AzurermNetworkSecurityRuleInvalidDirectionRule) Name() string {
	return "azurerm_network_security_rule_invalid_direction"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNetworkSecurityRuleInvalidDirectionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNetworkSecurityRuleInvalidDirectionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNetworkSecurityRuleInvalidDirectionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermNetworkSecurityRuleInvalidDirectionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as direction`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
