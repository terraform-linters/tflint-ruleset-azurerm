mapping "azurerm_batch_pool" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/BatchManagement.json"

  name                = ApplicationNameParameter
  resource_group_name = ResourceGroupNameParameter
  account_name        = AccountNameParameter
  node_agent_sku_id   = VirtualMachineConfiguration.nodeAgentSkuId
  vm_size             = PoolProperties.vmSize
  display_name        = PoolProperties.displayName
  max_tasks_per_node  = PoolProperties.taskSlotsPerNode
}
