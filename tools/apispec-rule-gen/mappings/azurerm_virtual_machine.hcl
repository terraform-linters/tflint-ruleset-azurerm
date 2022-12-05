mapping "azurerm_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/virtualMachine.json"

  vm_size      = any //HardwareProfile.vmSize
  license_type = VirtualMachineProperties.licenseType
}
