package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermKubernetesClusterNodePoolInvalidVMSizeRule checks the pattern is valid
type AzurermKubernetesClusterNodePoolInvalidVMSizeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermKubernetesClusterNodePoolInvalidVMSizeRule returns new rule with default attributes
func NewAzurermKubernetesClusterNodePoolInvalidVMSizeRule() *AzurermKubernetesClusterNodePoolInvalidVMSizeRule {
	return &AzurermKubernetesClusterNodePoolInvalidVMSizeRule{
		resourceType:  "azurerm_kubernetes_cluster_node_pool",
		attributeName: "vm_size",
		enum:          validMachineSizes,
	}
}

// Name returns the rule name
func (r *AzurermKubernetesClusterNodePoolInvalidVMSizeRule) Name() string {
	return "azurerm_kubernetes_cluster_node_pool_invalid_vm_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermKubernetesClusterNodePoolInvalidVMSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermKubernetesClusterNodePoolInvalidVMSizeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermKubernetesClusterNodePoolInvalidVMSizeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermKubernetesClusterNodePoolInvalidVMSizeRule) Check(runner tflint.Runner) error {
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
		err := runner.EvaluateExpr(attribute.Expr, func(val string) error {
			found := false
			for _, item := range r.enum {
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
