mapping "azurerm_lb_nat_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/loadBalancer.json"

  protocol                = TransportProtocol
  frontend_port           = InboundNatRulePropertiesFormat.frontendPort
  backend_port            = InboundNatRulePropertiesFormat.backendPort
  idle_timeout_in_minutes = InboundNatRulePropertiesFormat.idleTimeoutInMinutes
  enable_floating_ip      = InboundNatRulePropertiesFormat.enableFloatingIP
  enable_tcp_reset        = InboundNatRulePropertiesFormat.enableTcpReset
}
