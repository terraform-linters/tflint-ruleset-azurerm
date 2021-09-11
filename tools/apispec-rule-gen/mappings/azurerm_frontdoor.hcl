mapping "azurerm_frontdoor" {
  import_path = "azure-rest-api-specs/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/frontdoor.json"

  name                                       = frontDoorNameParameter
  resource_group_name                        = resourceGroupNameParameter
  backend_pools_send_receive_timeout_seconds = BackendPoolsSettings.sendRecvTimeoutSeconds
  friendly_name                              = FrontDoorUpdateParameters.friendlyName
}
