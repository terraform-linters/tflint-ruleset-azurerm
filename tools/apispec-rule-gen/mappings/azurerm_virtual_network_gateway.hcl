mapping "azurerm_virtual_network_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/virtualNetworkGateway.json"

  type          = VirtualNetworkGatewayPropertiesFormat.gatewayType
  vpn_type      = VirtualNetworkGatewayPropertiesFormat.vpnType
  enable_bgp    = VirtualNetworkGatewayPropertiesFormat.enableBgp
  active_active = VirtualNetworkGatewayPropertiesFormat.activeActive
  sku           = VirtualNetworkGatewaySku.name
  generation    = VirtualNetworkGatewayPropertiesFormat.vpnGatewayGeneration
}
