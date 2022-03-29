mapping "azurerm_sentinel_alert_rule_ms_security_incident" {
  import_path = "azure-rest-api-specs/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/AlertRules.json"

  name            = MicrosoftSecurityIncidentCreationAlertRuleProperties.alertRuleTemplateName
  display_name    = MicrosoftSecurityIncidentCreationAlertRuleProperties.displayName
  product_filter  = MicrosoftSecurityIncidentCreationAlertRuleCommonProperties.productFilter
  severity_filter = any // AlertSeverity
  description     = MicrosoftSecurityIncidentCreationAlertRuleProperties.description
  enabled         = MicrosoftSecurityIncidentCreationAlertRuleProperties.enabled
}
