mapping "azurerm_frontdoor_firewall_policy" {
  import_path = "azure-rest-api-specs/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-04-01/webapplicationfirewall.json"

  name                              = PolicyNameParameter
  resource_group_name               = ResourceGroupNameParameter
  mode                              = PolicySettings.mode
  redirect_url                      = PolicySettings.redirectUrl
  custom_block_response_status_code = PolicySettings.customBlockResponseStatusCode
  custom_block_response_body        = PolicySettings.customBlockResponseBody
}
