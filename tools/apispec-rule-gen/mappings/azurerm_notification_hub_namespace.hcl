mapping "azurerm_notification_hub_namespace" {
  import_path = "azure-rest-api-specs/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/notificationhubs.json"

  name           = NamespaceProperties.name
  namespace_type = NamespaceProperties.namespaceType
  sku_name       = Sku.name
  enabled        = NamespaceProperties.enabled
}
