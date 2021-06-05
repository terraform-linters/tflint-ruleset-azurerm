mapping "azurerm_kubernetes_cluster" {
  import_path = "azure-rest-api-specs/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-03-01/managedClusters.json"

  name                       = ResourceNameParameter
  resource_group_name        = ResourceGroupNameParameter
  dns_prefix                 = ContainerServiceMasterProfile.dnsPrefix
  enable_pod_security_policy = ManagedClusterProperties.enablePodSecurityPolicy
  kubernetes_version         = ManagedClusterProperties.kubernetesVersion
  node_resource_group        = ManagedClusterProperties.nodeResourceGroup
}
