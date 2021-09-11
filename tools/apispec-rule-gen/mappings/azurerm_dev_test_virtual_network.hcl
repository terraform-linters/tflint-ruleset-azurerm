mapping "azurerm_dev_test_virtual_network" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/DTL.json"

  resource_group_name = resourceGroupName
  description         = VirtualNetworkProperties.description
}
