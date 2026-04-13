# azurerm_linux_virtual_machine_retired_size

Warns about size values that are retired or announced-for-retirement based on https://learn.microsoft.com/azure/virtual-machines/sizes/retirement/retired-sizes-list

Values that produce a warning are:

- Standard_B1ms
- Standard_B2ms
- Standard_B1s
- Standard_B2s
- Standard_D1_v2
- Standard_D2_v2
- Standard_D3_v2
- Standard_DS1_v2
- Standard_DS2_v2
- Standard_F1
- Standard_F2
- Standard_FS2
- Standard_F2s_v2
- Standard_G1
- Standard_G2
- Standard_GS1
- Standard_NC6_v3
- Standard_NC12_v3
- Standard_NC24_v3

## Example

```hcl
resource "azurerm_linux_virtual_machine" "foo" {
  size = "Standard_B1ms"
}
```

```
$ tflint
1 issue(s) found:

Notice: "Standard_B1ms" is a retired VM size and may no longer available (azurerm_linux_virtual_machine_retired_size)

  on template.tf line 2:
  2:   size = "Standard_B1ms"

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.32.0/docs/rules/azurerm_linux_virtual_machine_retired_size.md

```

## Why

Linux Virtual Machines configured to use these sizes may retrun an error when calling the API by `terraform apply`.

## How To Fix

Update the Virtual Machine size to a supports size.

## Source

https://learn.microsoft.com/en-gb/azure/virtual-machines/sizes/retirement/retired-sizes-list
