mapping "azurerm_dev_test_lab" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  artifacts_storage_account_id         = LabProperties.artifactsStorageAccount
  default_storage_account_id           = LabProperties.defaultStorageAccount
  default_premium_storage_account_id   = LabProperties.defaultPremiumStorageAccount
  key_vault_id                         = UserSecretStore.keyVaultId
  premium_data_disk_storage_account_id = LabProperties.premiumDataDiskStorageAccount
  unique_identifier                    = LabProperties.uniqueIdentifier
}