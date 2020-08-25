package rules

import (
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/rules/apispec"
)

// Rules is a list of all rules
var Rules = append([]tflint.Rule{
	NewAzurermLinuxVirtualMachineInvalidSizeRule(),
	NewAzurermLinuxVirtualMachineScaleSetInvalidSkuRule(),
	NewAzurermVirtualMachineInvalidVMSizeRule(),
	NewAzurermWindowsVirtualMachineInvalidSizeRule(),
	NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule(),
}, apispec.Rules...)

func truncateLongMessage(str string) string {
	limit := 80

	str = strings.Replace(str, "\r\n", "\n", -1)
	str = strings.Replace(str, "\n", "\\n", -1)

	r := []rune(str)
	if len(r) > limit {
		return string(r[0:limit]) + "..."
	}

	return str
}
