mapping "azurerm_search_service" {
  import_path = "azure-rest-api-specs/specification/search/resource-manager/Microsoft.Search/stable/2022-09-01/search.json"

  location            = any //Resource.location
  name                = SearchServiceNameParameter
  resource_group_name = ResourceGroupNameParameter
  sku                 = Sku.name
  partition_count     = SearchServiceProperties.partitionCount
  replica_count       = SearchServiceProperties.replicaCount
}
