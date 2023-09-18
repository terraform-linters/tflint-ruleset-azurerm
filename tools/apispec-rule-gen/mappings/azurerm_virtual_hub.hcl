mapping "azurerm_virtual_hub" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/virtualWan.json"

  address_prefix = VirtualHubProperties.addressPrefix
}
