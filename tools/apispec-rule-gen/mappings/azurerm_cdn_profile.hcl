mapping "azurerm_cdn_profile" {
  import_path = "azure-rest-api-specs/specification/cdn/resource-manager/Microsoft.Cdn/stable/2019-04-15/cdn.json"

  resource_group_name = resourceGroupNameParameter
  sku                 = Sku.name
}
