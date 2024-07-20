mapping "azurerm_virtual_machine_scale_set_extension" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/virtualMachineScaleSet.json"

  publisher                  = VirtualMachineScaleSetExtensionProperties.publisher
  type                       = VirtualMachineScaleSetExtensionProperties.type
  type_handler_version       = VirtualMachineScaleSetExtensionProperties.typeHandlerVersion
  auto_upgrade_minor_version = VirtualMachineScaleSetExtensionProperties.autoUpgradeMinorVersion
  force_update_tag           = VirtualMachineScaleSetExtensionProperties.forceUpdateTag
  protected_settings         = VirtualMachineScaleSetExtensionProperties.protectedSettings
  provision_after_extensions = VirtualMachineScaleSetExtensionProperties.provisionAfterExtensions
  settings                   = VirtualMachineScaleSetExtensionProperties.settings
}
