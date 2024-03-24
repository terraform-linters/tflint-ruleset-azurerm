mapping "azurerm_monitor_activity_log_alert" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/activityLogAlerts_API.json"

  name                = ActivityLogAlertNameParameter
  resource_group_name = any //ResourceGroupNameParameter
  scopes              = AlertRuleProperties.scopes
  enabled             = AlertRuleProperties.enabled
  description         = AlertRuleProperties.description
}
