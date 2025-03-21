mapping "azurerm_key_vault_secret" {
  import_path = "azure-rest-api-specs/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.4/secrets.json"

  value           = SecretBundle.value
  content_type    = SecretBundle.contentType
  not_before_date = any //Attributes.nbf
  expiration_date = any //Attributes.exp
}
