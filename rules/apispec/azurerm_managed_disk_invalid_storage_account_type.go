// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermManagedDiskInvalidStorageAccountTypeRule checks the pattern is valid
type AzurermManagedDiskInvalidStorageAccountTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermManagedDiskInvalidStorageAccountTypeRule returns new rule with default attributes
func NewAzurermManagedDiskInvalidStorageAccountTypeRule() *AzurermManagedDiskInvalidStorageAccountTypeRule {
	return &AzurermManagedDiskInvalidStorageAccountTypeRule{
		resourceType:  "azurerm_managed_disk",
		attributeName: "storage_account_type",
		enum: []string{
			"Standard_LRS",
			"Premium_LRS",
			"StandardSSD_LRS",
			"UltraSSD_LRS",
			"Premium_ZRS",
			"StandardSSD_ZRS",
			"PremiumV2_LRS",
		},
	}
}

// Name returns the rule name
func (r *AzurermManagedDiskInvalidStorageAccountTypeRule) Name() string {
	return "azurerm_managed_disk_invalid_storage_account_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermManagedDiskInvalidStorageAccountTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermManagedDiskInvalidStorageAccountTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermManagedDiskInvalidStorageAccountTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermManagedDiskInvalidStorageAccountTypeRule) Check(runner tflint.Runner) error {
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
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as storage_account_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
