mapping "azurerm_eventhub_namespace" {
  import_path = "azure-rest-api-specs/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/namespaces-preview.json"

  sku      = Sku.name
  capacity = Sku.capacity
}
