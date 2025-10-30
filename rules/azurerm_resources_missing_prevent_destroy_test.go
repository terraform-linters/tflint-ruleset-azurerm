package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermResourcesMissingPreventDestroy(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "Missing lifecycle block entirely",
			Content: `
resource "azurerm_storage_account" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourcesMissingPreventDestroyRule(),
					Message: "Resource is missing lifecycle { prevent_destroy = true }. This resource contains data that should be protected from accidental deletion.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "Has lifecycle but prevent_destroy = false",
			Content: `
resource "azurerm_storage_account" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
  
  lifecycle {
    prevent_destroy = false
  }
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourcesMissingPreventDestroyRule(),
					Message: "Resource has lifecycle { prevent_destroy = false }. Consider setting prevent_destroy = true to protect data from accidental deletion.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "Has lifecycle with prevent_destroy = true",
			Content: `
resource "azurerm_storage_account" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
  
  lifecycle {
    prevent_destroy = true
  }
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Has lifecycle but missing prevent_destroy attribute",
			Content: `
resource "azurerm_key_vault" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
  
  lifecycle {
    create_before_destroy = true
  }
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourcesMissingPreventDestroyRule(),
					Message: "Resource is missing lifecycle { prevent_destroy = true }. This resource contains data that should be protected from accidental deletion.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 39},
					},
				},
			},
		},
		{
			Name: "Excluding resource type",
			Content: `
resource "azurerm_storage_account" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
  exclude = ["azurerm_storage_account"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Custom resources list - not in default list",
			Content: `
resource "azurerm_virtual_machine" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
  resources = ["azurerm_virtual_machine"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourcesMissingPreventDestroyRule(),
					Message: "Resource is missing lifecycle { prevent_destroy = true }. This resource contains data that should be protected from accidental deletion.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 45},
					},
				},
			},
		},
		{
			Name: "Resource not in default list and no custom resources specified",
			Content: `
resource "azurerm_virtual_machine" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Multiple default resources - some missing lifecycle",
			Content: `
resource "azurerm_storage_account" "example1" {
  name                = "example1"
  location            = "West Europe"
  resource_group_name = "example"
}

resource "azurerm_key_vault" "example2" {
  name                = "example2"
  location            = "West Europe"
  resource_group_name = "example"
  
  lifecycle {
    prevent_destroy = true
  }
}

resource "azurerm_mssql_server" "example3" {
  name                = "example3"
  location            = "West Europe"
  resource_group_name = "example"
}`,
			Config: `
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourcesMissingPreventDestroyRule(),
					Message: "Resource is missing lifecycle { prevent_destroy = true }. This resource contains data that should be protected from accidental deletion.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 46},
					},
				},
				{
					Rule:    NewAzurermResourcesMissingPreventDestroyRule(),
					Message: "Resource is missing lifecycle { prevent_destroy = true }. This resource contains data that should be protected from accidental deletion.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 18, Column: 1},
						End:      hcl.Pos{Line: 18, Column: 43},
					},
				},
			},
		},
	}

	rule := NewAzurermResourcesMissingPreventDestroyRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content, ".tflint.hcl": tc.Config})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
