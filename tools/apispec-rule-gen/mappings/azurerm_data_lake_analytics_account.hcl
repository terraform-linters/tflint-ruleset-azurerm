mapping "azurerm_data_lake_analytics_account" {
  import_path = "azure-rest-api-specs/specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/stable/2016-11-01/account.json"

  name                       = AccountNameParameter
  resource_group_name        = ResourceGroupNameParameter
  default_store_account_name = CreateDataLakeAnalyticsAccountProperties.defaultDataLakeStoreAccount
  tier                       = CreateDataLakeAnalyticsAccountProperties.newTier
}