mapping "azurerm_dev_test_policy" {
  import_path = "azure-rest-api-specs/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/DevTestLabs/stable/2018-09-15/DTL.json"

  name                = PolicyFactName
  resource_group_name = any //resourceGroupName
  location            = any //locationName
  description         = PolicyProperties.description
  evaluator_type      = PolicyEvaluatorType
  threshold           = PolicyProperties.threshold
  fact_data           = PolicyProperties.factData
}
