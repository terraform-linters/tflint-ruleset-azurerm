mapping "azurerm_mysql_configuration" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name        = ConfigurationNameParameter
  server_name = ServerNameParameter
  value       = ConfigurationProperties.value
}