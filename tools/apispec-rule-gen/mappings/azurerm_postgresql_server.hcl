mapping "azurerm_postgresql_server" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/FlexibleServers.json"

  name                         = any //ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerProperties.administratorLogin
  administrator_login_password = ServerProperties.administratorLoginPassword
  version                      = ServerVersion
}
