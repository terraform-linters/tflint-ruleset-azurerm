mapping "azurerm_storage_account_customer_managed_key" {
  import_path = "azure-rest-api-specs/specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/storage.json"

  key_name    = KeyVaultProperties.keyname
  key_version = KeyVaultProperties.keyversion
}
