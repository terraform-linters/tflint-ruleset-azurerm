mapping "azurerm_bastion_host" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/Network/stable/2025-07-01/virtualNetwork.json"

  name                = bastionHostName
  resource_group_name = ResourceGroupNameParameter
}
