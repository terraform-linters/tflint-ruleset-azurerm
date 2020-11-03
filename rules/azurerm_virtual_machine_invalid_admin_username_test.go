package rules

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermVirtualMachineInvalidAdminUserNameRule(t *testing.T) {
	for _, test := range AzurermVirtualMachineInvalidAdminUserNameRuleTestCases {
		testDisplay := fmt.Sprintf("%d - %s", test.testID, test.testName)
		t.Run(testDisplay, func(t *testing.T) {
			//arrange
			rule := *(NewAzurermVirtualMachineInvalidAdminUserNameRule())
			runner := helper.TestRunner(t, map[string]string{"instances.tf": test.hcl})
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//act
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			//assert
			helper.AssertIssues(t, test.expected, runner.Issues)

		})
	}
}

var AzurermVirtualMachineInvalidAdminUserNameRuleTestCases = []struct {
	testID   int
	testName string
	hcl      string
	expected helper.Issues
}{
	{
		testID:   1,
		testName: "123 is not allowed.",
		hcl: `
			resource "azurerm_virtual_machine" "test" {
				admin_username = "123"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermVirtualMachineInvalidAdminUserNameRule(),
				Message: "\"123\" is not a valid VM Admin username",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 22},
					End:      hcl.Pos{Line: 3, Column: 27},
				},
			},
		},
	},
	{
		testID:   2,
		testName: "root is not allowed.",
		hcl: `
			resource "azurerm_virtual_machine" "test" {
				admin_username = "root"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermVirtualMachineInvalidAdminUserNameRule(),
				Message: "\"root\" is not a valid VM Admin username",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 22},
					End:      hcl.Pos{Line: 3, Column: 28},
				},
			},
		},
	},
	{
		testID:   3,
		testName: "azureUser is not allowed.",
		hcl: `
			resource "azurerm_virtual_machine" "test" {
				admin_username = "AzureUser"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermVirtualMachineInvalidAdminUserNameRule(),
				Message: "\"azureuser\" is not a valid VM Admin username",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 22},
					End:      hcl.Pos{Line: 3, Column: 33},
				},
			},
		},
	},
}
