<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_network_packet_capture_invalid_maximum_capture_duration

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

The rule requires the value to be in the range 0 <= x <= 18000.

## Example

```hcl
resource "azurerm_network_packet_capture" "foo" {
  maximum_capture_duration = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: maximum_capture_duration must be 18000 or less (azurerm_network_packet_capture_invalid_maximum_capture_duration)

  on template.tf line 2:
  2:   maximum_capture_duration = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_network_packet_capture_invalid_maximum_capture_duration.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/networkWatcher.json
