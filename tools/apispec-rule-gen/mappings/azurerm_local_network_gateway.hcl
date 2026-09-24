mapping "azurerm_local_network_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/networkGateway.json"

  gateway_address = LocalNetworkGatewayPropertiesFormat.gatewayIpAddress
}
