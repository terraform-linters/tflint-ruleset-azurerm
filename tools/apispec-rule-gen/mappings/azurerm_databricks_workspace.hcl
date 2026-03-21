mapping "azurerm_databricks_workspace" {
  import_path = "azure-rest-api-specs/specification/databricks/resource-manager/Microsoft.Databricks/Databricks/stable/2026-01-01/openapi.json"

  name                        = any // workspaceName is now an inline path parameter in the 2026-01-01 spec.
  resource_group_name         = any // resourceGroupName is now referenced from common-types/resource-management/v5.
  sku                         = Sku.name
  managed_resource_group_name = any // managedResourceGroupId in the spec is not a Terraform RG name.
}
