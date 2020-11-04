mapping "azurerm_postgresql_server" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name                         = ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerPropertiesForDefaultCreate.administratorLogin
  administrator_login_password = ServerPropertiesForDefaultCreate.administratorLoginPassword
  version                      = ServerVersion
  ssl_enforcement              = SslEnforcement
}