mapping "azurerm_firewall_application_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/firewall.json"

  priority = AzureFirewallApplicationRuleCollectionPropertiesFormat.priority
  action   = AzureFirewallRCActionType
}
