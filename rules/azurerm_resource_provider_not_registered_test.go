package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermResourceProviderNotRegistered(t *testing.T) {
	grafana := `
resource "azurerm_dashboard_grafana" "this" {
  grafana_major_version = 11
}`
	grafanaRange := hcl.Range{
		Filename: "resource.tf",
		Start:    hcl.Pos{Line: 2, Column: 1},
		End:      hcl.Pos{Line: 2, Column: 44},
	}
	enabled := `
rule "azurerm_resource_provider_not_registered" {
  enabled = true
}`

	cases := []struct {
		Name     string
		Provider string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "namespace not in registration set",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "core"
  features {}
}`,
			Content: grafana,
			Config:  enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" is not registered by provider "azurerm" (resource_provider_registrations = "core"). Add it to resource_providers_to_register.`,
					Range:   grafanaRange,
				},
			},
		},
		{
			Name: "namespace in resource_providers_to_register",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "core"
  resource_providers_to_register  = ["Microsoft.Dashboard"]
  features {}
}`,
			Content:  grafana,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "namespace with lowercase Azure casing",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "core"
  resource_providers_to_register  = ["microsoft.insights"]
  features {}
}`,
			Content: `
resource "azurerm_monitor_diagnostic_setting" "this" {
  name = "example"
}`,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "namespace with documented casing instead of Azure casing",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "core"
  resource_providers_to_register  = ["Microsoft.Insights"]
  features {}
}`,
			Content: `
resource "azurerm_monitor_diagnostic_setting" "this" {
  name = "example"
}`,
			Config: enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Insights" resource provider in resource_providers_to_register of provider "azurerm" does not match the casing Azure uses, so the azurerm provider skips it. Use "microsoft.insights" instead.`,
					Range: hcl.Range{
						Filename: "provider.tf",
						Start:    hcl.Pos{Line: 4, Column: 37},
						End:      hcl.Pos{Line: 4, Column: 59},
					},
				},
			},
		},
		{
			Name: "missing namespace reported with Azure casing",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "core"
  features {}
}`,
			Content: `
resource "azurerm_monitor_diagnostic_setting" "this" {
  name = "example"
}`,
			Config: enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "microsoft.insights" resource provider required by "azurerm_monitor_diagnostic_setting" is not registered by provider "azurerm" (resource_provider_registrations = "core"). Add it to resource_providers_to_register.`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 53},
					},
				},
			},
		},
		{
			Name: "namespace in ignored_namespaces",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "core"
  features {}
}`,
			Content: grafana,
			Config: `
rule "azurerm_resource_provider_not_registered" {
  enabled            = true
  ignored_namespaces = ["microsoft.dashboard"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "typo in resource_providers_to_register",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "core"
  resource_providers_to_register  = ["Microsoft.Dashboardd"]
  features {}
}`,
			Content: grafana,
			Config:  enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" is not registered by provider "azurerm" (resource_provider_registrations = "core"). Add it to resource_providers_to_register.`,
					Range:   grafanaRange,
				},
			},
		},
		{
			Name: "namespace in registration set",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "all"
  features {}
}`,
			Content:  grafana,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "registration mode defaults to none",
			Provider: `
provider "azurerm" {
  features {}
}`,
			Content: `
resource "azurerm_storage_account" "this" {
  name = "example"
}`,
			Config: enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Storage" resource provider required by "azurerm_storage_account" is not registered by provider "azurerm" (resource_provider_registrations = "none"). Add it to resource_providers_to_register.`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 42},
					},
				},
			},
		},
		{
			Name: "default_registrations configured",
			Provider: `
provider "azurerm" {
  features {}
}`,
			Content: `
resource "azurerm_storage_account" "this" {
  name = "example"
}`,
			Config: `
rule "azurerm_resource_provider_not_registered" {
  enabled               = true
  default_registrations = "core"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "namespace registered by default in Azure",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "none"
  features {}
}`,
			Content: `
resource "azurerm_role_assignment" "this" {
  scope = "/subscriptions/00000000-0000-0000-0000-000000000000"
}`,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "aliased provider",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "all"
  features {}
}

provider "azurerm" {
  alias                           = "other"
  resource_provider_registrations = "core"
  features {}
}`,
			Content: `
resource "azurerm_dashboard_grafana" "this" {
  provider = azurerm.other
}`,
			Config: enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" is not registered by provider "azurerm.other" (resource_provider_registrations = "core"). Add it to resource_providers_to_register.`,
					Range:   grafanaRange,
				},
			},
		},
		{
			Name: "quoted provider reference",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "all"
  features {}
}

