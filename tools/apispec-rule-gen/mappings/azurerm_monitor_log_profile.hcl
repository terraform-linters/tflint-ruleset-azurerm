mapping "azurerm_monitor_log_profile" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/Insights/stable/2016-03-01/logProfiles.json"

  name               = any //LogProfileNameParameter
  categories         = LogProfileProperties.categories
  locations          = LogProfileProperties.locations
  storage_account_id = LogProfileProperties.storageAccountId
  servicebus_rule_id = LogProfileProperties.serviceBusRuleId
}
