mapping "azurerm_maintenance_configuration" {
  import_path = "azure-rest-api-specs/specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2022-07-01-preview/Maintenance.json"

  location = MaintenanceConfiguration.location
  scope    = MaintenanceConfigurationProperties.maintenanceScope
}
