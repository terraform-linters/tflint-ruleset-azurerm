mapping "azurerm_iothub_endpoint_eventhub" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/IoTHub/preview/2022-04-30-preview/iothub.json"

  name              = resourceName
  connection_string = RoutingEventHubProperties.connectionString
}
