mapping "azurerm_signalr_service" {
  import_path = "azure-rest-api-specs/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2018-10-01/signalr.json"

  name                = SignalRServiceName
  resource_group_name = ResourceGroupParameter
}
