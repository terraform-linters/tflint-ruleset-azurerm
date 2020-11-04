mapping "azurerm_container_registry_webhook" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2019-05-01/containerregistry.json"

  name                = WebhookNameParameter
  resource_group_name = ResourceGroupParameter
  registry_name       = RegistryNameParameter
  service_uri         = WebhookPropertiesCreateParameters.serviceUri
  actions             = WebhookProperties.actions
  status              = WebhookProperties.status
  scope               = WebhookProperties.scope
}