<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_frontdoor_invalid_backend_pools_send_receive_timeout_seconds

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

The rule requires the value to be 16 or higher.

## Example

```hcl
resource "azurerm_frontdoor" "foo" {
  backend_pools_send_receive_timeout_seconds = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: backend_pools_send_receive_timeout_seconds must be 16 or higher (azurerm_frontdoor_invalid_backend_pools_send_receive_timeout_seconds)

  on template.tf line 2:
  2:   backend_pools_send_receive_timeout_seconds = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_frontdoor_invalid_backend_pools_send_receive_timeout_seconds.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/frontdoor.json
