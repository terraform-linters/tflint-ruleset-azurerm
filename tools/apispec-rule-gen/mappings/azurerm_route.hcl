mapping "azurerm_route" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/routeTable.json"

  address_prefix         = RoutePropertiesFormat.addressPrefix
  next_hop_type          = RouteNextHopType
  next_hop_in_ip_address = RoutePropertiesFormat.nextHopIpAddress
}
