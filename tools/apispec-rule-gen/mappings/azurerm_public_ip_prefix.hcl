mapping "azurerm_public_ip_prefix" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/publicIpPrefix.json"

  sku           = PublicIPPrefixSku.name
  prefix_length = PublicIPPrefixPropertiesFormat.prefixLength
  zones         = PublicIPPrefix.zones
}
