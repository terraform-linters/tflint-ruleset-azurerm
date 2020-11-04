mapping "azurerm_virtual_machine_data_disk_attachment" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/compute.json"

  lun                       = VirtualMachineScaleSetDataDisk.lun
  caching                   = Caching
  create_option             = CreateOption
  write_accelerator_enabled = VirtualMachineScaleSetDataDisk.writeAcceleratorEnabled
}