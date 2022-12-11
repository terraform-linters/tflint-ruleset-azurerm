mapping "azurerm_iothub_endpoint_servicebus_topic" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/iothub.json"

  name              = resourceName
  connection_string = RoutingServiceBusTopicEndpointProperties.connectionString
}
