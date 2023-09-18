mapping "azurerm_nat_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/natGateway.json"

  idle_timeout_in_minutes = NatGatewayPropertiesFormat.idleTimeoutInMinutes
  public_ip_address_ids   = NatGatewayPropertiesFormat.publicIpAddresses
  public_ip_prefix_ids    = NatGatewayPropertiesFormat.publicIpPrefixes
  sku_name                = NatGatewaySku.name
  zones                   = NatGateway.zones
}
