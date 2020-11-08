mapping "azurerm_monitor_diagnostic_setting" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-07-01/serviceDiagnosticsSettings_API.json"

  name                       = Resource.name
  log_analytics_workspace_id = ServiceDiagnosticSettings.workspaceId
  storage_account_id         = ServiceDiagnosticSettings.storageAccountId
}
