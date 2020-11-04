mapping "azurerm_monitor_scheduled_query_rules_alert" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/scheduledQueryRule_API.json"

  name                 = RuleNameParameter
  resource_group_name  = ResourceGroupNameParameter
  data_source_id       = Source.dataSourceId
  frequency            = Schedule.frequencyInMinutes
  query                = Source.query
  time_window          = Schedule.timeWindowInMinutes
  description          = LogSearchRule.description
  enabled              = any // LogSearchRule.enabled
  severity             = AlertSeverity
  throttling           = AlertingAction.throttlingInMin
}
