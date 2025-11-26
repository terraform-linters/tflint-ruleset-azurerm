mapping "azurerm_service_fabric_cluster" {
  import_path = "azure-rest-api-specs/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/ServiceFabric/stable/2021-06-01/cluster.json"

  name                 = clusterNameParameter
  resource_group_name  = resourceGroupNameParameter
  reliability_level    = ReliabilityLevel
  management_endpoint  = ClusterProperties.managementEndpoint
  upgrade_mode         = UpgradeMode
  vm_image             = ClusterProperties.vmImage
  cluster_code_version = ClusterProperties.clusterCodeVersion
  add_on_features      = ClusterProperties.addOnFeatures
}
