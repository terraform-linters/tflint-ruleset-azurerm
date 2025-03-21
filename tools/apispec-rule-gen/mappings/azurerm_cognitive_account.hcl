mapping "azurerm_cognitive_account" {
  import_path = "azure-rest-api-specs/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/cognitiveservices.json"

  name                 = accountNameParameter
  location             = locationParameter
  kind                 = Kind
  sku_name             = SkuName
  qna_runtime_endpoint = ApiProperties.qnaRuntimeEndpoint
}
