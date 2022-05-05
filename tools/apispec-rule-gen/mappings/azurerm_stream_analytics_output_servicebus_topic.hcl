mapping "azurerm_stream_analytics_output_servicebus_topic" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/outputs.json"

  name                      = OutputNameParameter
  resource_group_name       = any //ResourceGroupNameParameter
  stream_analytics_job_name = any //StreamingJobNameParameter
  queue_name                = ServiceBusTopicOutputDataSourceProperties.topicName
}