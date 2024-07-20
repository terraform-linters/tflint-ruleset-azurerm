mapping "azurerm_maps_account" {
  import_path = "azure-rest-api-specs/specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/maps-management.json"

  name                = AccountNameParameter
  resource_group_name = any //ResourceGroupNameParameter
  sku_name            = Sku.name
}
