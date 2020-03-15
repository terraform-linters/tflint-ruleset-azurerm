mapping "azurerm_netapp_account" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2019-11-01/netapp.json"

  name                = AccountName
  resource_group_name = ResourceGroup
  location            = location
}

mapping "azurerm_netapp_pool" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2019-11-01/netapp.json"

  name                = PoolName
  resource_group_name = ResourceGroup
  account_name        = AccountName
  location            = location
  service_level       = poolProperties.serviceLevel
}

mapping "azurerm_netapp_volume" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2019-11-01/netapp.json"

  name                = VolumeName
  resource_group_name = ResourceGroup
  location            = location
  account_name        = AccountName
  pool_name           = PoolName
  service_level       = volumeProperties.serviceLevel
  protocols           = volumeProperties.protocolTypes
}

mapping "azurerm_netapp_snapshot" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2019-11-01/netapp.json"

  name                = SnapshotName
  resource_group_name = ResourceGroup
  account_name        = AccountName
  pool_name           = PoolName
  volume_name         = VolumeName
  location            = location
}
