mapping "azurerm_local_network_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2025-01-01/virtualNetworkGateway.json"

  gateway_address = LocalNetworkGatewayPropertiesFormat.gatewayIpAddress
}
