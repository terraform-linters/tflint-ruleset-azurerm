mapping "azurerm_batch_account" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2020-03-01/BatchManagement.json"

  name                 = AccountNameParameter
  resource_group_name  = ResourceGroupNameParameter
  pool_allocation_mode = PoolAllocationMode
}
