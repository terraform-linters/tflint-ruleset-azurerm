package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
	"regexp"
)

// AzurermLinuxVirtualMachineInvalidNameRule checks the hostname is valid
type AzurermLinuxVirtualMachineInvalidNameRule struct {
	tflint.DefaultRule

	pattern               *regexp.Regexp
	resourceType          string
	attributeName         string
	attributeComputerName string
}

// NewAzurermLinuxVirtualMachineInvalidNameRule returns new rule with default attributes
func NewAzurermLinuxVirtualMachineInvalidNameRule() *AzurermLinuxVirtualMachineInvalidNameRule {
	return &AzurermLinuxVirtualMachineInvalidNameRule{
		resourceType:          "azurerm_linux_virtual_machine",
		attributeName:         "name",
		attributeComputerName: "computer_name",
		pattern:               regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-]{0,62}$`),
	}
}

// Name returns the rule name
func (r *AzurermLinuxVirtualMachineInvalidNameRule) Name() string {
	return "azurerm_linux_virtual_machine_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermLinuxVirtualMachineInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermLinuxVirtualMachineInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermLinuxVirtualMachineInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the name is valid
func (r *AzurermLinuxVirtualMachineInvalidNameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
			{Name: r.attributeComputerName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		_, exists := resource.Body.Attributes[r.attributeComputerName]
		if exists {
			continue
		}

		primaryAttribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(primaryAttribute.Expr, func(val string) error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, r.pattern),
					primaryAttribute.Expr.Range(),
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
