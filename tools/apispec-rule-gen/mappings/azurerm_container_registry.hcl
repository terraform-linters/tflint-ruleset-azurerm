mapping "azurerm_container_registry" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2021-08-01-preview/containerregistry.json"

  name                = RegistryNameParameter
  resource_group_name = ResourceGroupParameter
  admin_enabled       = RegistryProperties.adminUserEnabled
  storage_account_id  = any //StorageAccountProperties.id
  sku                 = Sku.name
}