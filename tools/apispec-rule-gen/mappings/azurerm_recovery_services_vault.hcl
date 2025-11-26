mapping "azurerm_recovery_services_vault" {
  import_path = "azure-rest-api-specs/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/RecoveryServices/stable/2023-02-01/vaults.json"

  name                = VaultName
  resource_group_name = ResourceGroupName
  sku                 = Sku.name
}
