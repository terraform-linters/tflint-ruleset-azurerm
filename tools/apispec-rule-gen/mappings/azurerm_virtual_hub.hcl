mapping "azurerm_virtual_hub" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-05-01/virtualWan.json"

  address_prefix = VirtualHubProperties.addressPrefix
}
