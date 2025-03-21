mapping "azurerm_signalr_service" {
  import_path = "azure-rest-api-specs/specification/signalr/resource-manager/Microsoft.SignalRService/stable/2024-03-01/signalr.json"

  name                = ResourceName
  resource_group_name = any //ResourceGroupParameter
}
