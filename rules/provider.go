package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/rules/apispec"
)

// Rules is a list of all rules
var Rules = append([]tflint.Rule{
	NewAzurermAppServiceAppInsightsHiddenLinkRule(),
	NewAzurermKubernetesClusterDefaultNodePoolInvalidVMSizeRule(),
	NewAzurermKubernetesClusterNodePoolInvalidVMSizeRule(),
	NewAzurermLinuxVirtualMachineInvalidSizeRule(),
	NewAzurermLinuxVirtualMachineScaleSetInvalidSkuRule(),
	NewAzurermVirtualMachineInvalidVMSizeRule(),
	NewAzurermWindowsVirtualMachineInvalidNameRule(),
	NewAzurermWindowsVirtualMachineInvalidSizeRule(),
	NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule(),
	NewAzurermResourceMissingTagsRule(),
	NewAzurermResourcesMissingPreventDestroyRule(),
	NewAzurermLinuxVirtualMachineInvalidNameRule(),
}, apispec.Rules...)
