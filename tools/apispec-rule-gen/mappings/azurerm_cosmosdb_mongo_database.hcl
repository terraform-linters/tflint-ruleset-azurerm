mapping "azurerm_cosmosdb_mongo_database" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/cosmos-db.json"

  name                = databaseNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}
