mapping "azurerm_lb_outbound_rule" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/loadBalancer.json"

  protocol                 = OutboundRulePropertiesFormat.protocol
  enable_tcp_reset         = OutboundRulePropertiesFormat.enableTcpReset
  allocated_outbound_ports = OutboundRulePropertiesFormat.allocatedOutboundPorts
  idle_timeout_in_minutes  = OutboundRulePropertiesFormat.idleTimeoutInMinutes
}
