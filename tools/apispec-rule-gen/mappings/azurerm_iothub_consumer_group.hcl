mapping "azurerm_iothub_consumer_group" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-03-31/iothub.json"

  name                = resourceName
  iothub_name         = resourceName
  resource_group_name = resourceGroupName
}
