<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_managed_application_definition_invalid_lock_level

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- CanNotDelete
- ReadOnly
- None

## Example

```hcl
resource "azurerm_managed_application_definition" "foo" {
  lock_level = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as lock_level (azurerm_managed_application_definition_invalid_lock_level)

  on template.tf line 2:
  2:   lock_level = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_managed_application_definition_invalid_lock_level.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/managedapplications.json
