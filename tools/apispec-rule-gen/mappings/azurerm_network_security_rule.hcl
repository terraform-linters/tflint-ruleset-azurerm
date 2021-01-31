mapping "azurerm_network_security_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-07-01/networkSecurityGroup.json"

  description                                = SecurityRulePropertiesFormat.description
  protocol                                   = SecurityRulePropertiesFormat.protocol
  source_port_range                          = SecurityRulePropertiesFormat.sourcePortRange
  source_port_ranges                         = SecurityRulePropertiesFormat.sourcePortRanges
  destination_port_range                     = SecurityRulePropertiesFormat.destinationPortRange
  destination_port_ranges                    = SecurityRulePropertiesFormat.destinationPortRanges
  source_address_prefix                      = SecurityRulePropertiesFormat.sourceAddressPrefix
  source_address_prefixes                    = SecurityRulePropertiesFormat.sourceAddressPrefixes
  source_application_security_group_ids      = SecurityRulePropertiesFormat.sourceApplicationSecurityGroups
  destination_address_prefix                 = SecurityRulePropertiesFormat.destinationAddressPrefix
  destination_address_prefixes               = SecurityRulePropertiesFormat.destinationAddressPrefixes
  destination_application_security_group_ids = SecurityRulePropertiesFormat.destinationApplicationSecurityGroups
  access                                     = SecurityRuleAccess
  priority                                   = SecurityRulePropertiesFormat.priority
  direction                                  = SecurityRuleDirection
}
