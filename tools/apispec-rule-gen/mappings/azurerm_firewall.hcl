mapping "azurerm_firewall" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/azureFirewall.json"

  zones = AzureFirewall.zones
}
