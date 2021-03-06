mapping "azurerm_dedicated_host_group" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/compute.json"

  platform_fault_domain_count = DedicatedHostGroupProperties.platformFaultDomainCount
  zones                       = DedicatedHostGroup.zones
}
