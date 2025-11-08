mapping "azurerm_bot_channel_slack" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/BotService/preview/2021-05-01-preview/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  client_id           = SlackChannelProperties.clientId
  client_secret       = SlackChannelProperties.clientSecret
  verification_token  = SlackChannelProperties.verificationToken
  landing_page_url    = SlackChannelProperties.landingPageUrl
}
