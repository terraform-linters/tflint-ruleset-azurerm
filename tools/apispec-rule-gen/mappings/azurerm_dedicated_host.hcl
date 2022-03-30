mapping "azurerm_dedicated_host" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/compute.json"

  sku_name                = Sku.name
  platform_fault_domain   = DedicatedHostProperties.platformFaultDomain
  auto_replace_on_failure = DedicatedHostProperties.autoReplaceOnFailure
  license_type            = DedicatedHostLicenseType
}
