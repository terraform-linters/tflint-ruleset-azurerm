mapping "azurerm_monitor_autoscale_setting" {
  import_path = "azure-rest-api-specs/specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/autoscale_API.json"

  name                = AutoscaleSettingNameParameter
  resource_group_name = any //ResourceGroupNameParameter
  location            = Resource.location
  target_resource_id  = AutoscaleSetting.targetResourceUri
  enabled             = AutoscaleSetting.enabled
}
