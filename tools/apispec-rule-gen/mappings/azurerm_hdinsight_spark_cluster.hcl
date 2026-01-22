mapping "azurerm_hdinsight_spark_cluster" {
  import_path = "azure-rest-api-specs/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsight/stable/2021-06-01/cluster.json"

  name                = ClusterNameParameter
  resource_group_name = ResourceGroupNameParameter
  cluster_version     = ClusterCreateProperties.clusterVersion
  tier                = ClusterCreateProperties.tier
  min_tls_version     = ClusterCreateProperties.minSupportedTlsVersion
}
