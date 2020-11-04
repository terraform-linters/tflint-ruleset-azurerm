package rules

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermSubnetInvalidAddressPrefixesRule checks the pattern is valid
type AzurermSubnetInvalidAddressPrefixesRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// AddressPrefixes holds the value from the hcl
type AddressPrefixes struct {
	Data string
}

// NewAzurermSubnetInvalidAddressPrefixesRule returns new rule with default attributes
func NewAzurermSubnetInvalidAddressPrefixesRule() *AzurermSubnetInvalidAddressPrefixesRule {
	return &AzurermSubnetInvalidAddressPrefixesRule{
		resourceType:  "azurerm_subnet",
		attributeName: "address_prefixes",
		pattern:       regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))?$`),
	}
}

// Name returns the rule name
func (r *AzurermSubnetInvalidAddressPrefixesRule) Name() string {
	return "azurerm_subnet_invalid_address_prefixes"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSubnetInvalidAddressPrefixesRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSubnetInvalidAddressPrefixesRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSubnetInvalidAddressPrefixesRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func getAddressPrefixes(attribute *hcl.Attribute) ([]string, error) {
	var result []string
	var ctx hcl.EvalContext
	value, _ := attribute.Expr.Value(&ctx)
	for _, item := range value.AsValueSlice() {
		addressPrefix := item.AsString()
		result = append(result, addressPrefix)
	}
	return result, nil
}

// Check checks the pattern is valid
func (r *AzurermSubnetInvalidAddressPrefixesRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		val, err := getAddressPrefixes(attribute)

		return runner.EnsureNoError(err, func() error {
			for _, addressPrefix := range val {
				if !r.pattern.MatchString(addressPrefix) {
					runner.EmitIssueOnExpr(
						r,
						fmt.Sprintf(`"%s" is not a valid CIDR`, addressPrefix),
						attribute.Expr,
					)
				}
			}
			return nil
		})
	})
}
