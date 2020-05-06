mapping "azurerm_availability_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  platform_update_domain_count = AvailabilitySetProperties.platformUpdateDomainCount
  platform_fault_domain_count  = AvailabilitySetProperties.platformFaultDomainCount
}

mapping "azurerm_dedicated_host" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  sku_name                = Sku.name
  platform_fault_domain   = DedicatedHostProperties.platformFaultDomain
  auto_replace_on_failure = DedicatedHostProperties.autoReplaceOnFailure
  license_type            = DedicatedHostLicenseType
}

mapping "azurerm_dedicated_host_group" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  platform_fault_domain_count = DedicatedHostGroupProperties.platformFaultDomainCount
  zones                       = DedicatedHostGroup.zones
}

mapping "azurerm_disk_encryption_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-11-01/disk.json"

  name                = DiskEncryptionSetNameParameter
  resource_group_name = ResourceGroupNameParameter
}

mapping "azurerm_image" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  zone_resilient     = ImageStorageProfile.zoneResilient
  hyper_v_generation = HyperVGenerationType
}

mapping "azurerm_linux_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  admin_username                  = OSProfile.adminUsername
  size                            = HardwareProfile.vmSize
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

mapping "azurerm_linux_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  admin_username                                    = VirtualMachineScaleSetOSProfile.adminUsername
  sku                                               = HardwareProfile.vmSize
  admin_password                                    = VirtualMachineScaleSetOSProfile.adminPassword
  computer_name_prefix                              = VirtualMachineScaleSetOSProfile.computerNamePrefix
  custom_data                                       = VirtualMachineScaleSetOSProfile.customData
  disable_password_authentication                   = LinuxConfiguration.disablePasswordAuthentication
  do_not_run_extensions_on_overprovisioned_machines = VirtualMachineScaleSetProperties.doNotRunExtensionsOnOverprovisionedVMs
  eviction_policy                                   = evictionPolicy
  max_bid_price                                     = BillingProfile.maxPrice
  overprovision                                     = VirtualMachineScaleSetProperties.overprovision
  priority                                          = priority
  provision_vm_agent                                = LinuxConfiguration.provisionVMAgent
  scale_in_policy                                   = ScaleInPolicy.rules
  single_placement_group                            = VirtualMachineScaleSetProperties.singlePlacementGroup
  upgrade_mode                                      = UpgradePolicy.mode
  zone_balance                                      = VirtualMachineScaleSetProperties.zoneBalance
}

mapping "azurerm_managed_application_definition" {
  import_path = "azure-rest-api-specs/specification/resources/resource-manager/Microsoft.Solutions/stable/2019-07-01/managedapplications.json"

  lock_level           = ApplicationLockLevel
  create_ui_definition = ApplicationDefinitionProperties.createUiDefinition
  display_name         = ApplicationDefinitionProperties.displayName
  description          = ApplicationDefinitionProperties.description
  package_enabled      = ApplicationDefinitionProperties.isEnabled
  main_template        = ApplicationDefinitionProperties.mainTemplate
  package_file_uri     = ApplicationDefinitionProperties.packageFileUri
}

mapping "azurerm_managed_disk" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-11-01/disk.json"

  name                 = DiskNameParameter
  resource_group_name  = ResourceGroupNameParameter
  storage_account_type = DiskSku.name
  create_option        = CreationData.createOption
  disk_iops_read_write = DiskProperties.diskIOPSReadWrite
  disk_mbps_read_write = DiskProperties.diskMBpsReadWrite
  disk_size_gb         = DiskProperties.diskSizeGB
  os_type              = DiskProperties.osType
  source_resource_id   = CreationData.sourceResourceId
  source_uri           = CreationData.sourceUri
}

mapping "azurerm_shared_image" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/gallery.json"

  resource_group_name   = ResourceGroupNameParameter
  os_type               = GalleryImageProperties.osType
  description           = GalleryImageProperties.description
  eula                  = GalleryImageProperties.eula
  privacy_statement_uri = GalleryImageProperties.privacyStatementUri
  release_note_uri      = GalleryImageProperties.releaseNoteUri
}

mapping "azurerm_shared_image_gallery" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/gallery.json"

  resource_group_name = ResourceGroupNameParameter
  description         = GalleryImageProperties.description
}

mapping "azurerm_shared_image_version" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/gallery.json"

  resource_group_name = ResourceGroupNameParameter
  exclude_from_latest = GalleryArtifactPublishingProfileBase.excludeFromLatest
}

