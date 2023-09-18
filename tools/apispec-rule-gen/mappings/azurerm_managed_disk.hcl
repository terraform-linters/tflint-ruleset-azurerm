mapping "azurerm_managed_disk" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/disk.json"

  name                 = any // DiskNameParameter
  resource_group_name  = any // ResourceGroupNameParameter
  storage_account_type = DiskSku.name
  create_option        = any // CreationData.createOption
  disk_iops_read_write = DiskProperties.diskIOPSReadWrite
  disk_mbps_read_write = DiskProperties.diskMBpsReadWrite
  disk_size_gb         = DiskProperties.diskSizeGB
  os_type              = DiskProperties.osType
  source_resource_id   = any // CreationData.sourceResourceId
  source_uri           = any // CreationData.sourceUri
}
