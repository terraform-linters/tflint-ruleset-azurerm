mapping "azurerm_availability_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/compute.json"

  platform_update_domain_count = AvailabilitySetProperties.platformUpdateDomainCount
  platform_fault_domain_count  = AvailabilitySetProperties.platformFaultDomainCount
}
