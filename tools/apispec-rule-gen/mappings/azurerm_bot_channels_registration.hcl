mapping "azurerm_bot_channels_registration" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  resource_group_name                   = resourceGroupNameParameter
  sku                                   = SkuName
  microsoft_app_id                      = BotProperties.msaAppId
  display_name                          = BotProperties.displayName
  endpoint                              = BotProperties.endpoint
  developer_app_insights_key            = BotProperties.developerAppInsightKey
  developer_app_insights_api_key        = BotProperties.developerAppInsightsApiKey
  developer_app_insights_application_id = BotProperties.developerAppInsightsApplicationId
}