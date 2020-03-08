package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermVirtualMachineExampleSize(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "azurerm_virtual_machine" "main" {
    vm_size = "Standard_DS1_v2"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermVirtualMachineExampleSizeRule(),
					Message: "VM size is Standard_DS1_v2",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 15},
						End:      hcl.Pos{Line: 3, Column: 32},
					},
				},
			},
		},
	}

	rule := NewAzurermVirtualMachineExampleSizeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
