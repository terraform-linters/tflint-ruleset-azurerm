mapping "azurerm_search_service" {
  import_path = "azure-rest-api-specs/specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/search.json"

  location            = any //Resource.location
  name                = any //SearchServiceNameParameter
  resource_group_name = ResourceGroupNameParameter
  sku                 = Sku.name
  partition_count     = SearchServiceProperties.partitionCount
  replica_count       = SearchServiceProperties.replicaCount
}
