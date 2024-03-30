mapping "azurerm_lb" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/loadBalancer.json"

  sku = LoadBalancerSku.name
}
