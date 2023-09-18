mapping "azurerm_virtual_network_gateway_connection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/virtualNetworkGateway.json"

  type                               = VirtualNetworkGatewayConnectionType
  authorization_key                  = VirtualNetworkGatewayConnectionPropertiesFormat.authorizationKey
  routing_weight                     = VirtualNetworkGatewayConnectionPropertiesFormat.routingWeight
  shared_key                         = VirtualNetworkGatewayConnectionPropertiesFormat.sharedKey
  connection_protocol                = ConnectionProtocol
  enable_bgp                         = VirtualNetworkGatewayConnectionPropertiesFormat.enableBgp
  express_route_gateway_bypass       = VirtualNetworkGatewayConnectionPropertiesFormat.expressRouteGatewayBypass
  use_policy_based_traffic_selectors = VirtualNetworkGatewayConnectionPropertiesFormat.usePolicyBasedTrafficSelectors
}
