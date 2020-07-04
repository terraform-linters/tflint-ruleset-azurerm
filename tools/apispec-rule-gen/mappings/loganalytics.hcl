mapping "azurerm_log_analytics_workspace" {
  import_path = "azure-rest-api-specs/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2015-11-01-preview/OperationalInsights.json"

  sku               = any
  retention_in_days = WorkspaceProperties.retentionInDays
}
