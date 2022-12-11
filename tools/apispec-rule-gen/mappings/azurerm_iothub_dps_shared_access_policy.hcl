mapping "azurerm_iothub_dps_shared_access_policy" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iothub_dps_name     = resourceName
}
