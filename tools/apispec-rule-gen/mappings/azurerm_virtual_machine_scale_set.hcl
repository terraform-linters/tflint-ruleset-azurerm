mapping "azurerm_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/compute.json"

  upgrade_policy_mode    = UpgradePolicy.mode
  eviction_policy        = evictionPolicy
  license_type           = VirtualMachineProperties.licenseType
  overprovision          = VirtualMachineScaleSetProperties.overprovision
  priority               = priority
  single_placement_group = VirtualMachineScaleSetProperties.singlePlacementGroup
}