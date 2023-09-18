mapping "azurerm_postgresql_configuration" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/Configuration.json"

  name        = ConfigurationNameParameter
  server_name = any //ServerNameParameter
  value       = ConfigurationProperties.value
}
