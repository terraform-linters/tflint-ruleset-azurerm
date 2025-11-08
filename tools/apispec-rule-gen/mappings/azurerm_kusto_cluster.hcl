mapping "azurerm_kusto_cluster" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/Kusto/stable/2024-04-13/kusto.json"

  name                    = ClusterNameParameter
  resource_group_name     = any //ResourceGroupParameter
  enable_disk_encryption  = ClusterProperties.enableDiskEncryption
  enable_streaming_ingest = ClusterProperties.enableStreamingIngest
}
