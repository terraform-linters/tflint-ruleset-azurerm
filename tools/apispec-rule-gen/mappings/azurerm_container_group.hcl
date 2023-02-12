mapping "azurerm_container_group" {
  import_path = "azure-rest-api-specs/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/containerInstance.json"

  name                = ContainerGroupNameParameter
  resource_group_name = ResourceGroupNameParameter
  location            = LocationParameter
  dns_name_label      = IpAddress.dnsNameLabel
  ip_address_type     = any //IpAddress.type
  network_profile_id  = any //ContainerGroupNetworkProfile.id
}

