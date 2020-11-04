mapping "azurerm_virtual_wan" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-05-01/virtualWan.json"

  disable_vpn_encryption            = VirtualWanProperties.disableVpnEncryption
  allow_branch_to_branch_traffic    = VirtualWanProperties.allowBranchToBranchTraffic
  allow_vnet_to_vnet_traffic        = VirtualWanProperties.allowVnetToVnetTraffic
  office365_local_breakout_category = OfficeTrafficCategory
  type                              = VirtualWanProperties.type
}
