mapping "azurerm_container_group" {
  import_path = "azure-rest-api-specs/specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2018-10-01/containerInstance.json"

  name                = ContainerGroupNameParameter
  resource_group_name = ResourceGroupNameParameter
  location            = LocationParameter
  dns_name_label      = IpAddress.dnsNameLabel
  ip_address_type     = IpAddress.type
  network_profile_id  = ContainerGroupNetworkProfile.id
}

mapping "azurerm_container_registry" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2017-10-01/containerregistry.json"

  name                = RegistryNameParameter
  resource_group_name = ResourceGroupParameter
  admin_enabled       = RegistryProperties.adminUserEnabled
  storage_account_id  = StorageAccountProperties.id
  sku                 = Sku.name
}

mapping "azurerm_container_registry_webhook" {
  import_path = "azure-rest-api-specs/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2017-10-01/containerregistry.json"

  name                = WebhookNameParameter
  resource_group_name = ResourceGroupParameter
  registry_name       = RegistryNameParameter
  service_uri         = WebhookPropertiesCreateParameters.serviceUri
  actions             = WebhookProperties.actions
  status              = WebhookProperties.status
  scope               = WebhookProperties.scope
}

mapping "azurerm_kubernetes_cluster" {
  import_path = "azure-rest-api-specs/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2020-04-01/managedClusters.json"

  name                       = ResourceNameParameter
  resource_group_name        = ResourceGroupNameParameter
  dns_prefix                 = ContainerServiceMasterProfile.dnsPrefix
  enable_pod_security_policy = ManagedClusterProperties.enablePodSecurityPolicy
  kubernetes_version         = ManagedClusterProperties.kubernetesVersion
  node_resource_group        = ManagedClusterProperties.nodeResourceGroup
}

mapping "azurerm_kubernetes_cluster_node_pool" {
  import_path = "azure-rest-api-specs/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2020-04-01/managedClusters.json"

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
