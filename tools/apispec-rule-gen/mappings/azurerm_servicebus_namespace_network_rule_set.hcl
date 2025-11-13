mapping "azurerm_servicebus_namespace_network_rule_set" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2022-10-01-preview/namespace-preview.json"

  resource_group_name = any
  namespace_name      = any
}
