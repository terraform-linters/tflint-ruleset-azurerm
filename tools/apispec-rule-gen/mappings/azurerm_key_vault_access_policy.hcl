mapping "azurerm_key_vault_access_policy" {
  import_path = "azure-rest-api-specs/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2019-09-01/keyvault.json"

  tenant_id      = AccessPolicyEntry.tenantId
  object_id      = AccessPolicyEntry.objectId
  application_id = AccessPolicyEntry.applicationId
}
