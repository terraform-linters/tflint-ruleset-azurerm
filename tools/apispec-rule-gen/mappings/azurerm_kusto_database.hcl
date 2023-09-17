mapping "azurerm_kusto_database" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/kusto.json"

  name                = DatabaseNameParameter
  resource_group_name = any //ResourceGroupParameter
  cluster_name        = ClusterNameParameter
  hot_cache_period    = ReadWriteDatabaseProperties.hotCachePeriod
  soft_delete_period  = ReadWriteDatabaseProperties.softDeletePeriod
}
