mapping "azurerm_subnet" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/virtualNetwork.json"

  name              = any // Subnet.name is no longer defined
  address_prefix    = SubnetPropertiesFormat.addressPrefix
  address_prefixes  = SubnetPropertiesFormat.addressPrefixes
  service_endpoints = SubnetPropertiesFormat.serviceEndpoints
}
