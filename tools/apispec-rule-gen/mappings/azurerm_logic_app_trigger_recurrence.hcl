mapping "azurerm_logic_app_trigger_recurrence" {
  import_path = "azure-rest-api-specs/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/logic.json"

  frequency  = RecurrenceFrequency
  interval   = WorkflowTriggerRecurrence.interval
  start_time = WorkflowTriggerRecurrence.startTime
}
