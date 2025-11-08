mapping "azurerm_express_route_circuit" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2025-01-01/expressRouteCircuit.json"

  service_provider_name    = ExpressRouteCircuitServiceProviderProperties.serviceProviderName
  peering_location         = ExpressRouteCircuitServiceProviderProperties.peeringLocation
  bandwidth_in_mbps        = ExpressRouteCircuitServiceProviderProperties.bandwidthInMbps
  allow_classic_operations = ExpressRouteCircuitPropertiesFormat.allowClassicOperations
}
