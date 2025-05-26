mapping "azurerm_linux_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/computeRPCommon.json"

  admin_username                                    = any //VirtualMachineScaleSetOSProfile.adminUsername
  sku                                               = any //HardwareProfile.vmSize
  admin_password                                    = any //VirtualMachineScaleSetOSProfile.adminPassword
  computer_name_prefix                              = any //VirtualMachineScaleSetOSProfile.computerNamePrefix
  custom_data                                       = any //VirtualMachineScaleSetOSProfile.customData
  disable_password_authentication                   = LinuxConfiguration.disablePasswordAuthentication
  do_not_run_extensions_on_overprovisioned_machines = any //VirtualMachineScaleSetProperties.doNotRunExtensionsOnOverprovisionedVMs
  eviction_policy                                   = evictionPolicy
  max_bid_price                                     = BillingProfile.maxPrice
  overprovision                                     = any //VirtualMachineScaleSetProperties.overprovision
  priority                                          = priority
  provision_vm_agent                                = LinuxConfiguration.provisionVMAgent
  scale_in_policy                                   = any //ScaleInPolicy.rules
  single_placement_group                            = any //VirtualMachineScaleSetProperties.singlePlacementGroup
  upgrade_mode                                      = any //UpgradePolicy.mode
  zone_balance                                      = any //VirtualMachineScaleSetProperties.zoneBalance
}
