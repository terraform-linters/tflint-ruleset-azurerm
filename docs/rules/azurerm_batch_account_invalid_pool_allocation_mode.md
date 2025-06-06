<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_batch_account_invalid_pool_allocation_mode

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- BatchService
- UserSubscription

## Example

```hcl
resource "azurerm_batch_account" "foo" {
  pool_allocation_mode = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as pool_allocation_mode (azurerm_batch_account_invalid_pool_allocation_mode)

  on template.tf line 2:
  2:   pool_allocation_mode = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_batch_account_invalid_pool_allocation_mode.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/BatchManagement.json
