mapping "azurerm_vpn_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/virtualWan.json"

  scale_unit = VpnGatewayProperties.vpnGatewayScaleUnit
}
