package apispec

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_generatedEnumRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid",
			Content: `
resource "azurerm_analysis_services_server" "main" {
    querypool_connection_mode = "Write"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule(),
					Message: `"Write" is an invalid value as querypool_connection_mode`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 33},
						End:      hcl.Pos{Line: 3, Column: 40},
					},
				},
			},
		},
		{
			Name: "valid",
			Content: `
resource "azurerm_analysis_services_server" "main" {
    querypool_connection_mode = "ReadOnly"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermAnalysisServicesServerInvalidQuerypoolConnectionModeRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func Test_generatedPatternRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid",
			Content: `
resource "azurerm_mysql_firewall_rule" "main" {
    start_ip_address = "192.168.0.256"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermMysqlFirewallRuleInvalidStartIPAddressRule(),
					Message: `"192.168.0.256" does not match valid pattern ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 24},
						End:      hcl.Pos{Line: 3, Column: 39},
					},
				},
			},
		},
		{
			Name: "valid",
			Content: `
resource "azurerm_mysql_firewall_rule" "main" {
    start_ip_address = "192.168.0.1"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermMysqlFirewallRuleInvalidStartIPAddressRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func Test_generatedMinimumRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid",
			Content: `
resource "azurerm_search_service" "main" {
    partition_count = -1
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermSearchServiceInvalidPartitionCountRule(),
					Message: `partition_count must be 1 or higher`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 23},
						End:      hcl.Pos{Line: 3, Column: 25},
					},
				},
			},
		},
		{
			Name: "valid",
			Content: `
resource "azurerm_search_service" "main" {
    partition_count = 1
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermSearchServiceInvalidPartitionCountRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func Test_generatedMaximumRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid",
			Content: `
resource "azurerm_search_service" "main" {
    partition_count = 100
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermSearchServiceInvalidPartitionCountRule(),
					Message: `partition_count must be 12 or less`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 23},
						End:      hcl.Pos{Line: 3, Column: 26},
					},
				},
			},
		},
		{
			Name: "valid",
			Content: `
resource "azurerm_search_service" "main" {
    partition_count = 10
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermSearchServiceInvalidPartitionCountRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
