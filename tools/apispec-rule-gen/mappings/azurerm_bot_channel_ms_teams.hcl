mapping "azurerm_bot_channel_ms_teams" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/BotService/preview/2021-05-01-preview/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  calling_web_hook    = MsTeamsChannelProperties.callingWebhook
  enable_calling      = MsTeamsChannelProperties.enableCalling
}
