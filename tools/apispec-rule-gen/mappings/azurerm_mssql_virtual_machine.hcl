
mapping "azurerm_mssql_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/sqlvm.json"

  sql_license_type                 = SqlVirtualMachineProperties.sqlServerLicenseType
  r_services_enabled               = AdditionalFeaturesServerConfigurations.isRServicesEnabled
  sql_connectivity_port            = SqlConnectivityUpdateSettings.port
  sql_connectivity_type            = SqlConnectivityUpdateSettings.connectivityType
  sql_connectivity_update_password = SqlConnectivityUpdateSettings.sqlAuthUpdatePassword
  sql_connectivity_update_username = SqlConnectivityUpdateSettings.sqlAuthUpdateUserName
}
