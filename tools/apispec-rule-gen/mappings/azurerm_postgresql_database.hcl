mapping "azurerm_postgresql_database" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/Databases.json"

  name        = DatabaseNameParameter
  server_name = any //ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}
