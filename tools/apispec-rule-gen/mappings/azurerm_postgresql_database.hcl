mapping "azurerm_postgresql_database" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/Databases.json"

  name        = DatabaseNameParameter
  server_name = any //ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}
