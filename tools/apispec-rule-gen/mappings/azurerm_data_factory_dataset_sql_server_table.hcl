mapping "azurerm_data_factory_dataset_sql_server_table" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = datasetName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
  linked_service_name = linkedServiceName
}