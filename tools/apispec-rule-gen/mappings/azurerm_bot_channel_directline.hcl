mapping "azurerm_bot_channel_directline" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
}