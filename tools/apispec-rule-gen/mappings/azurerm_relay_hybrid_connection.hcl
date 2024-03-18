mapping "azurerm_relay_hybrid_connection" {
  import_path = "azure-rest-api-specs/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/hybridConnections.json"

  name                          = any //hybridConnectionNameParameter
  resource_group_name           = any //resourceGroupNameParameter
  relay_namespace_name          = any //namespaceNameParameter
}
