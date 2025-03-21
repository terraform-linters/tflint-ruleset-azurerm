mapping "azurerm_container_registry_webhook" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/containerregistry.json"

  name                = WebhookNameParameter
  resource_group_name = any //ResourceGroupParameter
  registry_name       = RegistryNameParameter
  service_uri         = WebhookPropertiesCreateParameters.serviceUri
  actions             = WebhookProperties.actions
  status              = WebhookProperties.status
  scope               = WebhookProperties.scope
}
