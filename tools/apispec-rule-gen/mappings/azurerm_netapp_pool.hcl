mapping "azurerm_netapp_pool" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/netapp.json"

  name                = PoolName
  resource_group_name = any //ResourceGroup
  account_name        = AccountName
  location            = any //location
  service_level       = serviceLevel
}
