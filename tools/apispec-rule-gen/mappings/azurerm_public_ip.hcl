mapping "azurerm_public_ip" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/publicIpAddress.json"

  sku                     = PublicIPAddressSku.name
  idle_timeout_in_minutes = PublicIPAddressPropertiesFormat.idleTimeoutInMinutes
  domain_name_label       = PublicIPAddressDnsSettings.domainNameLabel
  reverse_fqdn            = PublicIPAddressDnsSettings.reverseFqdn
  zones                   = PublicIPAddress.zones
}
