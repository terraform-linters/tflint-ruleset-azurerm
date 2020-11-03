package rules

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermSubnetInvalidAddressPrefixRule checks the pattern is valid
type AzurermSubnetInvalidAddressPrefixRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermSubnetInvalidAddressPrefixRule returns new rule with default attributes
func NewAzurermSubnetInvalidAddressPrefixRule() *AzurermSubnetInvalidAddressPrefixRule {
	return &AzurermSubnetInvalidAddressPrefixRule{
		resourceType:  "azurerm_subnet",
		attributeName: "address_prefixes",
		pattern:       regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))?$`),
	}
}

// Name returns the rule name
func (r *AzurermSubnetInvalidAddressPrefixRule) Name() string {
	return "azurerm_subnet_invalid_address"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSubnetInvalidAddressPrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSubnetInvalidAddressPrefixRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSubnetInvalidAddressPrefixRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermSubnetInvalidAddressPrefixRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val []string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val[0]) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is not a valid CIDR`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
