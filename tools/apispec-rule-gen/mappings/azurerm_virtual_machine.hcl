mapping "azurerm_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/compute.json"

  vm_size      = any //HardwareProfile.vmSize
  license_type = VirtualMachineProperties.licenseType
}