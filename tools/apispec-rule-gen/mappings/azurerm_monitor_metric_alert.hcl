mapping "azurerm_monitor_metric_alert" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/metricAlert_API.json"

  name                = RuleNameParameter
  resource_group_name = ResourceGroupNameParameter
  scopes              = MetricAlertProperties.scopes
  enabled             = MetricAlertProperties.enabled
  auto_mitigate       = MetricAlertProperties.autoMitigate
  description         = MetricAlertProperties.description
  frequency           = MetricAlertProperties.evaluationFrequency
  severity            = MetricAlertProperties.severity
  window_size         = MetricAlertProperties.windowSize
}
