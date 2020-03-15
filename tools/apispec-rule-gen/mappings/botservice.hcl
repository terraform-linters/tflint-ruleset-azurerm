mapping "azurerm_bot_channel_email" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  email_address       = EmailChannelProperties.emailAddress
  email_password      = EmailChannelProperties.password
}

mapping "azurerm_bot_channel_ms_teams" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  calling_web_hook    = MsTeamsChannelProperties.callingWebHook
  enable_calling      = MsTeamsChannelProperties.enableCalling
}

mapping "azurerm_bot_channel_slack" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  client_id           = SlackChannelProperties.clientId
  client_secret       = SlackChannelProperties.clientSecret
  verification_token  = SlackChannelProperties.verificationToken
  landing_page_url    = SlackChannelProperties.landingPageUrl
}

mapping "azurerm_bot_channel_directline" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
}

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

mapping "azurerm_bot_connection" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  name                  = connectionNameParameter
  resource_group_name   = resourceGroupNameParameter
  bot_name              = resourceNameParameter
  service_provider_name = ServiceProviderProperties.serviceProviderName
  client_id             = ConnectionSettingProperties.clientId
  client_secret         = ConnectionSettingProperties.clientSecret
  scopes                = ConnectionSettingProperties.scopes
  parameters            = ConnectionSettingProperties.parameters
}

mapping "azurerm_bot_web_app" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

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
