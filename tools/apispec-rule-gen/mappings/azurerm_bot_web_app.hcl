mapping "azurerm_bot_web_app" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/botservice.json"

  name                                  = resourceNameParameter
  resource_group_name                   = resourceGroupNameParameter
  sku                                   = SkuName
  microsoft_app_id                      = BotProperties.msaAppId
  display_name                          = BotProperties.displayName
  endpoint                              = BotProperties.endpoint
  developer_app_insights_key            = BotProperties.developerAppInsightKey
  developer_app_insights_api_key        = BotProperties.developerAppInsightsApiKey
  developer_app_insights_application_id = BotProperties.developerAppInsightsApplicationId
  luis_app_ids                          = BotProperties.luisAppIds
  luis_key                              = BotProperties.luisKey
}
