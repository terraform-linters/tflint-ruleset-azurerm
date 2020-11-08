mapping "azurerm_iothub_route" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2020-03-01/iothub.json"

  name                = RouteProperties.name
  resource_group_name = resourceGroupName
  iothub_name         = resourceName
  source              = RouteProperties.source
  condition           = RouteProperties.condition
  endpoint_names      = RouteProperties.endpointNames
  enabled             = RouteProperties.isEnabled
}
