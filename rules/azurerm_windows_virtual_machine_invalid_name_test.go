package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermWindowsVirtualMachineInvalidName(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "InvalidName - Ends with hyphen",
			Content: `
resource "azurerm_windows_virtual_machine" "vm" {
  name = "dummy-"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermWindowsVirtualMachineInvalidNameRule(),
					Message: `"dummy-" does not match valid pattern ^[a-zA-Z0-9]{0,1}[a-zA-Z0-9-]{0,13}[a-zA-Z0-9]$`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 18},
					},
				},
			},
		},
		{
			Name: "InvalidName - too long",
			Content: `
resource "azurerm_windows_virtual_machine" "vm" {
  name = "dummyhostname123"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermWindowsVirtualMachineInvalidNameRule(),
					Message: `"dummyhostname123" does not match valid pattern ^[a-zA-Z0-9]{0,1}[a-zA-Z0-9-]{0,13}[a-zA-Z0-9]$`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 28},
					},
				},
			},
		},
		{
			Name: "ValidName - Name is valid",
			Content: `
resource "azurerm_windows_virtual_machine" "vm" {
  name = "dummyhostname1"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "ValidName - Hostname is specified",
			Content: `
resource "azurerm_windows_virtual_machine" "vm" {
  name = "dummy-"
  computer_name = "dummyhostname1"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermWindowsVirtualMachineInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
