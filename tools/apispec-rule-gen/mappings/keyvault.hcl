mapping "azurerm_key_vault" {
  import_path = "azure-rest-api-specs/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2018-02-14/keyvault.json"

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

mapping "azurerm_key_vault_access_policy" {
  import_path = "azure-rest-api-specs/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2018-02-14/keyvault.json"

  tenant_id      = AccessPolicyEntry.tenantId
  object_id      = AccessPolicyEntry.objectId
  application_id = AccessPolicyEntry.applicationId
}

mapping "azurerm_key_vault_key" {
  import_path = "azure-rest-api-specs/specification/keyvault/data-plane/Microsoft.KeyVault/stable/2016-10-01/keyvault.json"

  key_type        = KeyCreateParameters.kty
  key_size        = KeyCreateParameters.key_size
  curve           = KeyCreateParameters.crv
  key_opts        = KeyCreateParameters.key_ops
  not_before_date = Attributes.nbf
  expiration_date = Attributes.exp
}

mapping "azurerm_key_vault_secret" {
  import_path = "azure-rest-api-specs/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2018-02-14/secrets.json"

  value           = SecretProperties.value
  content_type    = SecretProperties.contentType
  not_before_date = Attributes.nbf
  expiration_date = Attributes.exp
}
