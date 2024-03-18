mapping "azurerm_postgresql_database" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name        = DatabaseNameParameter
  server_name = ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}
