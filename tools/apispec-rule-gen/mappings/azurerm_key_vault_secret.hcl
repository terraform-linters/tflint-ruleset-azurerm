mapping "azurerm_key_vault_secret" {
  import_path = "azure-rest-api-specs/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/secrets.json"

  value           = SecretProperties.value
  content_type    = SecretProperties.contentType
  not_before_date = Attributes.nbf
  expiration_date = Attributes.exp
}
