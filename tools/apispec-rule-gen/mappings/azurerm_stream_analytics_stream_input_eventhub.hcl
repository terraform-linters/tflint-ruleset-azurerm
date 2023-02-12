mapping "azurerm_stream_analytics_stream_input_eventhub" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/inputs.json"

  name                         = InputNameParameter
  resource_group_name          = any //ResourceGroupNameParameter
  stream_analytics_job_name    = any //StreamingJobNameParameter
  eventhub_name                = EventHubDataSourceProperties.eventHubName
  eventhub_consumer_group_name = EventHubStreamInputDataSourceProperties.consumerGroupName
  servicebus_namespace         = ServiceBusDataSourceProperties.serviceBusNamespace
  shared_access_policy_key     = ServiceBusDataSourceProperties.sharedAccessPolicyKey
  shared_access_policy_name    = ServiceBusDataSourceProperties.sharedAccessPolicyName
}
