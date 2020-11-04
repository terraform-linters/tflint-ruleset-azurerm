mapping "azurerm_mariadb_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name        = virtualNetworkRuleNameParameter
  server_name = ServerNameParameter
  subnet_id   = VirtualNetworkRuleProperties.virtualNetworkSubnetId
}