mapping "azurerm_eventhub_namespace" {
  import_path = "azure-rest-api-specs/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2018-01-01-preview/namespaces-preview.json"

  sku      = Sku.name
  capacity = Sku.capacity
}
