mapping "azurerm_dashboard" {
  import_path = "azure-rest-api-specs/specification/portal/resource-manager/Microsoft.Portal/preview/2019-01-01-preview/portal.json"

  name                 = DashboardNameParameter
  resource_group_name  = ResourceGroupNameParameter
  location             = Dashboard.location
}
