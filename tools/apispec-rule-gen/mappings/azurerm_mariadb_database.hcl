mapping "azurerm_mariadb_database" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name        = DatabaseNameParameter
  server_name = ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}
