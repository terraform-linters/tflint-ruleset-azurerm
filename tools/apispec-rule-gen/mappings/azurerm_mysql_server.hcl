mapping "azurerm_mysql_server" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2017-12-01/mysql.json"

  name                         = ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerPropertiesForDefaultCreate.administratorLogin
  administrator_login_password = ServerPropertiesForDefaultCreate.administratorLoginPassword
  version                      = ServerVersion
}
