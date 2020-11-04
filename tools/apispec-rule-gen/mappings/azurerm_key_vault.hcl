mapping "azurerm_key_vault" {
  import_path = "azure-rest-api-specs/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2019-09-01/keyvault.json"

  name                            = VaultName
  resource_group_name             = ResourceGroupName
  sku_name                        = Sku.name
  tenant_id                       = VaultProperties.tenantId
  enabled_for_deployment          = VaultProperties.enabledForDeployment
  enabled_for_disk_encryption     = VaultProperties.enabledForDiskEncryption
  enabled_for_template_deployment = VaultProperties.enabledForTemplateDeployment
  purge_protection_enabled        = VaultProperties.enablePurgeProtection
  soft_delete_enabled             = VaultProperties.enableSoftDelete
}
