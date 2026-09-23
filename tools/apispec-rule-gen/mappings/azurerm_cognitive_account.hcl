mapping "azurerm_cognitive_account" {
  import_path = "azure-rest-api-specs/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2026-03-01/cognitiveservices.json"

  name                 = accountName
  location             = LocationParameter
  kind                 = Account.kind
  sku_name             = Sku.name
  qna_runtime_endpoint = ApiProperties.qnaRuntimeEndpoint
}
