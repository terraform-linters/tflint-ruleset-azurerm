mapping "azurerm_virtual_network_peering" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/virtualNetwork.json"

  allow_virtual_network_access = VirtualNetworkPeeringPropertiesFormat.allowVirtualNetworkAccess
  allow_forwarded_traffic      = VirtualNetworkPeeringPropertiesFormat.allowForwardedTraffic
  allow_gateway_transit        = VirtualNetworkPeeringPropertiesFormat.allowGatewayTransit
  use_remote_gateways          = VirtualNetworkPeeringPropertiesFormat.useRemoteGateways
}
