mapping "azurerm_stream_analytics_reference_input_blob" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/inputs.json"

  name                      = InputNameParameter
  resource_group_name       = any //ResourceGroupNameParameter
  stream_analytics_job_name = any //StreamingJobNameParameter
  date_format               = BlobDataSourceProperties.dateFormat
  path_pattern              = BlobDataSourceProperties.pathPattern
  storage_account_name      = StorageAccount.accountName
  storage_account_key       = StorageAccount.accountKey
  storage_container_name    = BlobDataSourceProperties.container
  time_format               = BlobDataSourceProperties.timeFormat
}
