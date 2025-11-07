package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermWindowsWebAppAppInsightsHiddenLink(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no Application Insights configuration - should pass",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example"
  }
	
  lifecycle {
    ignore_changes = [
      "tags[\"hidden-link: /app-insights-conn-string\"]",
      "tags[\"hidden-link: /app-insights-instrumentation-key\"]",
      "tags[\"hidden-link: /app-insights-resource-id\"]",
    ]
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermWindowsWebAppAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "Application Insights connection string unquoted key without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    APPLICATIONINSIGHTS_CONNECTION_STRING = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermWindowsWebAppAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "Application Insights with lifecycle but no ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {

  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
	
  lifecycle {
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermWindowsWebAppAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "instrumentation key with proper ignore_changes - should pass",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example-key"
  }
	
  lifecycle {
    ignore_changes = [
      "tags[\"hidden-link: /app-insights-conn-string\"]",
      "tags[\"hidden-link: /app-insights-instrumentation-key\"]",
      "tags[\"hidden-link: /app-insights-resource-id\"]",
    ]
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Application Insights with partial ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
	
  lifecycle {
    ignore_changes = [
      "tags[\"hidden-link: /app-insights-conn-string\"]",
    ]
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermWindowsWebAppAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "instrumentation key without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example-key"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermWindowsWebAppAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
	}

	rule := NewAzurermWindowsWebAppAppInsightsHiddenLinkRule()

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
