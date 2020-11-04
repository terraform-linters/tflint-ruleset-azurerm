mapping "azurerm_virtual_hub_connection" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-05-01/virtualWan.json"

  hub_to_vitual_network_traffic_allowed          = HubVirtualNetworkConnectionProperties.allowHubToRemoteVnetTransit
  vitual_network_to_hub_gateways_traffic_allowed = HubVirtualNetworkConnectionProperties.allowRemoteVnetToUseHubVnetGateways
  internet_security_enabled                      = HubVirtualNetworkConnectionProperties.enableInternetSecurity
}
