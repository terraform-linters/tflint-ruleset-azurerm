mapping "azurerm_hpc_cache" {
  import_path = "azure-rest-api-specs/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2020-03-01/storagecache.json"

  name     = ResourceName
  location = Cache.location
}

mapping "azurerm_hpc_cache_nfs_target" {
  import_path = "azure-rest-api-specs/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2020-03-01/storagecache.json"

  usage_model = UsageModel.targetType
}

mapping "azurerm_storage_account" {
  import_path = "azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json"

  name                      = StorageAccountName
  resource_group_name       = ResourceGroupName
  account_kind              = StorageAccount.kind
  access_tier               = StorageAccountProperties.accessTier
  enable_https_traffic_only = StorageAccountProperties.supportsHttpsTrafficOnly
  is_hns_enabled            = StorageAccountProperties.isHnsEnabled
}

mapping "azurerm_storage_account_customer_managed_key" {
  import_path = "azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json"

  key_name    = KeyVaultProperties.keyname
  key_version = KeyVaultProperties.keyversion
}

mapping "azurerm_storage_account_network_rules" {
  import_path = "azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json"

  storage_account_name = StorageAccountName
  resource_group_name  = ResourceGroupName
  default_action       = NetworkRuleSet.defaultAction
  bypass               = any // NetworkRuleSet.bypass
  ip_rules             = NetworkRuleSet.ipRules
}
