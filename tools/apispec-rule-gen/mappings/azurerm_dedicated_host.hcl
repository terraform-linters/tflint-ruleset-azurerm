mapping "azurerm_dedicated_host" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/dedicatedHost.json"

  sku_name                = any //Sku.name
  platform_fault_domain   = DedicatedHostProperties.platformFaultDomain
  auto_replace_on_failure = DedicatedHostProperties.autoReplaceOnFailure
  license_type            = DedicatedHostLicenseType
}
