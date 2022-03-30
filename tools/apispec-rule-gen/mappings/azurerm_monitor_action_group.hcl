mapping "azurerm_monitor_action_group" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2019-06-01/actionGroups_API.json"

  name                = ActionGroupNameParameter
  resource_group_name = any //ResourceGroupNameParameter
  short_name          = ActionGroup.groupShortName
  enabled             = ActionGroup.enabled
}
