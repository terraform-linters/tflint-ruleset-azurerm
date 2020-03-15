mapping "azurerm_cost_management_export_resource_group" {
  import_path = "azure-rest-api-specs/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2019-11-01/costmanagement.json"

  name                    = exportNameParameter
  recurrence_type         = ExportSchedule.recurrence
  recurrence_period_start = ExportRecurrencePeriod.from
  recurrence_period_end   = ExportRecurrencePeriod.to
}
