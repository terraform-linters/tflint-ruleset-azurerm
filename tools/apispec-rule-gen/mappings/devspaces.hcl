mapping "azurerm_devspace_controller" {
  import_path = "azure-rest-api-specs/specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/devspaces.json"

  name                                     = NameParameter
  resource_group_name                      = ResourceGroupParameter
  sku_name                                 = Sku.name
  target_container_host_resource_id        = ControllerProperties.targetContainerHostResourceId
  target_container_host_credentials_base64 = ControllerProperties.targetContainerHostCredentialsBase64
}
