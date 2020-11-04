mapping "azurerm_postgresql_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name             = FirewallRuleNameParameter
  server_name      = ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}