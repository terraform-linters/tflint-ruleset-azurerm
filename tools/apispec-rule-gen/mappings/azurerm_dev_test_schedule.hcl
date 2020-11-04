mapping "azurerm_dev_test_schedule" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  location            = locationName
  resource_group_name = resourceGroupName
  status              = ScheduleProperties.status
  task_type           = ScheduleProperties.taskType
  time_zone_id        = ScheduleProperties.timeZoneId
}