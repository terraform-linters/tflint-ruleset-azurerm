mapping "azurerm_monitor_action_group" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2019-06-01/actionGroups_API.json"

  name                = ActionGroupNameParameter
  resource_group_name = ResourceGroupNameParameter
  short_name          = ActionGroup.groupShortName
  enabled             = ActionGroup.enabled
}

mapping "azurerm_monitor_activity_log_alert" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/activityLogAlerts_API.json"

  name                = ActivityLogAlertNameParameter
  resource_group_name = ResourceGroupNameParameter
  scopes              = ActivityLogAlert.scopes
  enabled             = ActivityLogAlert.enabled
  description         = ActivityLogAlert.description
}

mapping "azurerm_monitor_autoscale_setting" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/autoscale_API.json"

  name                = AutoscaleSettingNameParameter
  resource_group_name = ResourceGroupNameParameter
  location            = Resource.location
  target_resource_id  = AutoscaleSetting.targetResourceUri
  enabled             = AutoscaleSetting.enabled
}

mapping "azurerm_monitor_diagnostic_setting" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-07-01/serviceDiagnosticsSettings_API.json"

  name                       = Resource.name
  log_analytics_workspace_id = ServiceDiagnosticSettings.workspaceId
  storage_account_id         = ServiceDiagnosticSettings.storageAccountId
}

mapping "azurerm_monitor_log_profile" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/logProfiles_API.json"

  name               = LogProfileNameParameter
  categories         = LogProfileProperties.categories
  locations          = LogProfileProperties.locations
  storage_account_id = LogProfileProperties.storageAccountId
  servicebus_rule_id = LogProfileProperties.serviceBusRuleId
}

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

mapping "azurerm_monitor_scheduled_query_rules_log" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2018-04-16/scheduledQueryRule_API.json"

  name                = RuleNameParameter
  resource_group_name = ResourceGroupNameParameter
  data_source_id      = Source.dataSourceId
  description         = LogSearchRule.description
  enabled             = any // LogSearchRule.enabled
}
