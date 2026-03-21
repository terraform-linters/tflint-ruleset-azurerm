mapping "azurerm_stream_analytics_output_eventhub" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/StreamAnalytics/preview/2021-10-01-preview/outputs.json"

  resource_group_name       = any //ResourceGroupNameParameter
  stream_analytics_job_name = any //StreamingJobNameParameter
  eventhub_name             = any // The output spec references EventHubDataSourceProperties from inputs.json.
  servicebus_namespace      = any // The output spec references ServiceBusDataSourceProperties from inputs.json.
  shared_access_policy_key  = any // The output spec references ServiceBusDataSourceProperties from inputs.json.
  shared_access_policy_name = any // The output spec references ServiceBusDataSourceProperties from inputs.json.
}