mapping "azurerm_snapshot" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-11-01/disk.json"

  name                = SnapshotNameParameter
  resource_group_name = ResourceGroupNameParameter
  create_option       = CreationData.createOption
  source_uri          = CreationData.sourceUri
  source_resource_id  = CreationData.sourceResourceId
  storage_account_id  = CreationData.storageAccountId
  disk_size_gb        = DiskProperties.diskSizeGB
}

mapping "azurerm_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  vm_size      = HardwareProfile.vmSize
  license_type = VirtualMachineProperties.licenseType
}

mapping "azurerm_virtual_machine_data_disk_attachment" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  lun                       = VirtualMachineScaleSetDataDisk.lun
  caching                   = Caching
  create_option             = CreateOption
  write_accelerator_enabled = VirtualMachineScaleSetDataDisk.writeAcceleratorEnabled
}

mapping "azurerm_virtual_machine_extension" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  publisher                  = VirtualMachineExtensionProperties.publisher
  type                       = VirtualMachineExtensionProperties.type
  type_handler_version       = VirtualMachineExtensionProperties.typeHandlerVersion
  auto_upgrade_minor_version = VirtualMachineExtensionProperties.autoUpgradeMinorVersion
  settings                   = VirtualMachineExtensionProperties.settings
  protected_settings         = VirtualMachineExtensionProperties.protectedSettings
}

mapping "azurerm_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  upgrade_policy_mode    = UpgradePolicy.mode
  eviction_policy        = evictionPolicy
  license_type           = VirtualMachineProperties.licenseType
  overprovision          = VirtualMachineScaleSetProperties.overprovision
  priority               = priority
  single_placement_group = VirtualMachineScaleSetProperties.singlePlacementGroup
}

mapping "azurerm_virtual_machine_scale_set_extension" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  publisher                  = VirtualMachineScaleSetExtensionProperties.publisher
  type                       = VirtualMachineScaleSetExtensionProperties.type
  type_handler_version       = VirtualMachineScaleSetExtensionProperties.typeHandlerVersion
  auto_upgrade_minor_version = VirtualMachineScaleSetExtensionProperties.autoUpgradeMinorVersion
  force_update_tag           = VirtualMachineScaleSetExtensionProperties.forceUpdateTag
  protected_settings         = VirtualMachineScaleSetExtensionProperties.protectedSettings
  provision_after_extensions = VirtualMachineScaleSetExtensionProperties.provisionAfterExtensions
  settings                   = VirtualMachineScaleSetExtensionProperties.settings
}

mapping "azurerm_windows_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  admin_password             = OSProfile.adminPassword
  admin_username             = OSProfile.adminUsername
  size                       = HardwareProfile.vmSize
  allow_extension_operations = OSProfile.allowExtensionOperations
  computer_name              = OSProfile.computerName
  custom_data                = OSProfile.customData
  enable_automatic_updates   = WindowsConfiguration.enableAutomaticUpdates
  eviction_policy            = evictionPolicy
  license_type               = VirtualMachineProperties.licenseType
  max_bid_price              = BillingProfile.maxPrice
  priority                   = priority
  provision_vm_agent         = WindowsConfiguration.provisionVMAgent
}

mapping "azurerm_windows_virtual_machine_scale_set" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2019-12-01/compute.json"

  admin_password                                    = VirtualMachineScaleSetOSProfile.adminPassword
  admin_username                                    = VirtualMachineScaleSetOSProfile.adminUsername
  sku                                               = HardwareProfile.vmSize
  computer_name_prefix                              = VirtualMachineScaleSetOSProfile.computerNamePrefix
  custom_data                                       = VirtualMachineScaleSetOSProfile.customData
  do_not_run_extensions_on_overprovisioned_machines = VirtualMachineScaleSetProperties.doNotRunExtensionsOnOverprovisionedVMs
  enable_automatic_updates                          = WindowsConfiguration.enableAutomaticUpdates
  eviction_policy                                   = evictionPolicy
  license_type                                      = VirtualMachineProperties.licenseType
  max_bid_price                                     = BillingProfile.maxPrice
  overprovision                                     = VirtualMachineScaleSetProperties.overprovision
  priority                                          = priority
  provision_vm_agent                                = WindowsConfiguration.provisionVMAgent
  scale_in_policy                                   = ScaleInPolicy.rules
  single_placement_group                            = VirtualMachineScaleSetProperties.singlePlacementGroup
  upgrade_mode                                      = UpgradePolicy.mode
  zone_balance                                      = VirtualMachineScaleSetProperties.zoneBalance
}
