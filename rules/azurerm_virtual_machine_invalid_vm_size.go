package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualMachineInvalidVMSizeRule checks the pattern is valid
type AzurermVirtualMachineInvalidVMSizeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewAzurermVirtualMachineInvalidVMSizeRule returns new rule with default attributes
func NewAzurermVirtualMachineInvalidVMSizeRule() *AzurermVirtualMachineInvalidVMSizeRule {
	return &AzurermVirtualMachineInvalidVMSizeRule{
		resourceType:  "azurerm_virtual_machine",
		attributeName: "vm_size",
	}
}

// Name returns the rule name
func (r *AzurermVirtualMachineInvalidVMSizeRule) Name() string {
	return "azurerm_virtual_machine_invalid_vm_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineInvalidVMSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineInvalidVMSizeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineInvalidVMSizeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualMachineInvalidVMSizeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as vm_size`, val),
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
