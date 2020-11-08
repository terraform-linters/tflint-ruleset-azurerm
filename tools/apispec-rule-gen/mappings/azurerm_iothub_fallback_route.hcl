mapping "azurerm_iothub_fallback_route" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2020-03-01/iothub.json"

  resource_group_name = resourceGroupName
  iothub_name         = resourceName
  enabled             = FallbackRouteProperties.isEnabled
  endpoint_names      = FallbackRouteProperties.endpointNames
  condition           = FallbackRouteProperties.condition
}
