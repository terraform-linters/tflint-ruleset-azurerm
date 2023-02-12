mapping "azurerm_stream_analytics_stream_input_iothub" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/inputs.json"

  name                         = InputNameParameter
  resource_group_name          = any //ResourceGroupNameParameter
  stream_analytics_job_name    = any //StreamingJobNameParameter
  eventhub_consumer_group_name = IoTHubStreamInputDataSourceProperties.consumerGroupName
  endpoint                     = IoTHubStreamInputDataSourceProperties.endpoint
  iothub_namespace             = IoTHubStreamInputDataSourceProperties.iotHubNamespace
  shared_access_policy_key     = IoTHubStreamInputDataSourceProperties.sharedAccessPolicyKey
  shared_access_policy_name    = IoTHubStreamInputDataSourceProperties.sharedAccessPolicyName
}
