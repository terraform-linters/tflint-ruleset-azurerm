mapping "azurerm_mysql_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name             = FirewallRuleNameParameter
  server_name      = ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}