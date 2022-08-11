mapping "azurerm_snapshot" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2019-11-01/disk.json"

  name                = SnapshotNameParameter
  resource_group_name = ResourceGroupNameParameter
  create_option       = CreationData.createOption
  source_uri          = CreationData.sourceUri
  source_resource_id  = CreationData.sourceResourceId
  storage_account_id  = CreationData.storageAccountId
  disk_size_gb        = DiskProperties.diskSizeGB
}