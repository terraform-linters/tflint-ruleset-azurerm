mapping "azurerm_machine_learning_workspace" {
  import_path = "azure-rest-api-specs/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/workspaceRP.json"

  name                = WorkspaceNameParameter
  resource_group_name = any //ResourceGroupParameter
  key_vault_id        = WorkspaceProperties.keyVault
  storage_account_id  = WorkspaceProperties.storageAccount
  sku_name            = any //Sku.name
}
