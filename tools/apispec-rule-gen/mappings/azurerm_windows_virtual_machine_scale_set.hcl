mapping "azurerm_windows_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/compute.json"

  admin_password                                    = VirtualMachineScaleSetOSProfile.adminPassword
  admin_username                                    = VirtualMachineScaleSetOSProfile.adminUsername
  sku                                               = any //HardwareProfile.vmSize
  computer_name_prefix                              = VirtualMachineScaleSetOSProfile.computerNamePrefix
  custom_data                                       = VirtualMachineScaleSetOSProfile.customData
  do_not_run_extensions_on_overprovisioned_machines = VirtualMachineScaleSetProperties.doNotRunExtensionsOnOverprovisionedVMs
  enable_automatic_updates                          = WindowsConfiguration.enableAutomaticUpdates
  eviction_policy                                   = evictionPolicy
  license_type                                      = VirtualMachineProperties.licenseType
  max_bid_price                                     = BillingProfile.maxPrice
  overprovision                                     = VirtualMachineScaleSetProperties.overprovision
  priority                                          = priority
  provision_vm_agent                                = WindowsConfiguration.provisionVMAgent
  scale_in_policy                                   = ScaleInPolicy.rules
  single_placement_group                            = VirtualMachineScaleSetProperties.singlePlacementGroup
  upgrade_mode                                      = UpgradePolicy.mode
  zone_balance                                      = VirtualMachineScaleSetProperties.zoneBalance
}
