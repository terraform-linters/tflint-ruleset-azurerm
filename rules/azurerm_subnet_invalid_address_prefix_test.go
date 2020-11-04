package rules

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermSubnetAddressPrefixRule(t *testing.T) {
	for _, test := range testCases {
		testDisplay := fmt.Sprintf("%d - %s", test.testID, test.testName)
		t.Run(testDisplay, func(t *testing.T) {
			//arrange
			rule := *(NewAzurermSubnetInvalidAddressPrefixRule())
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

var testCases = []struct {
	testID   int
	testName string
	hcl      string
	expected helper.Issues
}{
	{
		testID:   1,
		testName: "Invalid CIDR",
		hcl: `
			resource "azurerm_subnet" "test" {
				address_prefix     = "invalid_cidr"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermSubnetInvalidAddressPrefixRule(),
				Message: "\"invalid_cidr\" is not a valid CIDR",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 26},
					End:      hcl.Pos{Line: 3, Column: 40},
				},
			},
		},
	},
	{
		testID:   2,
		testName: "Invalid IP Range",
		hcl: `
			resource "azurerm_subnet" "test" {
				address_prefix     = "20000.0.10.10/24"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermSubnetInvalidAddressPrefixRule(),
				Message: "\"20000.0.10.10/24\" is not a valid CIDR",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 26},
					End:      hcl.Pos{Line: 3, Column: 44},
				},
			},
		},
	},
	{
		testID:   3,
		testName: "Invalid Bits",
		hcl: `
			resource "azurerm_subnet" "test" {
				address_prefix     = "10.0.10.10/86"
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermSubnetInvalidAddressPrefixRule(),
				Message: "\"10.0.10.10/86\" is not a valid CIDR",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 26},
					End:      hcl.Pos{Line: 3, Column: 41},
				},
			},
		},
	},
	{
		testID:   3,
		testName: "Valid CIDR",
		hcl: `
			resource "azurerm_subnet" "test" {
				address_prefixe    = "10.0.1.0/24"
			}
		`,
		expected: helper.Issues{},
	},
}
