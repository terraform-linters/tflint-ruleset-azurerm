mapping "azurerm_cosmosdb_cassandra_keyspace" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/cosmos-db.json"

  name                = keyspaceNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}
