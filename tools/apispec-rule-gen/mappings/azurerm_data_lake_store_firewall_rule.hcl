mapping "azurerm_data_lake_store_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/account.json"

  resource_group_name = ResourceGroupNameParameter
  account_name        = AccountNameParameter
  start_ip_address    = CreateOrUpdateFirewallRuleProperties.startIpAddress
  end_ip_address      = CreateOrUpdateFirewallRuleProperties.endIpAddress
}
