mapping "azurerm_traffic_manager_endpoint" {
  import_path = "azure-rest-api-specs/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/trafficmanager.json"

  endpoint_status     = EndpointProperties.endpointStatus
  target              = EndpointProperties.target
  target_resource_id  = EndpointProperties.targetResourceId
  weight              = EndpointProperties.weight
  priority            = EndpointProperties.priority
  endpoint_location   = EndpointProperties.endpointLocation
  min_child_endpoints = EndpointProperties.minChildEndpoints
  geo_mappings        = EndpointProperties.geoMapping
}
