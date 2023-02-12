mapping "azurerm_postgresql_configuration" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/Configuration.json"

  name        = ConfigurationNameParameter
  server_name = any //ServerNameParameter
  value       = ConfigurationProperties.value
}
