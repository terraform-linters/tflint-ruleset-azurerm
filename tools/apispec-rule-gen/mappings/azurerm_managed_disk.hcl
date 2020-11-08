mapping "azurerm_managed_disk" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-11-01/disk.json"

  name                 = DiskNameParameter
  resource_group_name  = ResourceGroupNameParameter
  storage_account_type = DiskSku.name
  create_option        = CreationData.createOption
  disk_iops_read_write = DiskProperties.diskIOPSReadWrite
  disk_mbps_read_write = DiskProperties.diskMBpsReadWrite
  disk_size_gb         = DiskProperties.diskSizeGB
  os_type              = DiskProperties.osType
  source_resource_id   = CreationData.sourceResourceId
  source_uri           = CreationData.sourceUri
}