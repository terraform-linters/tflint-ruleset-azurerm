mapping "azurerm_cosmosdb_account" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                              = accountNameParameter
  resource_group_name               = resourceGroupNameParameter
  offer_type                        = DatabaseAccountOfferType
  kind                              = DatabaseAccountCreateUpdateParameters.kind
  ip_range_filter                   = IPRangeFilter
  enable_automatic_failover         = DatabaseAccountCreateUpdateProperties.enableAutomaticFailover
  is_virtual_network_filter_enabled = DatabaseAccountCreateUpdateProperties.isVirtualNetworkFilterEnabled
  enable_multiple_write_locations   = DatabaseAccountCreateUpdateProperties.enableMultipleWriteLocations
}

mapping "azurerm_cosmosdb_cassandra_keyspace" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = keyspaceNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}

mapping "azurerm_cosmosdb_gremlin_database" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = databaseNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}

mapping "azurerm_cosmosdb_gremlin_graph" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = graphNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  database_name       = databaseNameParameter
  throughput          = any
}

mapping "azurerm_cosmosdb_mongo_collection" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = collectionNameParameter
  resource_group_name = resourceGroupNameParameter
  database_name       = databaseNameParameter
  throughput          = any
}

mapping "azurerm_cosmosdb_mongo_database" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = databaseNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}

mapping "azurerm_cosmosdb_sql_container" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = containerNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  database_name       = databaseNameParameter
  throughput          = any
  default_ttl         = SqlContainerResource.defaultTtl
}

mapping "azurerm_cosmosdb_sql_database" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = databaseNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}

mapping "azurerm_cosmosdb_table" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2015-04-08/cosmos-db.json"

  name                = tableNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  throughput          = any
}
