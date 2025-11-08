mapping "azurerm_kusto_eventhub_data_connection" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/Kusto/stable/2024-04-13/kusto.json"

  name                = DataConnectionNameParameter
  resource_group_name = any //ResourceGroupParameter
  cluster_name        = any //ResourceGroupParameter
  database_name       = DatabaseNameParameter
  eventhub_id         = EventHubConnectionProperties.eventHubResourceId
  consumer_group      = EventHubConnectionProperties.consumerGroup
  table_name          = EventHubConnectionProperties.tableName
  mapping_rule_name   = EventHubConnectionProperties.mappingRuleName
  data_format         = EventHubDataFormat
}
