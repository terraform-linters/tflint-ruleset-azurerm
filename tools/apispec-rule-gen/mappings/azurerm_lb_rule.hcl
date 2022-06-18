mapping "azurerm_lb_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/loadBalancer.json"

  protocol                = TransportProtocol
  frontend_port           = LoadBalancingRulePropertiesFormat.frontendPort
  backend_port            = LoadBalancingRulePropertiesFormat.backendPort
  enable_floating_ip      = LoadBalancingRulePropertiesFormat.enableFloatingIP
  idle_timeout_in_minutes = LoadBalancingRulePropertiesFormat.idleTimeoutInMinutes
  load_distribution       = LoadBalancingRulePropertiesFormat.loadDistribution
  disable_outbound_snat   = LoadBalancingRulePropertiesFormat.disableOutboundSnat
  enable_tcp_reset        = LoadBalancingRulePropertiesFormat.enableTcpReset
}
