# azurerm_resource_missing_tags

Require specific tags for all Azurerm resource types that support them.


## Configuration

```hcl
rule "azurerm_resource_missing_tags" {
  enabled = true
  tags = ["Foo", "Bar"]
  exclude = [] # (Optional) Exclude some resource types from tag checks
}
```

## Example

```hcl
resource "azurerm_virtual_machine" "vm" {
	vm_size = "Standard_DS1_v2"
  tags = {
    foo = "Bar"
    bar = "Baz"
  }
}
```

```
$ tflint
1 issue(s) found:

Notice: The resource is missing the following tags: "Bar", "Foo". (azurerm_resource_missing_tags)

  on main.tf line 3:
   3:   tags = {
   4:     foo = "Bar"
   5:     bar = "Baz"
   6:   }

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.17.1/docs/rules/azurerm_resource_missing_tags.md

```

## Why

You want to set a standardized set of tags for your Azurerm resources.

## How To Fix

For each resource type that supports tags, ensure that each missing tag is present.
