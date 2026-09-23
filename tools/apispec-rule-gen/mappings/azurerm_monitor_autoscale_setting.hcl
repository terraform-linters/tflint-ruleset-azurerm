mapping "azurerm_monitor_autoscale_setting" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/Insights/stable/2022-10-01/autoScale.json"

  name                = any //AutoscaleSettingNameParameter
  resource_group_name = any //ResourceGroupNameParameter
  location            = any //Resource.location
  target_resource_id  = AutoscaleSetting.targetResourceUri
  enabled             = AutoscaleSetting.enabled
}
