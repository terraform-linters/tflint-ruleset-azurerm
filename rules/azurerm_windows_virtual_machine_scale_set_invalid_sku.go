package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsVirtualMachineScaleSetInvalidSkuRule checks the pattern is valid
type AzurermWindowsVirtualMachineScaleSetInvalidSkuRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule() *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule {
	return &AzurermWindowsVirtualMachineScaleSetInvalidSkuRule{
		resourceType:  "azurerm_windows_virtual_machine_scale_set",
		attributeName: "sku",
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Name() string {
	return "azurerm_windows_virtual_machine_scale_set_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			found := false
			for _, item := range validMachineSizes {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku`, val),
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
