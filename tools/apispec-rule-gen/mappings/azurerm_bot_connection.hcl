mapping "azurerm_bot_connection" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/botservice.json"

  name                  = connectionNameParameter
  resource_group_name   = resourceGroupNameParameter
  bot_name              = resourceNameParameter
  service_provider_name = ServiceProviderProperties.serviceProviderName
  client_id             = ConnectionSettingProperties.clientId
  client_secret         = ConnectionSettingProperties.clientSecret
  scopes                = ConnectionSettingProperties.scopes
  parameters            = ConnectionSettingProperties.parameters
}
