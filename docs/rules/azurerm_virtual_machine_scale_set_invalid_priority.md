<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_virtual_machine_scale_set_invalid_priority

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- Regular
- Low
- Spot

## Example

```hcl
resource "azurerm_virtual_machine_scale_set" "foo" {
  priority = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as priority (azurerm_virtual_machine_scale_set_invalid_priority)

  on template.tf line 2:
  2:   priority = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_virtual_machine_scale_set_invalid_priority.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/ComputeRP.json
