mapping "azurerm_data_lake_store" {
  import_path = "azure-rest-api-specs/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/account.json"

  name                     = AccountNameParameter
  resource_group_name      = ResourceGroupNameParameter
  tier                     = CreateDataLakeStoreAccountProperties.newTier
  encryption_state         = CreateDataLakeStoreAccountProperties.encryptionState
  encryption_type          = EncryptionConfig.type
  firewall_allow_azure_ips = CreateDataLakeStoreAccountProperties.firewallAllowAzureIps
  firewall_state           = CreateDataLakeStoreAccountProperties.firewallState
}