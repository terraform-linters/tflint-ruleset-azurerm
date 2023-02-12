mapping "azurerm_postgresql_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/VirtualNetwork.json"

  name        = any //virtualNetworkRuleNameParameter
  server_name = any //ServerNameParameter
  subnet_id   = any //VirtualNetworkRuleProperties.virtualNetworkSubnetId
}
