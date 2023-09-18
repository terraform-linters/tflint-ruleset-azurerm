mapping "azurerm_express_route_circuit_peering" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/expressRouteCircuit.json"

  peering_type                  = ExpressRoutePeeringType
  primary_peer_address_prefix   = ExpressRouteCircuitPeeringPropertiesFormat.primaryPeerAddressPrefix
  secondary_peer_address_prefix = ExpressRouteCircuitPeeringPropertiesFormat.secondaryPeerAddressPrefix
  vlan_id                       = ExpressRouteCircuitPeeringPropertiesFormat.vlanId
  shared_key                    = ExpressRouteCircuitPeeringPropertiesFormat.sharedKey
  peer_asn                      = any // ExpressRouteCircuitPeeringPropertiesFormat.peerASN
}
