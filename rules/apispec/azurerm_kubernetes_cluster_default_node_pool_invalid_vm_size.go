package apispec

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
		enum: []string{
			"Standard_A1",
			"Standard_A10",
			"Standard_A11",
			"Standard_A1_v2",
			"Standard_A2",
			"Standard_A2_v2",
			"Standard_A2m_v2",
			"Standard_A3",
			"Standard_A4",
			"Standard_A4_v2",
			"Standard_A4m_v2",
			"Standard_A5",
			"Standard_A6",
			"Standard_A7",
			"Standard_A8",
			"Standard_A8_v2",
			"Standard_A8m_v2",
			"Standard_A9",
			"Standard_B2ms",
			"Standard_B2s",
			"Standard_B4ms",
			"Standard_B8ms",
			"Standard_D1",
			"Standard_D11",
			"Standard_D11_v2",
			"Standard_D11_v2_Promo",
			"Standard_D12",
			"Standard_D12_v2",
			"Standard_D12_v2_Promo",
			"Standard_D13",
			"Standard_D13_v2",
			"Standard_D13_v2_Promo",
			"Standard_D14",
			"Standard_D14_v2",
			"Standard_D14_v2_Promo",
			"Standard_D15_v2",
			"Standard_D16_v3",
			"Standard_D16s_v3",
			"Standard_D1_v2",
			"Standard_D2",
			"Standard_D2_v2",
			"Standard_D2_v2_Promo",
			"Standard_D2_v3",
			"Standard_D2s_v3",
			"Standard_D3",
			"Standard_D32_v3",
			"Standard_D32s_v3",
			"Standard_D3_v2",
			"Standard_D3_v2_Promo",
			"Standard_D4",
			"Standard_D4_v2",
			"Standard_D4_v2_Promo",
			"Standard_D4_v3",
			"Standard_D4s_v3",
			"Standard_D5_v2",
			"Standard_D5_v2_Promo",
			"Standard_D64_v3",
			"Standard_D64s_v3",
			"Standard_D8_v3",
			"Standard_D8s_v3",
			"Standard_DS1",
			"Standard_DS11",
			"Standard_DS11_v2",
			"Standard_DS11_v2_Promo",
			"Standard_DS12",
			"Standard_DS12_v2",
			"Standard_DS12_v2_Promo",
			"Standard_DS13",
			"Standard_DS13-2_v2",
			"Standard_DS13-4_v2",
			"Standard_DS13_v2",
			"Standard_DS13_v2_Promo",
			"Standard_DS14",
			"Standard_DS14-4_v2",
			"Standard_DS14-8_v2",
			"Standard_DS14_v2",
			"Standard_DS14_v2_Promo",
			"Standard_DS15_v2",
			"Standard_DS1_v2",
			"Standard_DS2",
			"Standard_DS2_v2",
			"Standard_DS2_v2_Promo",
			"Standard_DS3",
			"Standard_DS3_v2",
			"Standard_DS3_v2_Promo",
			"Standard_DS4",
			"Standard_DS4_v2",
			"Standard_DS4_v2_Promo",
			"Standard_DS5_v2",
			"Standard_DS5_v2_Promo",
			"Standard_E16_v3",
			"Standard_E16s_v3",
			"Standard_E2_v3",
			"Standard_E2s_v3",
			"Standard_E32-16s_v3",
			"Standard_E32-8s_v3",
			"Standard_E32_v3",
			"Standard_E32s_v3",
			"Standard_E4_v3",
			"Standard_E4s_v3",
			"Standard_E64-16s_v3",
			"Standard_E64-32s_v3",
			"Standard_E64_v3",
			"Standard_E64s_v3",
			"Standard_E8_v3",
			"Standard_E8s_v3",
			"Standard_F1",
			"Standard_F16",
			"Standard_F16s",
			"Standard_F16s_v2",
			"Standard_F1s",
			"Standard_F2",
			"Standard_F2s",
			"Standard_F2s_v2",
			"Standard_F32s_v2",
			"Standard_F4",
			"Standard_F4s",
			"Standard_F4s_v2",
			"Standard_F64s_v2",
			"Standard_F72s_v2",
			"Standard_F8",
			"Standard_F8s",
			"Standard_F8s_v2",
			"Standard_G1",
			"Standard_G2",
			"Standard_G3",
			"Standard_G4",
			"Standard_G5",
			"Standard_GS1",
			"Standard_GS2",
			"Standard_GS3",
			"Standard_GS4",
			"Standard_GS4-4",
			"Standard_GS4-8",
			"Standard_GS5",
			"Standard_GS5-16",
			"Standard_GS5-8",
			"Standard_H16",
			"Standard_H16m",
			"Standard_H16mr",
			"Standard_H16r",
			"Standard_H8",
			"Standard_H8m",
			"Standard_L16s",
			"Standard_L32s",
			"Standard_L4s",
			"Standard_L8s",
			"Standard_M128-32ms",
			"Standard_M128-64ms",
			"Standard_M128ms",
			"Standard_M128s",
			"Standard_M64-16ms",
			"Standard_M64-32ms",
			"Standard_M64ms",
			"Standard_M64s",
			"Standard_NC12",
			"Standard_NC12s_v2",
			"Standard_NC12s_v3",
			"Standard_NC24",
			"Standard_NC24r",
			"Standard_NC24rs_v2",
			"Standard_NC24rs_v3",
			"Standard_NC24s_v2",
			"Standard_NC24s_v3",
			"Standard_NC6",
			"Standard_NC6s_v2",
			"Standard_NC6s_v3",
			"Standard_ND12s",
			"Standard_ND24rs",
			"Standard_ND24s",
			"Standard_ND6s",
			"Standard_NV12",
			"Standard_NV24",
			"Standard_NV6",
		},
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
						fmt.Sprintf(`"%s" is an invalid value as vm_size`, truncateLongMessage(val)),
						attribute.Expr.Range(),
					)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
