// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermTemplateDeploymentInvalidDeploymentModeRule checks the pattern is valid
type AzurermTemplateDeploymentInvalidDeploymentModeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermTemplateDeploymentInvalidDeploymentModeRule returns new rule with default attributes
func NewAzurermTemplateDeploymentInvalidDeploymentModeRule() *AzurermTemplateDeploymentInvalidDeploymentModeRule {
	return &AzurermTemplateDeploymentInvalidDeploymentModeRule{
		resourceType:  "azurerm_template_deployment",
		attributeName: "deployment_mode",
		enum: []string{
			"Incremental",
			"Complete",
		},
	}
}

// Name returns the rule name
func (r *AzurermTemplateDeploymentInvalidDeploymentModeRule) Name() string {
	return "azurerm_template_deployment_invalid_deployment_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermTemplateDeploymentInvalidDeploymentModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermTemplateDeploymentInvalidDeploymentModeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermTemplateDeploymentInvalidDeploymentModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermTemplateDeploymentInvalidDeploymentModeRule) Check(runner tflint.Runner) error {
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
		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as deployment_mode`, truncateLongMessage(val)),
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
