package rules

import (
	"fmt"
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermSubnetAddressPrefixesRule(t *testing.T) {
	for _, test := range AzurermSubnetInvalidAddressPrefixesTestCases {
		testDisplay := fmt.Sprintf("%d - %s", test.testID, test.testName)
		t.Run(testDisplay, func(t *testing.T) {
			//arrange
			rule := *(NewAzurermSubnetInvalidAddressPrefixesRule())
			runner := helper.TestRunner(t, map[string]string{"instances.tf": test.hcl})

			//act
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			//assert
			helper.AssertIssues(t, test.expected, runner.Issues)

		})
	}
}

var AzurermSubnetInvalidAddressPrefixesTestCases = []struct {
	testID   int
	testName string
	hcl      string
	expected helper.Issues
}{
	{
		testID:   1,
		testName: "Invalid CIDR, one item in list",
		hcl: `
			resource "azurerm_subnet" "test" {
				address_prefixes     = ["invalid_cidr","10.0.1.0/24"]
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermSubnetInvalidAddressPrefixesRule(),
				Message: "\"invalid_cidr\" is not a valid CIDR",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 28},
					End:      hcl.Pos{Line: 3, Column: 58},
				},
			},
		},
	},
}
