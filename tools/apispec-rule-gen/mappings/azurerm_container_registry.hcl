mapping "azurerm_container_registry" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2019-05-01/containerregistry.json"

  name                = RegistryNameParameter
  resource_group_name = ResourceGroupParameter
  admin_enabled       = RegistryProperties.adminUserEnabled
  storage_account_id  = StorageAccountProperties.id
  sku                 = Sku.name
}