mapping "azurerm_kusto_database_principal" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-01-01/kusto.json"

  resource_group_name = ResourceGroupParameter
  cluster_name        = ClusterNameParameter
  database_name       = DatabaseNameParameter
  role                = DatabasePrincipalProperties.role
  type                = DatabasePrincipalProperties.principalType
}
