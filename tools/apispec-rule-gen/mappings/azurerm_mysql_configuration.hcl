mapping "azurerm_mysql_configuration" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/mysql.json"

  name        = ConfigurationNameParameter
  server_name = ServerNameParameter
  value       = ConfigurationProperties.value
}