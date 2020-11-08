mapping "azurerm_hdinsight_kafka_cluster" {
  import_path = "azure-rest-api-specs/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2018-06-01-preview/cluster.json"

  name                = ClusterNameParameter
  resource_group_name = ResourceGroupNameParameter
  cluster_version     = ClusterCreateProperties.clusterVersion
  tier                = ClusterCreateProperties.tier
  min_tls_version     = ClusterCreateProperties.minSupportedTlsVersion
}
