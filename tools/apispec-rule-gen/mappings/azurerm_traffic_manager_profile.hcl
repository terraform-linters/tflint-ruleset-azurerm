mapping "azurerm_traffic_manager_profile" {
  import_path = "azure-rest-api-specs/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/trafficmanager.json"

  profile_status         = ProfileProperties.profileStatus
  traffic_routing_method = ProfileProperties.trafficRoutingMethod
}
