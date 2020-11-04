mapping "azurerm_machine_learning_workspace" {
  import_path = "azure-rest-api-specs/specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2019-10-01/workspaces.json"

  name                = WorkspaceNameParameter
  resource_group_name = ResourceGroupParameter
  key_vault_id        = WorkspaceProperties.keyVaultIdentifierId
  storage_account_id  = WorkspaceProperties.userStorageAccountId
  sku_name            = Sku.name
}
