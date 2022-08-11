mapping "azurerm_linux_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2021-11-01/compute.json"

  admin_username                                    = VirtualMachineScaleSetOSProfile.adminUsername
  sku                                               = any //HardwareProfile.vmSize
  admin_password                                    = VirtualMachineScaleSetOSProfile.adminPassword
  computer_name_prefix                              = VirtualMachineScaleSetOSProfile.computerNamePrefix
  custom_data                                       = VirtualMachineScaleSetOSProfile.customData
  disable_password_authentication                   = LinuxConfiguration.disablePasswordAuthentication
  do_not_run_extensions_on_overprovisioned_machines = VirtualMachineScaleSetProperties.doNotRunExtensionsOnOverprovisionedVMs
  eviction_policy                                   = evictionPolicy
  max_bid_price                                     = BillingProfile.maxPrice
  overprovision                                     = VirtualMachineScaleSetProperties.overprovision
  priority                                          = priority
  provision_vm_agent                                = LinuxConfiguration.provisionVMAgent
  scale_in_policy                                   = ScaleInPolicy.rules
  single_placement_group                            = VirtualMachineScaleSetProperties.singlePlacementGroup
  upgrade_mode                                      = UpgradePolicy.mode
  zone_balance                                      = VirtualMachineScaleSetProperties.zoneBalance
}
