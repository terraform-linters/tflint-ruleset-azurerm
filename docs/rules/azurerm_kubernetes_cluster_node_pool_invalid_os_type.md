<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_kubernetes_cluster_node_pool_invalid_os_type

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- Linux
- Windows

## Example

```hcl
resource "azurerm_kubernetes_cluster_node_pool" "foo" {
  os_type = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as os_type (azurerm_kubernetes_cluster_node_pool_invalid_os_type)

  on template.tf line 2:
  2:   os_type = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_kubernetes_cluster_node_pool_invalid_os_type.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-02-01/managedClusters.json
