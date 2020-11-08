mapping "azurerm_monitor_activity_log_alert" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2017-04-01/activityLogAlerts_API.json"

  name                = ActivityLogAlertNameParameter
  resource_group_name = ResourceGroupNameParameter
  scopes              = ActivityLogAlert.scopes
  enabled             = ActivityLogAlert.enabled
  description         = ActivityLogAlert.description
}
