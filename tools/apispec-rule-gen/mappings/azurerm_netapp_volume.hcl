mapping "azurerm_netapp_volume" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-01-01/netapp.json"

  name                = VolumeName
  resource_group_name = any //ResourceGroup
  location            = any //location
  account_name        = AccountName
  pool_name           = PoolName
  service_level       = serviceLevel
  protocols           = volumeProperties.protocolTypes
}
