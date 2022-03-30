mapping "azurerm_image" {
  import_path = "azure-rest-api-specs/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/compute.json"

  zone_resilient     = ImageStorageProfile.zoneResilient
  hyper_v_generation = HyperVGenerationType
}
