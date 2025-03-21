mapping "azurerm_databricks_workspace" {
  import_path = "azure-rest-api-specs/specification/databricks/resource-manager/Microsoft.Databricks/stable/2024-05-01/databricks.json"

  name                        = WorkspaceName
  resource_group_name         = any //ResourceGroupName
  sku                         = Sku.name
  managed_resource_group_name = any //ResourceGroupName
}
