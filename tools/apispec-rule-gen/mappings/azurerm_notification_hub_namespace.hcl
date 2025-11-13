mapping "azurerm_notification_hub_namespace" {
  import_path = "azure-rest-api-specs/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/NotificationHubs/stable/2023-09-01/notificationhubs.json"

  name           = NamespaceProperties.name
  namespace_type = NamespaceType
  sku_name       = SkuName
  enabled        = NamespaceProperties.enabled
}
