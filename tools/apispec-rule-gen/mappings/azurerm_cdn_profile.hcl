mapping "azurerm_cdn_profile" {
  import_path = "azure-rest-api-specs/specification/cdn/resource-manager/Microsoft.Cdn/Cdn/stable/2020-09-01/cdn.json"

  resource_group_name = resourceGroupNameParameter
  sku                 = Sku.name
}
