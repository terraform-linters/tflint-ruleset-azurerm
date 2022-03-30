mapping "azurerm_lb_nat_pool" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/loadBalancer.json"

  protocol            = TransportProtocol
  frontend_port_start = InboundNatPoolPropertiesFormat.frontendPortRangeStart
  frontend_port_end   = InboundNatPoolPropertiesFormat.frontendPortRangeEnd
  backend_port        = InboundNatPoolPropertiesFormat.backendPort
}
