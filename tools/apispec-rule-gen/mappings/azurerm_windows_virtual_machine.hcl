mapping "azurerm_windows_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/computeRPCommon.json"

  admin_password             = any //OSProfile.adminPassword
  admin_username             = any //OSProfile.adminUsername
  size                       = any //HardwareProfile.vmSize
  name                       = any
  allow_extension_operations = OSProfile.allowExtensionOperations
  computer_name              = OSProfile.computerName
  custom_data                = OSProfile.customData
  enable_automatic_updates   = WindowsConfiguration.enableAutomaticUpdates
  eviction_policy            = evictionPolicy
  license_type               = any //VirtualMachineProperties.licenseType
  max_bid_price              = BillingProfile.maxPrice
  priority                   = priority
  provision_vm_agent         = WindowsConfiguration.provisionVMAgent
}
