mapping "azurerm_firewall_network_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/firewall.json"

  priority = AzureFirewallNetworkRuleCollectionPropertiesFormat.priority
  action   = AzureFirewallRCActionType
}
