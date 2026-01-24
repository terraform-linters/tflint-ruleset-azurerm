mapping "azurerm_search_service" {
  import_path = "azure-rest-api-specs/specification/search/resource-manager/Microsoft.Search/Search/stable/2025-05-01/search.json"

  location            = any //Resource.location
  name                = any //SearchServiceNameParameter
  resource_group_name = ResourceGroupNameParameter
  sku                 = Sku.name
  partition_count     = SearchServiceProperties.partitionCount
  replica_count       = SearchServiceProperties.replicaCount
}
