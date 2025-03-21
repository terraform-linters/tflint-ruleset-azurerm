mapping "azurerm_servicebus_subscription_rule" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-06-01-preview/Rules.json"

  namespace_name      = any
  topic_name          = any
  subscription_name   = any
  resource_group_name = any
  filter_type         = FilterType
}
