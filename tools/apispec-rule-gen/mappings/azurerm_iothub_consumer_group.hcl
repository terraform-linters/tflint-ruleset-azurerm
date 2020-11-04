mapping "azurerm_iothub_consumer_group" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2020-03-01/iothub.json"

  name                = resourceName
  iothub_name         = resourceName
  resource_group_name = resourceGroupName
}
