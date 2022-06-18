mapping "azurerm_lb" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/loadBalancer.json"

  sku = LoadBalancerSku.name
}
