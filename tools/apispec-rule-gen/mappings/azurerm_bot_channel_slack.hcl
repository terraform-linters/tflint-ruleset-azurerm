mapping "azurerm_bot_channel_slack" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  client_id           = SlackChannelProperties.clientId
  client_secret       = SlackChannelProperties.clientSecret
  verification_token  = SlackChannelProperties.verificationToken
  landing_page_url    = SlackChannelProperties.landingPageUrl
}
