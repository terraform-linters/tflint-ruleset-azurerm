mapping "azurerm_monitor_diagnostic_setting" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/diagnosticsSettings_API.json"

  name                       = NameParameter
  log_analytics_workspace_id = DiagnosticSettings.workspaceId
  storage_account_id         = DiagnosticSettings.storageAccountId
}
