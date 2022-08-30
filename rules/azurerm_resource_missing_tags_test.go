package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AzurermResourceMissingTags(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "Wanted tags: Bar,Foo, found: bar,foo",
			Content: `
resource "azurerm_virtual_machine" "vm" {
  vm_size = "Standard_DS1_v2"
  tags = {
    foo = "bar"
    bar = "baz"
  }
}`,
			Config: `
rule "azurerm_resource_missing_tags" {
  enabled = true
  tags = ["Foo", "Bar"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceMissingTagsRule(),
					Message: "The resource is missing the following tags: \"Bar\", \"Foo\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 10},
						End:      hcl.Pos{Line: 7, Column: 4},
					},
				},
			},
		},
		{
			Name: "No tags",
			Content: `
resource "azurerm_virtual_machine" "vm" {
	vm_size = "Standard_DS1_v2"
}`,
			Config: `
rule "azurerm_resource_missing_tags" {
  enabled = true
  tags = ["Foo", "Bar"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAzurermResourceMissingTagsRule(),
					Message: "The resource is missing the following tags: \"Bar\", \"Foo\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 40},
					},
				},
			},
		},
		{
			Name: "Tags are correct",
			Content: `
resource "azurerm_virtual_machine" "vm" {
	vm_size = "Standard_DS1_v2"
  tags = {
    Foo = "bar"
    Bar = "baz"
  }
}`,
			Config: `
rule "azurerm_resource_missing_tags" {
  enabled = true
  tags = ["Foo", "Bar"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "excluding resource type",
			Content: `
resource "azurerm_virtual_machine" "vm" {
  vm_size = "Standard_DS1_v2"
  tags = {
    foo = "bar"
    bar = "baz"
  }
}`,
			Config: `
rule "azurerm_resource_missing_tags" {
  enabled = true
  exclude = ["azurerm_virtual_machine"]
  tags = ["Foo", "Bar"]
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAzurermResourceMissingTagsRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content, ".tflint.hcl": tc.Config})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
