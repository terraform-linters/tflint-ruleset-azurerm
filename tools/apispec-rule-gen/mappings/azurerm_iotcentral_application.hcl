mapping "azurerm_iotcentral_application" {
  import_path = "azure-rest-api-specs/specification/iotcentral/resource-manager/Microsoft.IoTCentral/IoTCentral/preview/2021-11-01-preview/iotcentral.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  location            = any //Resource.location
  sub_domain          = AppProperties.subdomain
  display_name        = AppProperties.displayName
  sku                 = AppSkuInfo.name
  template            = AppProperties.template
}
