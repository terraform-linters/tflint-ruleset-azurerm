mapping "azurerm_relay_namespace" {
  import_path = "azure-rest-api-specs/specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/relay.json"

  name                = namespaceNameParameter
  resource_group_name = resourceGroupNameParameter
  sku_name            = Sku.name
}
