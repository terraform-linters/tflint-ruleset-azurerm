mapping "azurerm_kusto_cluster" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2020-02-15/kusto.json"

  name                    = ClusterNameParameter
  resource_group_name     = ResourceGroupParameter
  enable_disk_encryption  = ClusterProperties.enableDiskEncryption
  enable_streaming_ingest = ClusterProperties.enableStreamingIngest
}
