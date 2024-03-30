mapping "azurerm_hpc_cache" {
  import_path = "azure-rest-api-specs/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/storagecache.json"

  name     = ResourceName
  location = Cache.location
}
