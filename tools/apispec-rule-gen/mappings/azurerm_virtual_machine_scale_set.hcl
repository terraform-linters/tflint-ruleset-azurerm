mapping "azurerm_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/ComputeRP.json"

  upgrade_policy_mode    = any //UpgradePolicy.mode
  eviction_policy        = VirtualMachineEvictionPolicyTypes
  license_type           = any //VirtualMachineProperties.licenseType
  overprovision          = any //VirtualMachineScaleSetProperties.overprovision
  priority               = VirtualMachinePriorityTypes
  single_placement_group = any //VirtualMachineScaleSetProperties.singlePlacementGroup
}
