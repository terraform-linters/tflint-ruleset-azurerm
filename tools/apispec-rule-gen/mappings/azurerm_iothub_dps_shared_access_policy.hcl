mapping "azurerm_iothub_dps_shared_access_policy" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-03-31/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iothub_dps_name     = resourceName
}
