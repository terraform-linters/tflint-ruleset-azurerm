mapping "azurerm_data_factory_pipeline" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = pipelineName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
}