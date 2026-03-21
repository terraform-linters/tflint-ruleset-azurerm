mapping "azurerm_container_group" {
  import_path = "azure-rest-api-specs/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/ContainerInstance/stable/2025-09-01/containerInstance.json"

  name                = ContainerGroupNameParameter
  resource_group_name = any // ResourceGroupNameParameter is now referenced externally.
  location            = any // LocationParameter is now referenced externally.
  dns_name_label      = IpAddress.dnsNameLabel
  ip_address_type     = any //IpAddress.type
  network_profile_id  = any //ContainerGroupNetworkProfile.id
}
