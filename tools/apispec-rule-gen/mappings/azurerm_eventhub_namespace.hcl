mapping "azurerm_eventhub_namespace" {
  import_path = "azure-rest-api-specs/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/namespaces.json"

  sku      = Sku.name
  capacity = Sku.capacity
}
