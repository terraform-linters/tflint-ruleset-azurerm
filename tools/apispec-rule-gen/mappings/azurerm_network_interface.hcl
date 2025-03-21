mapping "azurerm_network_interface" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/networkInterface.json"

  dns_servers                   = NetworkInterfaceDnsSettings.dnsServers
  enable_ip_forwarding          = NetworkInterfacePropertiesFormat.enableIPForwarding
  enable_accelerated_networking = NetworkInterfacePropertiesFormat.enableAcceleratedNetworking
  internal_dns_name_label       = NetworkInterfaceDnsSettings.internalDnsNameLabel
}
