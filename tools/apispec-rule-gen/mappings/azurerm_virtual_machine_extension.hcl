mapping "azurerm_virtual_machine_extension" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/virtualMachine.json"

  publisher                  = VirtualMachineExtensionProperties.publisher
  type                       = VirtualMachineExtensionProperties.type
  type_handler_version       = VirtualMachineExtensionProperties.typeHandlerVersion
  auto_upgrade_minor_version = VirtualMachineExtensionProperties.autoUpgradeMinorVersion
  settings                   = VirtualMachineExtensionProperties.settings
  protected_settings         = VirtualMachineExtensionProperties.protectedSettings
}
