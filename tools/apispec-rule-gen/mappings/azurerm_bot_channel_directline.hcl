mapping "azurerm_bot_channel_directline" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
}
