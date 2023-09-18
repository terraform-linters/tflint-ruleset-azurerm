mapping "azurerm_image" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/computeRPCommon.json"

  zone_resilient     = any //ImageStorageProfile.zoneResilient
  hyper_v_generation = HyperVGenerationType
}
