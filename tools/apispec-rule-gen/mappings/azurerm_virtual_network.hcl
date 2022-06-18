mapping "azurerm_virtual_network" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/virtualNetwork.json"

  dns_servers = DhcpOptions.dnsServers
}
