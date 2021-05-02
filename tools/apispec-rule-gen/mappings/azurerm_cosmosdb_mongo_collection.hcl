mapping "azurerm_cosmosdb_mongo_collection" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-01-15/cosmos-db.json"

  name                = collectionNameParameter
  resource_group_name = resourceGroupNameParameter
  database_name       = databaseNameParameter
  throughput          = any
}
