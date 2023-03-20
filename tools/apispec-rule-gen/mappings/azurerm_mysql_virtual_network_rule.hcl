mapping "azurerm_mysql_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/mysql.json"

  name        = any //virtualNetworkRuleNameParameter
  server_name = ServerNameParameter
  subnet_id   = any //VirtualNetworkRuleProperties.virtualNetworkSubnetId
}