provider "azurerm" {
  alias                           = "other"
  resource_provider_registrations = "core"
  features {}
}`,
			Content: `
resource "azurerm_dashboard_grafana" "this" {
  provider = "azurerm.other"
}`,
			Config: enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" is not registered by provider "azurerm.other" (resource_provider_registrations = "core"). Add it to resource_providers_to_register.`,
					Range:   grafanaRange,
				},
			},
		},
		{
			Name: "implicit default provider next to an aliased provider",
			Provider: `
provider "azurerm" {
  alias                           = "other"
  resource_provider_registrations = "all"
  features {}
}`,
			Content: grafana,
			Config:  enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" is not registered by provider "azurerm" (resource_provider_registrations = "none"). Add it to resource_providers_to_register.`,
					Range:   grafanaRange,
				},
			},
		},
		{
			Name: "resource_providers_to_register from variable",
			Provider: `
variable "resource_providers" {
  type    = list(string)
  default = ["Microsoft.Dashboard"]
}

provider "azurerm" {
  resource_provider_registrations = "core"
  resource_providers_to_register  = var.resource_providers
  features {}
}`,
			Content:  grafana,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown resource_providers_to_register",
			Provider: `
variable "resource_providers" {
  type = list(string)
}

provider "azurerm" {
  resource_provider_registrations = "core"
  resource_providers_to_register  = var.resource_providers
  features {}
}`,
			Content:  grafana,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name:     "no azurerm provider configuration",
			Provider: ``,
			Content:  grafana,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "resource without API providers mapping",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "none"
  features {}
}`,
			Content: `
resource "azurerm_key_vault_secret" "this" {
  name = "example"
}`,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "multiple missing namespaces on one resource",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "none"
  features {}
}`,
			Content: `
resource "azurerm_container_app_environment" "this" {
  location = "westeurope"
}`,
			Config: enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.App", "Microsoft.OperationalInsights" resource providers required by "azurerm_container_app_environment" are not registered by provider "azurerm" (resource_provider_registrations = "none"). Add them to resource_providers_to_register.`,
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 52},
					},
				},
			},
		},
		{
			Name: "null resource_provider_registrations and resource_providers_to_register",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = null
  resource_providers_to_register  = null
  features {}
}`,
			Content: grafana,
			Config:  enabled,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceProviderNotRegisteredRule(),
					Message: `The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" is not registered by provider "azurerm" (resource_provider_registrations = "none"). Add it to resource_providers_to_register.`,
					Range:   grafanaRange,
				},
			},
		},
		{
			Name: "unknown resource_provider_registrations",
			Provider: `
variable "mode" {
  type = string
}

provider "azurerm" {
  resource_provider_registrations = var.mode
  features {}
}`,
			Content:  grafana,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "undeclared provider alias",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "none"
  features {}
}`,
			Content: `
resource "azurerm_dashboard_grafana" "this" {
  provider = azurerm.undeclared
}`,
			Config:   enabled,
			Expected: helper.Issues{},
		},
		{
			Name: "provider with another local name",
			Provider: `
provider "azurerm" {
  resource_provider_registrations = "none"
  features {}
}`,
			Content: `
resource "azurerm_dashboard_grafana" "this" {
  provider = other.alias
}`,
			Config:   enabled,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermResourceProviderNotRegisteredRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{
				"provider.tf": tc.Provider,
				"resource.tf": tc.Content,
				".tflint.hcl": tc.Config,
			})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}

func Test_AzurermResourceProviderNotRegistered_InvalidDefaultRegistrations(t *testing.T) {
	runner := helper.TestRunner(t, map[string]string{
		"main.tf": `
provider "azurerm" {
  features {}
}`,
		".tflint.hcl": `
rule "azurerm_resource_provider_not_registered" {
  enabled               = true
  default_registrations = "invalid"
}`,
	})

	err := NewAzurermResourceProviderNotRegisteredRule().Check(runner)
	if err == nil {
		t.Fatal("expected an error for an invalid default_registrations")
	}
	want := `invalid default_registrations "invalid" for rule azurerm_resource_provider_not_registered, must be one of: all, core, extended, legacy, none`
	if err.Error() != want {
		t.Errorf("unexpected error:\n got: %s\nwant: %s", err, want)
	}
}
