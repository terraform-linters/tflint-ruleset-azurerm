mapping "azurerm_subnet" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/virtualNetwork.json"

  name              = Subnet.name
  address_prefix    = SubnetPropertiesFormat.addressPrefix
  address_prefixes  = SubnetPropertiesFormat.addressPrefixes
  service_endpoints = SubnetPropertiesFormat.serviceEndpoints
}
