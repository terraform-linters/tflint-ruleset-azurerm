mapping "azurerm_databricks_workspace" {
  import_path = "azure-rest-api-specs/specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/databricks.json"

  name                        = WorkspaceName
  resource_group_name         = ResourceGroupName
  sku                         = Sku.name
  managed_resource_group_name = ResourceGroupName
}
