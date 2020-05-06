mapping "azurerm_kusto_cluster" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2020-02-15/kusto.json"

  name                    = ClusterNameParameter
  resource_group_name     = ResourceGroupParameter
  enable_disk_encryption  = ClusterProperties.enableDiskEncryption
  enable_streaming_ingest = ClusterProperties.enableStreamingIngest
}

mapping "azurerm_kusto_database" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2020-02-15/kusto.json"

  name                = DatabaseNameParameter
  resource_group_name = ResourceGroupParameter
  cluster_name        = ClusterNameParameter
  hot_cache_period    = ReadWriteDatabaseProperties.hotCachePeriod
  soft_delete_period  = ReadWriteDatabaseProperties.softDeletePeriod
}

mapping "azurerm_kusto_database_principal" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2020-02-15/kusto.json"

  resource_group_name = ResourceGroupParameter
  cluster_name        = ClusterNameParameter
  database_name       = DatabaseNameParameter
  role                = DatabasePrincipalProperties.role
  type                = DatabasePrincipalProperties.principalType
}

mapping "azurerm_kusto_eventhub_data_connection" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2020-02-15/kusto.json"

  name                = DataConnectionNameParameter
  resource_group_name = ResourceGroupParameter
  cluster_name        = ResourceGroupParameter
  database_name       = DatabaseNameParameter
  eventhub_id         = EventHubConnectionProperties.eventHubResourceId
  consumer_group      = EventHubConnectionProperties.consumerGroup
  table_name          = EventHubConnectionProperties.tableName
  mapping_rule_name   = EventHubConnectionProperties.mappingRuleName
  data_format         = EventHubDataFormat
}
