mapping "azurerm_point_to_site_vpn_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/virtualWan.json"

  scale_unit = P2SVpnGatewayProperties.vpnGatewayScaleUnit
}
