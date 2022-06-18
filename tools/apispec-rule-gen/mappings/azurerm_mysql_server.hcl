mapping "azurerm_mysql_server" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/mysql.json"

  name                         = ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerProperties.administratorLogin
  administrator_login_password = ServerProperties.administratorLoginPassword
  version                      = any //ServerVersion
  ssl_enforcement              = any //SslEnforcement
}
