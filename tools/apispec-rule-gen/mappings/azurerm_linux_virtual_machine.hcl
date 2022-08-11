mapping "azurerm_linux_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2021-11-01/compute.json"

  admin_username                  = OSProfile.adminUsername
  size                            = any //HardwareProfile.vmSize
  admin_password                  = OSProfile.adminPassword
  allow_extension_operations      = OSProfile.allowExtensionOperations
  computer_name                   = OSProfile.computerName
  custom_data                     = OSProfile.customData
  disable_password_authentication = LinuxConfiguration.disablePasswordAuthentication
  eviction_policy                 = evictionPolicy
  max_bid_price                   = BillingProfile.maxPrice
  priority                        = priority
  provision_vm_agent              = LinuxConfiguration.provisionVMAgent
}
