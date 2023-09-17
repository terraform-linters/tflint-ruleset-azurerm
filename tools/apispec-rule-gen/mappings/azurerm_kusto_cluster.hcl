mapping "azurerm_kusto_cluster" {
  import_path = "azure-rest-api-specs/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/kusto.json"

  name                    = ClusterNameParameter
  resource_group_name     = ResourceGroupParameter
  enable_disk_encryption  = ClusterProperties.enableDiskEncryption
  enable_streaming_ingest = ClusterProperties.enableStreamingIngest
}
