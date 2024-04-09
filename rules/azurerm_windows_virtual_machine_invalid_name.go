package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
	"regexp"
)

// AzurermWindowsVirtualMachineInvalidNameRule checks the hostname is valid
type AzurermWindowsVirtualMachineInvalidNameRule struct {
	tflint.DefaultRule

	pattern                *regexp.Regexp
	resourceType           string
	primaryAttributeName   string
	secondaryAttributeName string
}

// NewAzurermWindowsVirtualMachineInvalidNameRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineInvalidNameRule() *AzurermWindowsVirtualMachineInvalidNameRule {
	return &AzurermWindowsVirtualMachineInvalidNameRule{
		resourceType:           "azurerm_windows_virtual_machine",
		primaryAttributeName:   "name",
		secondaryAttributeName: "computer_name",
		pattern:                regexp.MustCompile(`^[a-zA-Z0-9]{0,1}[a-zA-Z0-9-]{0,13}[a-zA-Z0-9]$`),
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineInvalidNameRule) Name() string {
	return "azurerm_windows_virtual_machine_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineInvalidNameRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks the name is valid
func (r *AzurermWindowsVirtualMachineInvalidNameRule) Check(runner tflint.Runner) error {
	// TODO: Write the implementation here. See this documentation for what tflint.Runner can do.
	//       https://pkg.go.dev/github.com/terraform-linters/tflint-plugin-sdk/tflint#Runner

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.primaryAttributeName},
			{Name: r.secondaryAttributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		_, exists := resource.Body.Attributes[r.secondaryAttributeName]
		if exists {
			continue
		}

		primaryAttribute, exists := resource.Body.Attributes[r.primaryAttributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(primaryAttribute.Expr, func(val string) error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, val, `^[a-zA-Z0-9]{0,1}[a-zA-Z0-9]{0,13}[a-zA-Z0-9]$`),
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
