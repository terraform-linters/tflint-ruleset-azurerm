mapping "azurerm_eventgrid_domain" {
  import_path = "azure-rest-api-specs/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/EventGrid.json"

  input_schema = DomainProperties.inputSchema
}
