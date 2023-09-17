mapping "azurerm_windows_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/computeRPCommon.json"

  admin_password                                    = any //VirtualMachineScaleSetOSProfile.adminPassword
  admin_username                                    = any //VirtualMachineScaleSetOSProfile.adminUsername
  sku                                               = any //HardwareProfile.vmSize
  computer_name_prefix                              = any //VirtualMachineScaleSetOSProfile.computerNamePrefix
  custom_data                                       = any //VirtualMachineScaleSetOSProfile.customData
  do_not_run_extensions_on_overprovisioned_machines = any //VirtualMachineScaleSetProperties.doNotRunExtensionsOnOverprovisionedVMs
  enable_automatic_updates                          = WindowsConfiguration.enableAutomaticUpdates
  eviction_policy                                   = evictionPolicy
  license_type                                      = any //VirtualMachineProperties.licenseType
  max_bid_price                                     = BillingProfile.maxPrice
  overprovision                                     = any //VirtualMachineScaleSetProperties.overprovision
  priority                                          = priority
  provision_vm_agent                                = WindowsConfiguration.provisionVMAgent
  scale_in_policy                                   = any //ScaleInPolicy.rules
  single_placement_group                            = any //VirtualMachineScaleSetProperties.singlePlacementGroup
  upgrade_mode                                      = any //UpgradePolicy.mode
  zone_balance                                      = any //VirtualMachineScaleSetProperties.zoneBalance
}
