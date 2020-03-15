mapping "azurerm_lb" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/loadBalancer.json"

  sku = LoadBalancerSku.name
}

mapping "azurerm_lb_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/loadBalancer.json"

  protocol                = TransportProtocol
  frontend_port           = LoadBalancingRulePropertiesFormat.frontendPort
  backend_port            = LoadBalancingRulePropertiesFormat.backendPort
  enable_floating_ip      = LoadBalancingRulePropertiesFormat.enableFloatingIP
  idle_timeout_in_minutes = LoadBalancingRulePropertiesFormat.idleTimeoutInMinutes
  load_distribution       = LoadBalancingRulePropertiesFormat.loadDistribution
  disable_outbound_snat   = LoadBalancingRulePropertiesFormat.disableOutboundSnat
  enable_tcp_reset        = LoadBalancingRulePropertiesFormat.enableTcpReset
}

mapping "azurerm_lb_outbound_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/loadBalancer.json"

  protocol                 = OutboundRulePropertiesFormat.protocol
  enable_tcp_reset         = OutboundRulePropertiesFormat.enableTcpReset
  allocated_outbound_ports = OutboundRulePropertiesFormat.allocatedOutboundPorts
  idle_timeout_in_minutes  = OutboundRulePropertiesFormat.idleTimeoutInMinutes
}

mapping "azurerm_lb_nat_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/loadBalancer.json"

  protocol                = TransportProtocol
  frontend_port           = InboundNatRulePropertiesFormat.frontendPort
  backend_port            = InboundNatRulePropertiesFormat.backendPort
  idle_timeout_in_minutes = InboundNatRulePropertiesFormat.idleTimeoutInMinutes
  enable_floating_ip      = InboundNatRulePropertiesFormat.enableFloatingIP
  enable_tcp_reset        = InboundNatRulePropertiesFormat.enableTcpReset
}

mapping "azurerm_lb_nat_pool" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/loadBalancer.json"

  protocol            = TransportProtocol
  frontend_port_start = InboundNatPoolPropertiesFormat.frontendPortRangeStart
  frontend_port_end   = InboundNatPoolPropertiesFormat.frontendPortRangeEnd
  backend_port        = InboundNatPoolPropertiesFormat.backendPort
}

mapping "azurerm_lb_probe" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/loadBalancer.json"

  protocol            = ProbePropertiesFormat.protocol
  port                = ProbePropertiesFormat.port
  request_path        = ProbePropertiesFormat.requestPath
  interval_in_seconds = ProbePropertiesFormat.intervalInSeconds
  number_of_probes    = ProbePropertiesFormat.numberOfProbes
}
