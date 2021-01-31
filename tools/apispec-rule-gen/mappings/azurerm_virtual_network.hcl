mapping "azurerm_virtual_network" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-07-01/virtualNetwork.json"

  dns_servers = DhcpOptions.dnsServers
}
