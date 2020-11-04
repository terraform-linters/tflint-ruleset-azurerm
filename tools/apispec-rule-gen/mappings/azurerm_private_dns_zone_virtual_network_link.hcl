mapping "azurerm_private_dns_zone_virtual_network_link" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  registration_enabled = VirtualNetworkLinkProperties.registrationEnabled
}
