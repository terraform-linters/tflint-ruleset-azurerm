mapping "azurerm_eventgrid_domain" {
  import_path = "azure-rest-api-specs/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2019-02-01-preview/EventGrid.json"

  input_schema = DomainProperties.inputSchema
}

mapping "azurerm_eventgrid_event_subscription" {
  import_path = "azure-rest-api-specs/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2019-02-01-preview/EventGrid.json"

  event_delivery_schema = EventSubscriptionProperties.eventDeliverySchema
  topic_name            = EventSubscriptionProperties.topic
  included_event_types  = EventSubscriptionFilter.includedEventTypes
  labels                = EventSubscriptionProperties.labels
}

mapping "azurerm_eventhub_namespace" {
  import_path = "azure-rest-api-specs/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2017-04-01/namespaces.json"

  sku      = Sku.name
  capacity = Sku.capacity
}

mapping "azurerm_notification_hub_namespace" {
  import_path = "azure-rest-api-specs/specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/notificationhubs.json"

  name           = NamespaceProperties.name
  namespace_type = NamespaceProperties.namespaceType
  sku_name       = Sku.name
  enabled        = NamespaceProperties.enabled
}

mapping "azurerm_relay_hybrid_connection" {
  import_path = "azure-rest-api-specs/specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/relay.json"

  name                          = hybridConnectionNameParameter
  resource_group_name           = resourceGroupNameParameter
  relay_namespace_name          = namespaceNameParameter
}

mapping "azurerm_relay_namespace" {
  import_path = "azure-rest-api-specs/specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/relay.json"

  name                = namespaceNameParameter
  resource_group_name = resourceGroupNameParameter
  sku_name            = Sku.name
}

mapping "azurerm_servicebus_namespace" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  name                = NamespaceNameParameter
  resource_group_name = ResourceGroupNameParameter
  sku                 = SBSku.name
  capacity            = SBSku.capacity
}

mapping "azurerm_servicebus_namespace_authorization_rule" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  name                = RuleNameParameter
  namespace_name      = NamespaceNameParameter
  resource_group_name = ResourceGroupNameParameter
}

mapping "azurerm_servicebus_namespace_network_rule_set" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  resource_group_name = ResourceGroupNameParameter
  namespace_name      = NamespaceNameParameter
}

mapping "azurerm_servicebus_queue" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  name                                    = QueueNameParameter
  namespace_name                          = NamespaceNameParameter
  resource_group_name                     = ResourceGroupNameParameter
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

mapping "azurerm_servicebus_queue_authorization_rule" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  name                = AuthorizationRuleNameParameter
  namespace_name      = NamespaceNameParameter
  queue_name          = QueueNameParameter
  resource_group_name = ResourceGroupNameParameter
}

mapping "azurerm_servicebus_subscription" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  name                                 = SubscriptionNameParameter
  namespace_name                       = NamespaceNameParameter
  topic_name                           = TopicNameParameter
  resource_group_name                  = ResourceGroupNameParameter
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

mapping "azurerm_servicebus_subscription_rule" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  namespace_name      = NamespaceNameParameter
  topic_name          = TopicNameParameter
  subscription_name   = SubscriptionNameParameter
  resource_group_name = ResourceGroupNameParameter
  filter_type         = FilterType
}

mapping "azurerm_servicebus_topic" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  name                                    = TopicNameParameter
  namespace_name                          = NamespaceNameParameter
  resource_group_name                     = ResourceGroupNameParameter
  status                                  = EntityStatus
  auto_delete_on_idle                     = SBTopicProperties.autoDeleteOnIdle
  default_message_ttl                     = SBTopicProperties.defaultMessageTimeToLive
  duplicate_detection_history_time_window = SBTopicProperties.duplicateDetectionHistoryTimeWindow
  enable_batched_operations               = SBTopicProperties.enableBatchedOperations
  enable_express                          = SBTopicProperties.enableExpress
  enable_partitioning                     = SBTopicProperties.enablePartitioning
  max_size_in_megabytes                   = SBTopicProperties.maxSizeInMegabytes
  requires_duplicate_detection            = SBTopicProperties.requiresDuplicateDetection
  support_ordering                        = SBTopicProperties.supportOrdering
}

mapping "azurerm_servicebus_topic_authorization_rule" {
  import_path = "azure-rest-api-specs/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2017-04-01/servicebus.json"

  name                = AuthorizationRuleNameParameter
  namespace_name      = NamespaceNameParameter
  topic_name          = TopicNameParameter
  resource_group_name = ResourceGroupNameParameter
}

mapping "azurerm_signalr_service" {
  import_path = "azure-rest-api-specs/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2018-10-01/signalr.json"

  name                = SignalRServiceName
  resource_group_name = ResourceGroupParameter
}
