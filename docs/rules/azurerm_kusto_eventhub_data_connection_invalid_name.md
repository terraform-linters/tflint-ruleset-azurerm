<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_kusto_eventhub_data_connection_invalid_name

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

In this rule, the string must match the regular expression `^.*$``.

## Example

```hcl
resource "azurerm_kusto_eventhub_data_connection" "foo" {
  name = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." does not match valid pattern ^.*$ (azurerm_kusto_eventhub_data_connection_invalid_name)

  on template.tf line 2:
  2:   name = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_kusto_eventhub_data_connection_invalid_name.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/kusto.json