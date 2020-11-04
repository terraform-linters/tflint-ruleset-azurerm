mapping "azurerm_mariadb_configuration" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name        = ConfigurationNameParameter
  server_name = ServerNameParameter
  value       = ConfigurationProperties.value
}
