mapping "azurerm_servicebus_namespace" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/stable/2026-01-01/servicebus.json"

  name                = any
  resource_group_name = any
  sku                 = SBSku.name
  capacity            = SBSku.capacity
}
