mapping "azurerm_dev_test_linux_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/DevTestLabs/stable/2018-09-15/DTL.json"

  resource_group_name        = any //resourceGroupName
  location                   = any //locationName
  lab_subnet_name            = LabVirtualMachineCreationParameterProperties.labSubnetName
  lab_virtual_network_id     = LabVirtualMachineCreationParameterProperties.labVirtualNetworkId
  size                       = LabVirtualMachineCreationParameterProperties.size
  storage_type               = LabVirtualMachineCreationParameterProperties.storageType
  username                   = LabVirtualMachineCreationParameterProperties.userName
  allow_claim                = LabVirtualMachineCreationParameterProperties.allowClaim
  disallow_public_ip_address = LabVirtualMachineCreationParameterProperties.disallowPublicIpAddress
  notes                      = LabVirtualMachineCreationParameterProperties.notes
  password                   = LabVirtualMachineCreationParameterProperties.password
  ssh_key                    = LabVirtualMachineCreationParameterProperties.sshKey
}
