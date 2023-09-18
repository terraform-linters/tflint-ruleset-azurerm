mapping "azurerm_vpn_server_configuration" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/virtualWan.json"

  vpn_authentication_types = VpnServerConfigurationProperties.vpnAuthenticationTypes
  vpn_protocols            = VpnServerConfigurationProperties.vpnProtocols
}
