mapping "azurerm_iothub_consumer_group" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/IoTHub/preview/2022-04-30-preview/iothub.json"

  name                = resourceName
  iothub_name         = resourceName
  resource_group_name = resourceGroupName
}
