mapping "azurerm_dev_test_lab" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  artifacts_storage_account_id         = LabProperties.artifactsStorageAccount
  default_storage_account_id           = LabProperties.defaultStorageAccount
  default_premium_storage_account_id   = LabProperties.defaultPremiumStorageAccount
  key_vault_id                         = UserSecretStore.keyVaultId
  premium_data_disk_storage_account_id = LabProperties.premiumDataDiskStorageAccount
  unique_identifier                    = LabProperties.uniqueIdentifier
}

mapping "azurerm_dev_test_linux_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  resource_group_name        = resourceGroupName
  location                   = locationName
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

mapping "azurerm_dev_test_policy" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  name                = PolicyProperties.factName
  resource_group_name = resourceGroupName
  location            = locationName
  description         = PolicyProperties.description
  evaluator_type      = PolicyProperties.evaluatorType
  threshold           = PolicyProperties.threshold
  fact_data           = PolicyProperties.factData
}

mapping "azurerm_dev_test_schedule" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  location            = locationName
  resource_group_name = resourceGroupName
  status              = ScheduleProperties.status
  task_type           = ScheduleProperties.taskType
  time_zone_id        = ScheduleProperties.timeZoneId
}

mapping "azurerm_dev_test_virtual_network" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  resource_group_name = resourceGroupName
  description         = VirtualNetworkProperties.description
}

mapping "azurerm_dev_test_windows_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2016-05-15/DTL.json"

  resource_group_name        = resourceGroupName
  location                   = locationName
  lab_subnet_name            = LabVirtualMachineCreationParameterProperties.labSubnetName
  lab_virtual_network_id     = LabVirtualMachineCreationParameterProperties.labVirtualNetworkId
  size                       = LabVirtualMachineCreationParameterProperties.size
  storage_type               = LabVirtualMachineCreationParameterProperties.storageType
  username                   = LabVirtualMachineCreationParameterProperties.userName
  allow_claim                = LabVirtualMachineCreationParameterProperties.allowClaim
  disallow_public_ip_address = LabVirtualMachineCreationParameterProperties.disallowPublicIpAddress
  notes                      = LabVirtualMachineCreationParameterProperties.notes
  password                   = LabVirtualMachineCreationParameterProperties.password
}
