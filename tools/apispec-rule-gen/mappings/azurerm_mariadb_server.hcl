mapping "azurerm_mariadb_server" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name                         = ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerPropertiesForDefaultCreate.administratorLogin
  administrator_login_password = ServerPropertiesForDefaultCreate.administratorLoginPassword
  version                      = ServerVersion
  ssl_enforcement              = SslEnforcement
}