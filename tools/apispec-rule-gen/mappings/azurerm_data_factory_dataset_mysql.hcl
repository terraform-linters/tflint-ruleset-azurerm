mapping "azurerm_data_factory_dataset_mysql" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/openapi.json"

  name                = any // datasetName is now an inline path parameter.
  linked_service_name = any // linkedServiceName is now an inline path parameter.
}
