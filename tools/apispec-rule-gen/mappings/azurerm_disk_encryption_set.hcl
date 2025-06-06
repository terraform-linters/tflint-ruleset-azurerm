mapping "azurerm_disk_encryption_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/disk.json"

  name                = any // DiskEncryptionSetNameParameter
  resource_group_name = any // ResourceGroupNameParameter
}
