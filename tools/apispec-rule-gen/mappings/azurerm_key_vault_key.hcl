mapping "azurerm_key_vault_key" {
  import_path = "azure-rest-api-specs/specification/keyvault/data-plane/Microsoft.KeyVault/stable/2016-10-01/keyvault.json"

  key_type        = KeyCreateParameters.kty
  key_size        = KeyCreateParameters.key_size
  curve           = KeyCreateParameters.crv
  key_opts        = KeyCreateParameters.key_ops
  not_before_date = Attributes.nbf
  expiration_date = Attributes.exp
}
