mapping "azurerm_relay_namespace" {
  import_path = "azure-rest-api-specs/specification/relay/resource-manager/Microsoft.Relay/Relay/stable/2021-11-01/Namespaces.json"

  name                = any //namespaceNameParameter
  resource_group_name = any //resourceGroupNameParameter
  sku_name            = Sku.name
}
