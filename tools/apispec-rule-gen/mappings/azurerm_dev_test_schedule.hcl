mapping "azurerm_dev_test_schedule" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/DevTestLabs/stable/2018-09-15/DTL.json"

  location            = any //locationName
  resource_group_name = any //resourceGroupName
  status              = ScheduleProperties.status
  task_type           = ScheduleProperties.taskType
  time_zone_id        = ScheduleProperties.timeZoneId
}
