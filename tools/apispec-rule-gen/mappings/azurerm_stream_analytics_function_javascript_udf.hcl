mapping "azurerm_stream_analytics_function_javascript_udf" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/functions.json"

  name                      = FunctionNameParameter
  resource_group_name       = ResourceGroupNameParameter
  stream_analytics_job_name = StreamingJobNameParameter
  script                    = JavaScriptFunctionBindingProperties.script
}