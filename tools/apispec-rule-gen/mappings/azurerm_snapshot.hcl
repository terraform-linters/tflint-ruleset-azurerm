mapping "azurerm_snapshot" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/disk.json"

  name                = any // SnapshotNameParameter
  resource_group_name = any // ResourceGroupNameParameter
  create_option       = any // CreationData.createOption
  source_uri          = any // CreationData.sourceUri
  source_resource_id  = any // CreationData.sourceResourceId
  storage_account_id  = any // CreationData.storageAccountId
  disk_size_gb        = DiskProperties.diskSizeGB
}
