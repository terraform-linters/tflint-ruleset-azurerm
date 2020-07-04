mapping "azurerm_batch_account" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2019-08-01/BatchManagement.json"

  name                 = AccountNameParameter
  resource_group_name  = ResourceGroupNameParameter
  pool_allocation_mode = PoolAllocationMode
}

mapping "azurerm_batch_application" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2019-08-01/BatchManagement.json"

  name                = ApplicationNameParameter
  resource_group_name = ResourceGroupNameParameter
  account_name        = AccountNameParameter
  allow_updates       = ApplicationProperties.allowUpdates
  default_version     = ApplicationProperties.defaultVersion
  display_name        = ApplicationProperties.displayName
}

mapping "azurerm_batch_certificate" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2019-08-01/BatchManagement.json"

  account_name        = AccountNameParameter
  resource_group_name = ResourceGroupNameParameter
  certificate         = CertificateCreateOrUpdateProperties.data
  format              = CertificateBaseProperties.format
  password            = CertificateCreateOrUpdateProperties.password
  thumbprint          = CertificateBaseProperties.thumbprintAlgorithm
}

mapping "azurerm_batch_pool" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2019-08-01/BatchManagement.json"

  name                = ApplicationNameParameter
  resource_group_name = ResourceGroupNameParameter
  account_name        = AccountNameParameter
  node_agent_sku_id   = VirtualMachineConfiguration.nodeAgentSkuId
  vm_size             = PoolProperties.vmSize
  display_name        = PoolProperties.displayName
  max_tasks_per_node  = PoolProperties.maxTasksPerNode
}
