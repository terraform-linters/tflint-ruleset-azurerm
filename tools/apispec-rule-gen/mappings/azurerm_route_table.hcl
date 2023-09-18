mapping "azurerm_route_table" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/routeTable.json"

  disable_bgp_route_propagation = RouteTablePropertiesFormat.disableBgpRoutePropagation
}
