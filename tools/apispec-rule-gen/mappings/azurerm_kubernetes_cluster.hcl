mapping "azurerm_kubernetes_cluster" {
  import_path = "azure-rest-api-specs/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-07-01/managedClusters.json"

  name                       = ResourceNameParameter
  resource_group_name        = any //ResourceGroupNameParameter
  dns_prefix                 = ManagedClusterProperties.dnsPrefix
  kubernetes_version         = ManagedClusterProperties.kubernetesVersion
  node_resource_group        = ManagedClusterProperties.nodeResourceGroup

  default_node_pool = {
    vm_size = any //ContainerServiceVMSize
  }
}
