mapping "azurerm_netapp_volume" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/NetApp/stable/2025-06-01/netapp.json"

  name                = any //VolumeName
  resource_group_name = any //ResourceGroup
  location            = any //location
  account_name        = any //AccountName
  pool_name           = any //PoolName
  service_level       = PoolProperties.serviceLevel
  protocols           = VolumePatchProperties.protocolTypes
}
