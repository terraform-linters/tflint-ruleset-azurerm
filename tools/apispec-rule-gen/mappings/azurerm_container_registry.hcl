mapping "azurerm_container_registry" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-04-01/containerregistry.json"

  name                = RegistryNameParameter
  resource_group_name = any //ResourceGroupParameter
  admin_enabled       = RegistryProperties.adminUserEnabled
  storage_account_id  = any //StorageAccountProperties.id
  sku                 = Sku.name
}
