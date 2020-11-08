mapping "azurerm_storage_account_network_rules" {
  import_path = "azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/storage.json"

  storage_account_name = StorageAccountName
  resource_group_name  = ResourceGroupName
  default_action       = NetworkRuleSet.defaultAction
  bypass               = any // NetworkRuleSet.bypass
  ip_rules             = NetworkRuleSet.ipRules
}