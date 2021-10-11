mapping "azurerm_iothub_endpoint_servicebus_queue" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-03-31/iothub.json"

  name              = resourceName
  connection_string = RoutingServiceBusQueueEndpointProperties.connectionString
}
