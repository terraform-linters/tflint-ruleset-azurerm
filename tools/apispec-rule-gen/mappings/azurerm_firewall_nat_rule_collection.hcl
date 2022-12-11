mapping "azurerm_firewall_nat_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/azureFirewall.json"

  priority = AzureFirewallNatRuleCollectionProperties.priority
  action   = AzureFirewallNatRCActionType
}
