mapping "azurerm_lb_probe" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/loadBalancer.json"

  protocol            = ProbePropertiesFormat.protocol
  port                = ProbePropertiesFormat.port
  request_path        = ProbePropertiesFormat.requestPath
  interval_in_seconds = ProbePropertiesFormat.intervalInSeconds
  number_of_probes    = ProbePropertiesFormat.numberOfProbes
}
