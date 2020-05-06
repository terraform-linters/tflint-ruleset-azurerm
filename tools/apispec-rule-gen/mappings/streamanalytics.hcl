mapping "azurerm_stream_analytics_job" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/streamingjobs.json"

  name                                     = StreamingJobNameParameter
  resource_group_name                      = ResourceGroupNameParameter
  location                                 = Resource.location
  compatibility_level                      = CompatibilityLevel
  data_locale                              = StreamingJobProperties.dataLocale
  events_late_arrival_max_delay_in_seconds = StreamingJobProperties.eventsLateArrivalMaxDelayInSeconds
  events_out_of_order_max_delay_in_seconds = StreamingJobProperties.eventsOutOfOrderMaxDelayInSeconds
  events_out_of_order_policy               = EventsOutOfOrderPolicy
  output_error_policy                      = OutputErrorPolicy
}

mapping "azurerm_stream_analytics_function_javascript_udf" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/functions.json"

  name                      = FunctionNameParameter
  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  script                    = JavaScriptFunctionBindingProperties.script
}

mapping "azurerm_stream_analytics_output_blob" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/inputs.json"

  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  date_format               = BlobDataSourceProperties.dateFormat
  path_pattern              = BlobDataSourceProperties.pathPattern
  storage_account_name      = StorageAccount.accountName
  storage_account_key       = StorageAccount.accountKey
  storage_container_name    = BlobDataSourceProperties.container
  time_format               = BlobDataSourceProperties.timeFormat
}

mapping "azurerm_stream_analytics_output_mssql" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/outputs.json"

  name                      = OutputNameParameter
  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  server                    = AzureSqlDatabaseDataSourceProperties.server
  user                      = AzureSqlDatabaseDataSourceProperties.user
  password                  = AzureSqlDatabaseDataSourceProperties.password
  table                     = AzureSqlDatabaseDataSourceProperties.table
}

mapping "azurerm_stream_analytics_output_eventhub" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/inputs.json"

  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  eventhub_name             = EventHubDataSourceProperties.eventHubName
  servicebus_namespace      = ServiceBusDataSourceProperties.serviceBusNamespace
  shared_access_policy_key  = ServiceBusDataSourceProperties.sharedAccessPolicyKey
  shared_access_policy_name = ServiceBusDataSourceProperties.sharedAccessPolicyName
}

mapping "azurerm_stream_analytics_output_servicebus_queue" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/outputs.json"

  name                      = OutputNameParameter
  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  queue_name                = ServiceBusQueueOutputDataSourceProperties.queueName
}

mapping "azurerm_stream_analytics_output_servicebus_topic" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/outputs.json"

  name                      = OutputNameParameter
  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  queue_name                = ServiceBusTopicOutputDataSourceProperties.topicName
}

mapping "azurerm_stream_analytics_reference_input_blob" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/inputs.json"

  name                      = InputNameParameter
  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  date_format               = BlobDataSourceProperties.dateFormat
  path_pattern              = BlobDataSourceProperties.pathPattern
  storage_account_name      = StorageAccount.accountName
  storage_account_key       = StorageAccount.accountKey
  storage_container_name    = BlobDataSourceProperties.container
  time_format               = BlobDataSourceProperties.timeFormat
}

mapping "azurerm_stream_analytics_stream_input_blob" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/inputs.json"

  name                      = InputNameParameter
  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  date_format               = BlobDataSourceProperties.dateFormat
  path_pattern              = BlobDataSourceProperties.pathPattern
  storage_account_name      = StorageAccount.accountName
  storage_account_key       = StorageAccount.accountKey
  storage_container_name    = BlobDataSourceProperties.container
  time_format               = BlobDataSourceProperties.timeFormat
}

mapping "azurerm_stream_analytics_stream_input_eventhub" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/inputs.json"

  name                         = InputNameParameter
  resource_group_name          = ResourceGroupNameParameter
  stream_analytics_job_name    = StreamingJobNameParameter
  eventhub_name                = EventHubDataSourceProperties.eventHubName
  eventhub_consumer_group_name = EventHubStreamInputDataSourceProperties.consumerGroupName
  servicebus_namespace         = ServiceBusDataSourceProperties.serviceBusNamespace
  shared_access_policy_key     = ServiceBusDataSourceProperties.sharedAccessPolicyKey
  shared_access_policy_name    = ServiceBusDataSourceProperties.sharedAccessPolicyName
}

mapping "azurerm_stream_analytics_stream_input_iothub" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/inputs.json"

  name                         = InputNameParameter
  resource_group_name          = ResourceGroupNameParameter
  stream_analytics_job_name    = StreamingJobNameParameter
  eventhub_consumer_group_name = IoTHubStreamInputDataSourceProperties.consumerGroupName
  endpoint                     = IoTHubStreamInputDataSourceProperties.endpoint
  iothub_namespace             = IoTHubStreamInputDataSourceProperties.iotHubNamespace
  shared_access_policy_key     = IoTHubStreamInputDataSourceProperties.sharedAccessPolicyKey
  shared_access_policy_name    = IoTHubStreamInputDataSourceProperties.sharedAccessPolicyName
}
