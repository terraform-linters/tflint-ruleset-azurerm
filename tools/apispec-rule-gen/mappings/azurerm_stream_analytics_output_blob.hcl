mapping "azurerm_stream_analytics_output_blob" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/StreamAnalytics/preview/2021-10-01-preview/outputs.json"

  resource_group_name       = any //ResourceGroupNameParameter
  stream_analytics_job_name = any //StreamingJobNameParameter
  date_format               = any // BlobOutputDataSourceProperties inherits these from inputs.json.
  path_pattern              = any // BlobOutputDataSourceProperties inherits these from inputs.json.
  storage_account_name      = any // BlobDataSourceProperties.storageAccounts[] is not representable as a single property path.
  storage_account_key       = any // BlobDataSourceProperties.storageAccounts[] is not representable as a single property path.
  storage_container_name    = any // BlobOutputDataSourceProperties inherits these from inputs.json.
  time_format               = any // BlobOutputDataSourceProperties inherits these from inputs.json.
}
