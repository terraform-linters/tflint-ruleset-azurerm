mapping "azurerm_firewall_application_rule_collection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/azureFirewall.json"

  priority = AzureFirewallApplicationRuleCollectionPropertiesFormat.priority
  action   = AzureFirewallRCActionType
}
