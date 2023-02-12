mapping "azurerm_stream_analytics_function_javascript_udf" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/functions.json"

  name                      = FunctionNameParameter
  resource_group_name       = any //ResourceGroupNameParameter
  stream_analytics_job_name = any //StreamingJobNameParameter
  script                    = JavaScriptFunctionBindingProperties.script
}