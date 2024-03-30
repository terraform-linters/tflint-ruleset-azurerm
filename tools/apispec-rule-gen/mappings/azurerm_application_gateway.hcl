mapping "azurerm_application_gateway" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/applicationGateway.json"

  zones        = ApplicationGateway.zones
  enable_http2 = ApplicationGatewayPropertiesFormat.enableHttp2
}
