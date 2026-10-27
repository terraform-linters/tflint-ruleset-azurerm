package rules

import (
	"os"
	"path/filepath"
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermResourceProviderNotRegistered_Modules(t *testing.T) {
	const provider = `
provider "azurerm" {
  resource_provider_registrations = "core"
  features {}
}

provider "azurerm" {
  alias                           = "dashboard"
  resource_provider_registrations = "core"
  resource_providers_to_register  = ["Microsoft.Dashboard"]
  features {}
}
`
	const grafana = `
resource "azurerm_dashboard_grafana" "this" {
  grafana_major_version = 11
}
`
	const network = `
resource "azurerm_virtual_network" "this" {
  address_space = ["10.0.0.0/16"]
}
`
	moduleRange := hcl.Range{
		Filename: "main.tf",
		Start:    hcl.Pos{Line: 2, Column: 1},
		End:      hcl.Pos{Line: 2, Column: 13},
	}
	issue := func(address string) helper.Issues {
		return helper.Issues{
			{
				Rule:    NewAzurermResourceProviderNotRegisteredRule(),
				Message: `The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" in ` + address + ` is not registered by provider "azurerm" (resource_provider_registrations = "core"). Add it to resource_providers_to_register.`,
				Range:   moduleRange,
			},
		}
	}

	cases := []struct {
		Name     string
		Root     string
		Files    map[string]string
		Expected helper.Issues
	}{
		{
			Name: "local module with unregistered namespace",
			Root: `
module "app" {
  source = "./modules/app"
}`,
			Files:    map[string]string{"modules/app/main.tf": grafana},
			Expected: issue("module.app"),
		},
		{
			Name: "local module with registered namespace",
			Root: `
module "app" {
  source = "./modules/app"
}`,
			Files:    map[string]string{"modules/app/main.tf": network},
			Expected: helper.Issues{},
		},
		{
			Name: "duplicate resources are reported once",
			Root: `
module "app" {
  source = "./modules/app"
}`,
			Files: map[string]string{
				"modules/app/a.tf": grafana,
				"modules/app/b.tf": `
resource "azurerm_dashboard_grafana" "other" {
  grafana_major_version = 11
}`,
			},
			Expected: issue("module.app"),
		},
		{
			Name: "nested local module",
			Root: `
module "app" {
  source = "./modules/app"
}`,
			Files: map[string]string{
				"modules/app/main.tf": `
module "inner" {
  source = "../inner"
}`,
				"modules/inner/main.tf": grafana,
			},
			Expected: issue("module.app.module.inner"),
		},
		{
			Name: "aliased provider passed to module",
			Root: `
module "app" {
  source = "./modules/app"
  providers = {
    azurerm = azurerm.dashboard
  }
}`,
			Files:    map[string]string{"modules/app/main.tf": grafana},
			Expected: helper.Issues{},
		},
		{
			Name: "aliased provider passed to nested module",
			Root: `
module "app" {
  source = "./modules/app"
  providers = {
    azurerm      = azurerm
    azurerm.dash = azurerm.dashboard
  }
}`,
			Files: map[string]string{
				"modules/app/main.tf": `
module "inner" {
  source = "../inner"
  providers = {
    azurerm = azurerm.dash
  }
}`,
				"modules/inner/main.tf": grafana,
			},
			Expected: helper.Issues{},
		},
		{
			Name: "default provider implicitly inherited alongside explicit aliases",
			Root: `
module "app" {
  source = "./modules/app"
  providers = {
    azurerm.dash = azurerm.dashboard
  }
}`,
			Files:    map[string]string{"modules/app/main.tf": grafana},
			Expected: issue("module.app"),
		},
		{
			Name: "quoted provider reference in module",
			Root: `
module "app" {
  source = "./modules/app"
  providers = {
    azurerm      = azurerm.dashboard
    azurerm.dash = azurerm
  }
}`,
			Files: map[string]string{"modules/app/main.tf": `
resource "azurerm_dashboard_grafana" "this" {
  provider = "azurerm.dash"
}`},
			Expected: issue("module.app"),
		},
		{
			Name: "resource using alias not passed to module",
			Root: `
module "app" {
  source = "./modules/app"
}`,
			Files: map[string]string{"modules/app/main.tf": `
resource "azurerm_dashboard_grafana" "this" {
  provider = azurerm.dash
}`},
			Expected: helper.Issues{},
		},
		{
			Name: "module with its own provider configuration",
			Root: `
module "app" {
  source = "./modules/app"
}`,
			Files: map[string]string{"modules/app/main.tf": `
provider "azurerm" {
  features {}
}
` + grafana},
			Expected: helper.Issues{},
		},
		{
			Name: "installed remote module",
			Root: `
module "app" {
  source  = "example/app/azurerm"
  version = "1.0.0"
}`,
			Files: map[string]string{
				".terraform/modules/modules.json": `{"Modules":[{"Key":"","Source":"","Dir":"."},{"Key":"app","Source":"registry.terraform.io/example/app/azurerm","Version":"1.0.0","Dir":".terraform/modules/app"}]}`,
				".terraform/modules/app/main.tf":  grafana,
			},
			Expected: issue("module.app"),
		},
		{
			Name: "remote module not installed",
			Root: `
module "app" {
  source = "example/app/azurerm"
}`,
			Files:    map[string]string{},
			Expected: helper.Issues{},
		},
		{
			Name: "local module not found",
			Root: `
module "app" {
  source = "./modules/missing"
}`,
			Files:    map[string]string{},
			Expected: helper.Issues{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			dir := t.TempDir()
			for name, content := range tc.Files {
				path := filepath.Join(dir, name)
				if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
					t.Fatal(err)
				}
			}
			t.Setenv("TF_DATA_DIR", "")

			runner := helper.TestRunner(t, map[string]string{
				"main.tf":     tc.Root,
				"provider.tf": provider,
				".tflint.hcl": `
rule "azurerm_resource_provider_not_registered" {
  enabled = true
}`,
			})

			rule := NewAzurermResourceProviderNotRegisteredRule()
			rule.rootDir = dir
			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
