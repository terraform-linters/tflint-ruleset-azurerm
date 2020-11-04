mapping "azurerm_mariadb_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name             = FirewallRuleNameParameter
  server_name      = ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}