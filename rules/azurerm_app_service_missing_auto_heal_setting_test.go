package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermAppServiceMissingAutoHealSetting(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		// Linux Web App tests
		{
			Name: "Linux Web App - no site_config block - should pass",
			Content: `
resource "azurerm_linux_web_app" "example" {
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Linux Web App - site_config without auto_heal_setting - should fail",
			Content: `
resource "azurerm_linux_web_app" "example" {  
  site_config {
    always_on = true
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceMissingAutoHealSettingRule(),
					Message: "auto_heal_setting should be configured in site_config block for robust app services.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 43},
					},
				},
			},
		},
		{
			Name: "Linux Web App - site_config with auto_heal_setting - should pass",
			Content: `
resource "azurerm_linux_web_app" "example" {
  site_config {
    auto_heal_setting {
      action {
        action_type = "Recycle"
      }
      trigger {
        status_code {
          count             = 5
          interval          = "00:01:00"
          status_code_range = "500-599"
        }
      }
    }
  }
}
`,
			Expected: helper.Issues{},
		},
		// Linux Web App Slot tests
		{
			Name: "Linux Web App Slot - no site_config block - should pass",
			Content: `
resource "azurerm_linux_web_app_slot" "example" {
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Linux Web App Slot - site_config without auto_heal_setting - should fail",
			Content: `
resource "azurerm_linux_web_app_slot" "example" {  
  site_config {
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceMissingAutoHealSettingRule(),
					Message: "auto_heal_setting should be configured in site_config block for robust app services.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 48},
					},
				},
			},
		},
		{
			Name: "Linux Web App Slot - site_config with auto_heal_setting - should pass",
			Content: `
resource "azurerm_linux_web_app_slot" "example" {
  site_config {
    auto_heal_setting {
      action {
        action_type = "Recycle"
      }
      trigger {
        requests {
          count    = 100
          interval = "00:01:00"
        }
      }
    }
  }
}
`,
			Expected: helper.Issues{},
		},
		// Windows Web App tests
		{
			Name: "Windows Web App - no site_config block - should pass",
			Content: `
resource "azurerm_windows_web_app" "example" {
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Windows Web App - site_config without auto_heal_setting - should fail",
			Content: `
resource "azurerm_windows_web_app" "example" {
  site_config {
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceMissingAutoHealSettingRule(),
					Message: "auto_heal_setting should be configured in site_config block for robust app services.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "Windows Web App - site_config with auto_heal_setting - should pass",
			Content: `
resource "azurerm_windows_web_app" "example" {
  site_config {
    auto_heal_setting {
      action {
        action_type = "Recycle"
      }
      trigger {
        slow_request {
          count      = 10
          interval   = "00:01:00"
          time_taken = "00:00:30"
        }
      }
    }
  }
}
`,
			Expected: helper.Issues{},
		},
		// Windows Web App Slot tests
		{
			Name: "Windows Web App Slot - no site_config block - should pass",
			Content: `
resource "azurerm_windows_web_app_slot" "example" {
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "Windows Web App Slot - site_config without auto_heal_setting - should fail",
			Content: `
resource "azurerm_windows_web_app_slot" "example" {
  site_config {
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermAppServiceMissingAutoHealSettingRule(),
					Message: "auto_heal_setting should be configured in site_config block for robust app services.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 50},
					},
				},
			},
		},
		{
			Name: "Windows Web App Slot - site_config with auto_heal_setting - should pass",
			Content: `
resource "azurerm_windows_web_app_slot" "example" {
  site_config {
    auto_heal_setting {
      action {
        action_type = "Recycle"
      }
      trigger {
        status_code {
          count             = 3
          interval          = "00:05:00"
          status_code_range = "500-599"
        }
      }
    }
  }
}
`,
			Expected: helper.Issues{},
		},
		// Multiple triggers
		{
			Name: "Linux Web App - site_config with complex auto_heal_setting - should pass",
			Content: `
resource "azurerm_linux_web_app" "example" {
  site_config {
    auto_heal_setting {
      action {
        action_type = "Recycle"
        minimum_process_execution_time = "00:01:00"
      }
      trigger {
        status_code {
          count             = 5
          interval          = "00:01:00"
          status_code_range = "500-599"
        }
        requests {
          count    = 100
          interval = "00:01:00"
        }
        slow_request {
          count      = 10
          interval   = "00:02:00"
          time_taken = "00:00:45"
        }
      }
    }
  }
}
`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermAppServiceMissingAutoHealSettingRule()

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
