mapping "azurerm_servicebus_topic" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/common/v1/definitions.json"

  name                                    = any
  namespace_name                          = any
  resource_group_name                     = any
  status                                  = EntityStatus
  auto_delete_on_idle                     = any
  default_message_ttl                     = any
  duplicate_detection_history_time_window = any
  enable_batched_operations               = any
  enable_express                          = any
  enable_partitioning                     = any
  max_size_in_megabytes                   = any
  requires_duplicate_detection            = any
  support_ordering                        = any
}
