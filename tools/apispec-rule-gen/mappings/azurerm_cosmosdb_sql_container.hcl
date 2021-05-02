mapping "azurerm_cosmosdb_sql_container" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-01-15/cosmos-db.json"

  name                = containerNameParameter
  resource_group_name = resourceGroupNameParameter
  account_name        = accountNameParameter
  database_name       = databaseNameParameter
  throughput          = any
  default_ttl         = SqlContainerResource.defaultTtl
}
