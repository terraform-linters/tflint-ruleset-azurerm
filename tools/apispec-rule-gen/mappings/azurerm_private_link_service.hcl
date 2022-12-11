mapping "azurerm_private_link_service" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/privateLinkService.json"

  load_balancer_frontend_ip_configuration_ids = PrivateLinkServiceProperties.loadBalancerFrontendIpConfigurations
  enable_proxy_protocol                       = PrivateLinkServiceProperties.enableProxyProtocol
}
