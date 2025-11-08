mapping "azurerm_data_factory_trigger_schedule" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/entityTypes/Trigger.json"

  start_time = ScheduleTriggerRecurrence.startTime
  end_time   = ScheduleTriggerRecurrence.endTime
  interval   = ScheduleTriggerRecurrence.interval
  frequency  = RecurrenceFrequency
}
