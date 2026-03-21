mapping "azurerm_container_registry" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/containerregistry.json"

  name                = any // registryName is now an inline path parameter.
  resource_group_name = any //ResourceGroupParameter
  admin_enabled       = RegistryProperties.adminUserEnabled
  storage_account_id  = any //StorageAccountProperties.id
  sku                 = SkuName
}
