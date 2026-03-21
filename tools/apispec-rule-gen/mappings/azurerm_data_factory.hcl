mapping "azurerm_data_factory" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/openapi.json"

  name                = any // factoryName is now an inline path parameter.
  resource_group_name = any // resourceGroupName is now an inline path parameter.
}
