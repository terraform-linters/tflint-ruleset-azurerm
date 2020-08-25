package rules

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStorageAccountInvalidNameRule checks the pattern is valid
type AzurermStorageAccountInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermStorageAccountInvalidNameRule returns new rule with default attributes
func NewAzurermStorageAccountInvalidNameRule() *AzurermStorageAccountInvalidNameRule {
	return &AzurermStorageAccountInvalidNameRule{
		resourceType:  "azurerm_storage_account",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-z0-9]{3,24}$`),
	}
}

// Name returns the rule name
func (r *AzurermStorageAccountInvalidNameRule) Name() string {
	return "azurerm_storage_account_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStorageAccountInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStorageAccountInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStorageAccountInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStorageAccountInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z0-9]{3,24}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
