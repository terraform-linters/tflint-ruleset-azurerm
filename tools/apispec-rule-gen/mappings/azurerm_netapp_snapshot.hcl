mapping "azurerm_netapp_snapshot" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/NetApp/stable/2025-06-01/netapp.json"

  name                = any //SnapshotName
  resource_group_name = any //ResourceGroup
  account_name        = any //AccountName
  pool_name           = any //PoolName
  volume_name         = any //VolumeName
  location            = any //location
}
