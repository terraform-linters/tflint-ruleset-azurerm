mapping "azurerm_eventgrid_domain" {
  import_path = "azure-rest-api-specs/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2020-04-01-preview/EventGrid.json"

  input_schema = DomainProperties.inputSchema
}
