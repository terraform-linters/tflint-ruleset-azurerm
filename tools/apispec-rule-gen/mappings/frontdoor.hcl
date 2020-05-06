mapping "azurerm_frontdoor" {
  import_path = "azure-rest-api-specs/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-01-01/frontdoor.json"

  name                                       = frontDoorNameParameter
  resource_group_name                        = resourceGroupNameParameter
  backend_pools_send_receive_timeout_seconds = BackendPoolsSettings.sendRecvTimeoutSeconds
  friendly_name                              = FrontDoorUpdateParameters.friendlyName
}

mapping "azurerm_frontdoor_firewall_policy" {
  import_path = "azure-rest-api-specs/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-10-01/webapplicationfirewall.json"

  name                              = PolicyNameParameter
  resource_group_name               = ResourceGroupNameParameter
  mode                              = PolicySettings.mode
  redirect_url                      = PolicySettings.redirectUrl
  custom_block_response_status_code = PolicySettings.customBlockResponseStatusCode
  custom_block_response_body        = PolicySettings.customBlockResponseBody
}
