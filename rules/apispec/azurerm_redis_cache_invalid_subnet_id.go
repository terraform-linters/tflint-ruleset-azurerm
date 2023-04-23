// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermRedisCacheInvalidSubnetIDRule checks the pattern is valid
type AzurermRedisCacheInvalidSubnetIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermRedisCacheInvalidSubnetIDRule returns new rule with default attributes
func NewAzurermRedisCacheInvalidSubnetIDRule() *AzurermRedisCacheInvalidSubnetIDRule {
	return &AzurermRedisCacheInvalidSubnetIDRule{
		resourceType:  "azurerm_redis_cache",
		attributeName: "subnet_id",
		pattern:       regexp.MustCompile(`^/subscriptions/[^/]*/resourceGroups/[^/]*/providers/Microsoft.(ClassicNetwork|Network)/virtualNetworks/[^/]*/subnets/[^/]*$`),
	}
}

// Name returns the rule name
func (r *AzurermRedisCacheInvalidSubnetIDRule) Name() string {
	return "azurerm_redis_cache_invalid_subnet_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermRedisCacheInvalidSubnetIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermRedisCacheInvalidSubnetIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermRedisCacheInvalidSubnetIDRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermRedisCacheInvalidSubnetIDRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^/subscriptions/[^/]*/resourceGroups/[^/]*/providers/Microsoft.(ClassicNetwork|Network)/virtualNetworks/[^/]*/subnets/[^/]*$`),
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
