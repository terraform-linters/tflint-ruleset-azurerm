mapping "azurerm_data_factory_integration_runtime_managed" {
  import_path = "azure-rest-api-specs/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/entityTypes/IntegrationRuntime.json"

  node_size                        = IntegrationRuntimeComputeProperties.nodeSize
  number_of_nodes                  = IntegrationRuntimeComputeProperties.numberOfNodes
  max_parallel_executions_per_node = IntegrationRuntimeComputeProperties.maxParallelExecutionsPerNode
  edition                          = IntegrationRuntimeSsisProperties.edition
  license_type                     = IntegrationRuntimeSsisProperties.licenseType
}