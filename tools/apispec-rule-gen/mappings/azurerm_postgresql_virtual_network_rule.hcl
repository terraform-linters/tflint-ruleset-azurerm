mapping "azurerm_postgresql_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name        = virtualNetworkRuleNameParameter
  server_name = ServerNameParameter
  subnet_id   = VirtualNetworkRuleProperties.virtualNetworkSubnetId
}