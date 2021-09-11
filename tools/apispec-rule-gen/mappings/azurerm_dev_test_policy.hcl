mapping "azurerm_dev_test_policy" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/DTL.json"

  name                = PolicyProperties.factName
  resource_group_name = resourceGroupName
  location            = locationName
  description         = PolicyProperties.description
  evaluator_type      = PolicyProperties.evaluatorType
  threshold           = PolicyProperties.threshold
  fact_data           = PolicyProperties.factData
}
