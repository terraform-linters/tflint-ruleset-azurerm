mapping "azurerm_automation_account" {
  import_path = "azure-rest-api-specs/specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/account.json"

  sku_name = Sku.name
}

mapping "azurerm_automation_runbook" {
  import_path = "azure-rest-api-specs/specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/runbook.json"

  runbook_type = RunbookProperties.runbookType
}

mapping "azurerm_automation_schedule" {
  import_path = "azure-rest-api-specs/specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/schedule.json"

  frequency = scheduleFrequency
}
