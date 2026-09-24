mapping "azurerm_route_table" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/virtualNetwork.json"

  disable_bgp_route_propagation = RouteTablePropertiesFormat.disableBgpRoutePropagation
}
