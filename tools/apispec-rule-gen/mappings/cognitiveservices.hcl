mapping "azurerm_cognitive_account" {
  import_path = "azure-rest-api-specs/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2017-04-18/cognitiveservices.json"

  name                 = CognitiveServicesAccount.name
  location             = locationParameter
  kind                 = CognitiveServicesAccountKind
  sku_name             = SkuName
  qna_runtime_endpoint = CognitiveServicesAccountApiProperties.qnaRuntimeEndpoint
}
