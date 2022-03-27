package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsVirtualMachineInvalidSizeRule checks the pattern is valid
type AzurermWindowsVirtualMachineInvalidSizeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewAzurermWindowsVirtualMachineInvalidSizeRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineInvalidSizeRule() *AzurermWindowsVirtualMachineInvalidSizeRule {
	return &AzurermWindowsVirtualMachineInvalidSizeRule{
		resourceType:  "azurerm_windows_virtual_machine",
		attributeName: "size",
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Name() string {
	return "azurerm_windows_virtual_machine_invalid_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermWindowsVirtualMachineInvalidSizeRule) Check(runner tflint.Runner) error {
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

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range validMachineSizes {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as size`, val),
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
