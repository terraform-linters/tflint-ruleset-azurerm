mapping "azurerm_dedicated_host_group" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/dedicatedHost.json"

  platform_fault_domain_count = DedicatedHostGroupProperties.platformFaultDomainCount
  zones                       = DedicatedHostGroup.zones
}
