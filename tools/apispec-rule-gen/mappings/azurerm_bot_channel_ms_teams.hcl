mapping "azurerm_bot_channel_ms_teams" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  calling_web_hook    = MsTeamsChannelProperties.callingWebHook
  enable_calling      = MsTeamsChannelProperties.enableCalling
}
