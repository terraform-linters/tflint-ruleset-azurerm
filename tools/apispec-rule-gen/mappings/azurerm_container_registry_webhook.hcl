mapping "azurerm_container_registry_webhook" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/Registry/stable/2025-11-01/containerregistry.json"

  name                = any // webhookName is now an inline path parameter.
  resource_group_name = any //ResourceGroupParameter
  registry_name       = any // registryName is now an inline path parameter.
  service_uri         = WebhookPropertiesCreateParameters.serviceUri
  actions             = WebhookProperties.actions
  status              = WebhookStatus
  scope               = WebhookProperties.scope
}
