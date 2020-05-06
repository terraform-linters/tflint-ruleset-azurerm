mapping "azurerm_log_analytics_workspace" {
  import_path = "azure-rest-api-specs/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2020-03-01-preview/Workspaces.json"

  sku               = WorkspaceSku.name
  retention_in_days = WorkspaceProperties.retentionInDays
}
