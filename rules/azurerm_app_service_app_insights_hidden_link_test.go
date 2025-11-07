package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermAppServiceAppInsightsHiddenLink(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		// Linux Web App tests
		{
			Name: "Linux Web App - no Application Insights configuration - should pass",
			Content: `
resource "azurerm_linux_web_app" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Linux Web App - Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_linux_web_app" "example" {
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
			Name: "Linux Web App - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_linux_web_app" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 43},
					},
				},
			},
		},
		{
			Name: "Linux Web App - instrumentation key without ignore_changes - should fail",
			Content: `
resource "azurerm_linux_web_app" "example" {
  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example-key"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 43},
					},
				},
			},
		},
		// Linux Web App Slot tests
		{
			Name: "Linux Web App Slot - no Application Insights configuration - should pass",
			Content: `
resource "azurerm_linux_web_app_slot" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Linux Web App Slot - Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_linux_web_app_slot" "example" {
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
			Name: "Linux Web App Slot - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_linux_web_app_slot" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 48},
					},
				},
			},
		},
		// Windows Web App tests
		{
			Name: "Windows Web App - no Application Insights configuration - should pass",
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
			Name: "Windows Web App - Application Insights with proper ignore_changes - should pass",
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
			Name: "Windows Web App - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
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
			Name: "Windows Web App - instrumentation key without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {
  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example-key"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		// Windows Web App Slot tests
		{
			Name: "Windows Web App Slot - no Application Insights configuration - should pass",
			Content: `
resource "azurerm_windows_web_app_slot" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Windows Web App Slot - Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_windows_web_app_slot" "example" {
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
			Name: "Windows Web App Slot - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_web_app_slot" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 50},
					},
				},
			},
		},
		// Linux Function App tests
		{
			Name: "Linux Function App - no Application Insights configuration - should pass",
			Content: `
resource "azurerm_linux_function_app" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Linux Function App - Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_linux_function_app" "example" {
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
			Name: "Linux Function App - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_linux_function_app" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 48},
					},
				},
			},
		},
		// Linux Function App Slot tests
		{
			Name: "Linux Function App Slot - no Application Insights configuration - should pass",
			Content: `
resource "azurerm_linux_function_app_slot" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Linux Function App Slot - Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_linux_function_app_slot" "example" {
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
			Name: "Linux Function App Slot - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_linux_function_app_slot" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 53},
					},
				},
			},
		},
		// Windows Function App tests
		{
			Name: "Windows Function App - no Application Insights configuration - should pass",
			Content: `
resource "azurerm_windows_function_app" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Windows Function App - Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_windows_function_app" "example" {
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
			Name: "Windows Function App - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_function_app" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 50},
					},
				},
			},
		},
		// Windows Function App Slot tests
		{
			Name: "Windows Function App Slot - no Application Insights configuration - should pass",
			Content: `
resource "azurerm_windows_function_app_slot" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Windows Function App Slot - Application Insights with proper ignore_changes - should pass",
			Content: `
resource "azurerm_windows_function_app_slot" "example" {
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
			Name: "Windows Function App Slot - Application Insights connection string without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_function_app_slot" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 55},
					},
				},
			},
		},
		// Mixed scenarios
		{
			Name: "Application Insights with partial ignore_changes - should fail",
			Content: `
resource "azurerm_linux_web_app" "example" {
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
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 43},
					},
				},
			},
		},
		// Site Config tests
		{
			Name: "Linux Function App - site_config with application_insights_connection_string without ignore_changes - should fail",
			Content: `
resource "azurerm_linux_function_app" "example" {
  site_config {
    application_insights_connection_string = "example-connection"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 48},
					},
				},
			},
		},
		{
			Name: "Windows Function App - site_config with application_insights_key without ignore_changes - should fail",
			Content: `
resource "azurerm_windows_function_app" "example" {
  site_config {
    application_insights_key = "example-key"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 50},
					},
				},
			},
		},
		{
			Name: "Function App Slot - site_config with empty application_insights_connection_string - should pass",
			Content: `
resource "azurerm_linux_function_app_slot" "example" {
  site_config {
    application_insights_connection_string = ""
  }
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Mixed configuration - app_settings and site_config with Application Insights - should fail",
			Content: `
resource "azurerm_linux_function_app" "example" {
  app_settings = {
    "SOME_OTHER_SETTING" = "value"
  }
  
  site_config {
    application_insights_connection_string = "example-connection"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceAppInsightsHiddenLinkRule(),
					Message: "When Application Insights is configured, lifecycle { ignore_changes } should include all hidden-link tags: tags[\"hidden-link: /app-insights-conn-string\"], tags[\"hidden-link: /app-insights-instrumentation-key\"], tags[\"hidden-link: /app-insights-resource-id\"]",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 48},
					},
				},
			},
		},
	}

	rule := NewAzurermAppServiceAppInsightsHiddenLinkRule()

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
