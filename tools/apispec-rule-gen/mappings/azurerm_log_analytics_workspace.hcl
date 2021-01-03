mapping "azurerm_log_analytics_workspace" {
  import_path = "azure-rest-api-specs/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/Workspaces.json"

  sku               = any
  retention_in_days = WorkspaceProperties.retentionInDays
}
