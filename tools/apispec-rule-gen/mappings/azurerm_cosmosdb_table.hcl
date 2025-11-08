mapping "azurerm_cosmosdb_table" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2023-04-15/cosmos-db.json"

  name                = tableNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}
