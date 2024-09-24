package rules

import (
	// to print actual issues for debugging
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermCosmosdbSqlContainerDeprecatedPartitionKeyPath(t *testing.T) {
    cases := []struct {
        Name     string
        Content  string
        Expected helper.Issues
    }{
        {
            Name: "Deprecated partition_key_path",
            Content: `
resource "azurerm_cosmosdb_sql_container" "example" {
  partition_key_path = "/deprecated_key"
}`,
            Expected: helper.Issues{
                {
                    Rule:    NewAzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule(),
                    Message: "`partition_key_path` is deprecated and should be replaced with `partition_key_paths`.",
                    Range: hcl.Range{
                        Filename: "resource.tf",
                        Start:    hcl.Pos{Line: 3, Column: 24},
                        End:      hcl.Pos{Line: 3, Column: 41},
                    },
                },
            },
        },
        {
            Name: "Valid partition_key_paths",
            Content: `
resource "azurerm_cosmosdb_sql_container" "example" {
  partition_key_paths = ["/valid_key"]
}`,
            Expected: helper.Issues{}, // No issues since the correct attribute is used
        },
    }

    rule := NewAzurermCosmosdbSqlContainerDeprecatedPartitionKeyPathRule()

    for _, tc := range cases {
        runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

        if err := rule.Check(runner); err != nil {
            t.Fatalf("Unexpected error occurred: %s", err)
        }

        helper.AssertIssues(t, tc.Expected, runner.Issues)
    }
}
