mapping "azurerm_netapp_pool" {
  import_path = "azure-rest-api-specs/specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/netapp.json"

  name                = PoolName
  resource_group_name = ResourceGroup
  account_name        = AccountName
  location            = location
  service_level       = serviceLevel
}
