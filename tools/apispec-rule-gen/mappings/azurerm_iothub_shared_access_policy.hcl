mapping "azurerm_iothub_shared_access_policy" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-03-31/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iothub_name         = resourceName
}
