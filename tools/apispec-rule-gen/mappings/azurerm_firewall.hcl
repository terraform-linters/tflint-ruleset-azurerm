mapping "azurerm_firewall" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/firewall.json"

  zones = AzureFirewall.zones
}
