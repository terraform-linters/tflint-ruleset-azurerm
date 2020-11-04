mapping "azurerm_iotcentral_application" {
  import_path = "azure-rest-api-specs/specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2018-09-01/iotcentral.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  location            = Resource.location
  sub_domain          = AppProperties.subdomain
  display_name        = AppProperties.displayName
  sku                 = AppSkuInfo.name
  template            = AppProperties.template
}
