mapping "azurerm_healthcare_service" {
  import_path = "azure-rest-api-specs/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2022-12-01/healthcare-apis.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  location            = locationName
  cosmosdb_throughput = ServiceCosmosDbConfigurationInfo.offerThroughput
  kind                = ServicesResource.kind
}
