mapping "azurerm_mysql_database" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name        = DatabaseNameParameter
  server_name = ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}