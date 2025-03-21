mapping "azurerm_cosmosdb_account" {
  import_path = "azure-rest-api-specs/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json"

  name                              = accountNameParameter
  resource_group_name               = resourceGroupNameParameter
  offer_type                        = DatabaseAccountOfferType
  kind                              = DatabaseAccountCreateUpdateParameters.kind
  ip_range_filter                   = any // IPRangeFilter
  enable_automatic_failover         = DatabaseAccountCreateUpdateProperties.enableAutomaticFailover
  is_virtual_network_filter_enabled = DatabaseAccountCreateUpdateProperties.isVirtualNetworkFilterEnabled
  enable_multiple_write_locations   = DatabaseAccountCreateUpdateProperties.enableMultipleWriteLocations
}
