mapping "azurerm_netapp_account" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/netapp.json"

  name                = AccountName
  resource_group_name = any //ResourceGroup
  location            = any //location
}
