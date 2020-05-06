mapping "azurerm_data_lake_analytics_account" {
  import_path = "azure-rest-api-specs/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/account.json"

  name                       = AccountNameParameter
  resource_group_name        = ResourceGroupNameParameter
  default_store_account_name = CreateDataLakeAnalyticsAccountProperties.defaultDataLakeStoreAccount
  tier                       = CreateDataLakeAnalyticsAccountProperties.newTier
}

mapping "azurerm_data_lake_analytics_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/account.json"

  resource_group_name = ResourceGroupNameParameter
  account_name        = AccountNameParameter
  start_ip_address    = CreateOrUpdateFirewallRuleProperties.startIpAddress
  end_ip_address      = CreateOrUpdateFirewallRuleProperties.endIpAddress
}

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

mapping "azurerm_data_lake_store_file" {
  import_path = "azure-rest-api-specs/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/account.json"

  account_name = AccountNameParameter
}

mapping "azurerm_data_lake_store_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/account.json"

  resource_group_name = ResourceGroupNameParameter
  account_name        = AccountNameParameter
  start_ip_address    = CreateOrUpdateFirewallRuleProperties.startIpAddress
  end_ip_address      = CreateOrUpdateFirewallRuleProperties.endIpAddress
}
