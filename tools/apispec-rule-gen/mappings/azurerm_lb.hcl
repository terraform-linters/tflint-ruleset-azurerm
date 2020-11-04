mapping "azurerm_lb" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/loadBalancer.json"

  sku = LoadBalancerSku.name
}
