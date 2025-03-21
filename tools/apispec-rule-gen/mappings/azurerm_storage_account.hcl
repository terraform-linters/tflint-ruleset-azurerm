mapping "azurerm_storage_account" {
  import_path = "azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/storage.json"

  name                      = StorageAccountName
  resource_group_name       = ResourceGroupName
  account_kind              = StorageAccount.kind
  access_tier               = StorageAccountProperties.accessTier
  enable_https_traffic_only = StorageAccountProperties.supportsHttpsTrafficOnly
  is_hns_enabled            = StorageAccountProperties.isHnsEnabled
}
