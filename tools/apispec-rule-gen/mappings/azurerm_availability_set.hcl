mapping "azurerm_availability_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/availabilitySet.json"

  platform_update_domain_count = AvailabilitySetProperties.platformUpdateDomainCount
  platform_fault_domain_count  = AvailabilitySetProperties.platformFaultDomainCount
}
