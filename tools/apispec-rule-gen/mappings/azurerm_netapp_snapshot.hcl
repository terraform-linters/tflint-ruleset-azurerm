mapping "azurerm_netapp_snapshot" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2019-10-01/netapp.json"

  name                = SnapshotName
  resource_group_name = ResourceGroup
  account_name        = AccountName
  pool_name           = PoolName
  volume_name         = VolumeName
  location            = location
}
