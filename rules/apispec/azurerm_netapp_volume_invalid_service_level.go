// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermNetappVolumeInvalidServiceLevelRule checks the pattern is valid
type AzurermNetappVolumeInvalidServiceLevelRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermNetappVolumeInvalidServiceLevelRule returns new rule with default attributes
func NewAzurermNetappVolumeInvalidServiceLevelRule() *AzurermNetappVolumeInvalidServiceLevelRule {
	return &AzurermNetappVolumeInvalidServiceLevelRule{
		resourceType:  "azurerm_netapp_volume",
		attributeName: "service_level",
		enum: []string{
			"Standard",
			"Premium",
			"Ultra",
			"StandardZRS",
		},
	}
}

// Name returns the rule name
func (r *AzurermNetappVolumeInvalidServiceLevelRule) Name() string {
	return "azurerm_netapp_volume_invalid_service_level"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermNetappVolumeInvalidServiceLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermNetappVolumeInvalidServiceLevelRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermNetappVolumeInvalidServiceLevelRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermNetappVolumeInvalidServiceLevelRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as service_level`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
