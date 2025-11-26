mapping "azurerm_bot_channel_email" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/BotService/stable/2022-09-15/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  email_address       = EmailChannelProperties.emailAddress
  email_password      = EmailChannelProperties.password
}
