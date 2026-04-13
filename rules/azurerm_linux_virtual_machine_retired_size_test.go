package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermLinuxVirtualMachineRetiredSize(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "retired size emits issue",
			Content: `
resource "azurerm_linux_virtual_machine" "example" {
  size = "Standard_B1ms"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermLinuxVirtualMachineRetiredSizeRule(),
					Message: `"Standard_B1ms" is a retired or announced-for-retirement Virtual Machine Size`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 25},
					},
				},
			},
		},
		{
			Name: "valid size emits no issue",
			Content: `
resource "azurerm_linux_virtual_machine" "example" {
  size = "Standard_D4s_v5"
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermLinuxVirtualMachineRetiredSizeRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
