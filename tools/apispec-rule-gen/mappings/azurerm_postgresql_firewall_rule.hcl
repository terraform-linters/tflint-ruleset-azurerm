mapping "azurerm_postgresql_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/FirewallRules.json"

  name             = FirewallRuleNameParameter
  server_name      = any //ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}
