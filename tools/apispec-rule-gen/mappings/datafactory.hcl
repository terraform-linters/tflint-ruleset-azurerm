mapping "azurerm_data_factory" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = factoryName
  resource_group_name = resourceGroupName
}

mapping "azurerm_data_factory_dataset_mysql" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = datasetName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
  linked_service_name = linkedServiceName
}

mapping "azurerm_data_factory_dataset_postgresql" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = datasetName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
  linked_service_name = linkedServiceName
}

mapping "azurerm_data_factory_dataset_sql_server_table" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = datasetName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
  linked_service_name = linkedServiceName
}

mapping "azurerm_data_factory_integration_runtime_managed" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/entityTypes/IntegrationRuntime.json"

  node_size                        = IntegrationRuntimeComputeProperties.nodeSize
  number_of_nodes                  = IntegrationRuntimeComputeProperties.numberOfNodes
  max_parallel_executions_per_node = IntegrationRuntimeComputeProperties.maxParallelExecutionsPerNode
  edition                          = IntegrationRuntimeSsisProperties.edition
  license_type                     = IntegrationRuntimeSsisProperties.licenseType
}

mapping "azurerm_data_factory_pipeline" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = pipelineName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
}

mapping "azurerm_data_factory_linked_service_data_lake_storage_gen2" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = linkedServiceName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
}

mapping "azurerm_data_factory_linked_service_mysql" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = linkedServiceName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
}

mapping "azurerm_data_factory_linked_service_postgresql" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = linkedServiceName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
}

mapping "azurerm_data_factory_linked_service_sql_server" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/datafactory.json"

  name                = linkedServiceName
  resource_group_name = resourceGroupName
  data_factory_name   = factoryName
}

mapping "azurerm_data_factory_trigger_schedule" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/entityTypes/Trigger.json"

  start_time = ScheduleTriggerRecurrence.startTime
  end_time   = ScheduleTriggerRecurrence.endTime
  interval   = ScheduleTriggerRecurrence.interval
  frequency  = RecurrenceFrequency
}
