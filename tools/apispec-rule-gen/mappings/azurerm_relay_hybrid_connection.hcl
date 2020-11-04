mapping "azurerm_relay_hybrid_connection" {
  import_path = "azure-rest-api-specs/specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/relay.json"

  name                          = hybridConnectionNameParameter
  resource_group_name           = resourceGroupNameParameter
  relay_namespace_name          = namespaceNameParameter
}
