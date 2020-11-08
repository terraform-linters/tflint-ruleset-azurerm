mapping "azurerm_iothub_endpoint_servicebus_topic" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2020-03-01/iothub.json"

  name              = resourceName
  connection_string = RoutingServiceBusTopicEndpointProperties.connectionString
}
