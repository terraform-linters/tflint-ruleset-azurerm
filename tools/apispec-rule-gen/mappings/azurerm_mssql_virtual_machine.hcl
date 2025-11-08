
mapping "azurerm_mssql_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/SqlVirtualMachine/stable/2023-10-01/sqlvm.json"

  sql_license_type                 = SqlServerLicenseType
  r_services_enabled               = AdditionalFeaturesServerConfigurations.isRServicesEnabled
  sql_connectivity_port            = SqlConnectivityUpdateSettings.port
  sql_connectivity_type            = ConnectivityType
  sql_connectivity_update_password = SqlConnectivityUpdateSettings.sqlAuthUpdatePassword
  sql_connectivity_update_username = SqlConnectivityUpdateSettings.sqlAuthUpdateUserName
}
