mapping "azurerm_servicebus_subscription" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/subscriptions.json"

  name                                 = any
  namespace_name                       = any
  topic_name                           = any
  resource_group_name                  = any
  max_delivery_count                   = SBSubscriptionProperties.maxDeliveryCount
  auto_delete_on_idle                  = SBSubscriptionProperties.autoDeleteOnIdle
  default_message_ttl                  = SBSubscriptionProperties.defaultMessageTimeToLive
  lock_duration                        = SBSubscriptionProperties.lockDuration
  dead_lettering_on_message_expiration = SBSubscriptionProperties.deadLetteringOnMessageExpiration
  enable_batched_operations            = SBSubscriptionProperties.enableBatchedOperations
  requires_session                     = SBSubscriptionProperties.requiresSession
  forward_to                           = SBSubscriptionProperties.forwardTo
  forward_dead_lettered_messages_to    = SBSubscriptionProperties.forwardDeadLetteredMessagesTo
}
