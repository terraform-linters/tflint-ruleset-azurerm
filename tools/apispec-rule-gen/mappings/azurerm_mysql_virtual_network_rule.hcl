mapping "azurerm_mysql_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name        = virtualNetworkRuleNameParameter
  server_name = ServerNameParameter
  subnet_id   = VirtualNetworkRuleProperties.virtualNetworkSubnetId
}
