mapping "azurerm_bastion_host" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/bastionHost.json"

  name                = BastionHostName
  resource_group_name = ResourceGroupName
}
