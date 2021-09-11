mapping "azurerm_firewall_network_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2021-02-01/azureFirewall.json"

  priority = AzureFirewallNetworkRuleCollectionPropertiesFormat.priority
  action   = AzureFirewallRCActionType
}
