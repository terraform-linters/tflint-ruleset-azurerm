mapping "azurerm_kubernetes_cluster_node_pool" {
  import_path = "azure-rest-api-specs/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-05-01/managedClusters.json"

  vm_size               = ContainerServiceVMSize
  availability_zones    = ManagedClusterAgentPoolProfileProperties.availabilityZones
  enable_auto_scaling   = ManagedClusterAgentPoolProfileProperties.enableAutoScaling
  enable_node_public_ip = ManagedClusterAgentPoolProfileProperties.enableNodePublicIP
  max_pods              = ManagedClusterAgentPoolProfileProperties.maxPods
  node_labels           = ManagedClusterAgentPoolProfileProperties.nodeLabels
  node_taints           = ManagedClusterAgentPoolProfileProperties.nodeTaints
  os_disk_size_gb       = ContainerServiceOSDisk
  os_type               = OSType
  vnet_subnet_id        = ContainerServiceVnetSubnetID
}
