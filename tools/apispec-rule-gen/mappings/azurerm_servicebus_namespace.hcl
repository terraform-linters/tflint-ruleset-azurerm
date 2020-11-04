mapping "azurerm_servicebus_namespace" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2018-01-01-preview/namespace-preview.json"

  name                = any
  resource_group_name = any
  sku                 = SBSku.name
  capacity            = SBSku.capacity
}
