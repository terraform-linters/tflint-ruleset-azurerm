mapping "azurerm_storage_account_network_rules" {
  import_path = "azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2021-04-01/storage.json"

  default_action = NetworkRuleSet.defaultAction
  bypass         = any // NetworkRuleSet.bypass
  ip_rules       = NetworkRuleSet.ipRules
}
