mapping "azurerm_cosmosdb_gremlin_graph" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2024-08-15/cosmos-db.json"

  name                = graphNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  database_name       = databaseNameParameter
  throughput          = any
}
