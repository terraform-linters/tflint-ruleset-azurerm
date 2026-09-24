mapping "azurerm_cdn_profile" {
  import_path = "azure-rest-api-specs/specification/cdn/resource-manager/Microsoft.Cdn/Cdn/stable/2025-12-01/openapi.json"

  resource_group_name = ResourceGroupNameParameter
  sku                 = Sku.name
}
