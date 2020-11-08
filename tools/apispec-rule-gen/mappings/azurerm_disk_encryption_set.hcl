mapping "azurerm_disk_encryption_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-11-01/disk.json"

  name                = DiskEncryptionSetNameParameter
  resource_group_name = ResourceGroupNameParameter
}
