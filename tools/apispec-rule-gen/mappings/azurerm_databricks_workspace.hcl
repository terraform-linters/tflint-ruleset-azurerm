mapping "azurerm_databricks_workspace" {
  import_path = "azure-rest-api-specs/specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/databricks.json"

  name                        = WorkspaceName
  resource_group_name         = ResourceGroupName
  sku                         = Sku.name
  managed_resource_group_name = ResourceGroupName
}
