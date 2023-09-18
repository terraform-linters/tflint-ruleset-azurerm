mapping "azurerm_virtual_machine_data_disk_attachment" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/computeRPCommon.json"

  lun                       = any //VirtualMachineScaleSetDataDisk.lun
  caching                   = Caching
  create_option             = CreateOption
  write_accelerator_enabled = any //VirtualMachineScaleSetDataDisk.writeAcceleratorEnabled
}
