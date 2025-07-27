mapping "azurerm_servicebus_queue" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/Queue.json"

  name                                    = any
  namespace_name                          = any
  resource_group_name                     = any
  auto_delete_on_idle                     = SBQueueProperties.autoDeleteOnIdle
  default_message_ttl                     = SBQueueProperties.defaultMessageTimeToLive
  duplicate_detection_history_time_window = SBQueueProperties.duplicateDetectionHistoryTimeWindow
  enable_express                          = SBQueueProperties.enableExpress
  enable_partitioning                     = SBQueueProperties.enablePartitioning
  lock_duration                           = SBQueueProperties.lockDuration
  max_size_in_megabytes                   = SBQueueProperties.maxSizeInMegabytes
  requires_duplicate_detection            = SBQueueProperties.requiresDuplicateDetection
  requires_session                        = SBQueueProperties.requiresSession
  dead_lettering_on_message_expiration    = SBQueueProperties.deadLetteringOnMessageExpiration
  max_delivery_count                      = SBQueueProperties.maxDeliveryCount
}
