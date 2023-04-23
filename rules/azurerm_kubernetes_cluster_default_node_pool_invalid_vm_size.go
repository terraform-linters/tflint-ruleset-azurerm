package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule checks the pattern is valid
type AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule struct {
	tflint.DefaultRule

	resourceType  string
	blockType     string
	attributeName string
	enum          []string
}

// NewAzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule returns new rule with default attributes
func NewAzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule() *AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule {
	return &AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule{
		resourceType:  "azurerm_kubernetes_cluster",
		blockType:     "default_node_pool",
		attributeName: "vm_size",
		enum:          validMachineSizes,
	}
}

// Name returns the rule name
func (r *AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule) Name() string {
	return "azurerm_kubernetes_cluster_default_node_pool_invalid_vm_size"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether ...
func (r *AzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: r.blockType,
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: r.attributeName},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, inner := range resource.Body.Blocks {
			attribute, exists := inner.Body.Attributes[r.attributeName]
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
	}

	return nil
}
