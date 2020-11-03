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
				address_prefixes     = ["10.0.1.0/asd"]
			}
		`,
		expected: helper.Issues{
			{
				Rule:    NewAzurermSubnetInvalidAddressPrefixRule(),
				Message: "\"NotavalidCIDR\" does not match valid pattern",
				Range: hcl.Range{
					Filename: "instances.tf",
					Start:    hcl.Pos{Line: 3, Column: 12},
					End:      hcl.Pos{Line: 3, Column: 27},
				},
			},
		},
	},
}
