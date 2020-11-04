mapping "azurerm_netapp_volume" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2019-10-01/netapp.json"

  name                = VolumeName
  resource_group_name = ResourceGroup
  location            = location
  account_name        = AccountName
  pool_name           = PoolName
  service_level       = volumeProperties.serviceLevel
  protocols           = volumeProperties.protocolTypes
}
