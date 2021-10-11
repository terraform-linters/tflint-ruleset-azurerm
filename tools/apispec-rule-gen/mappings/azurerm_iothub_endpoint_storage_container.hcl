mapping "azurerm_iothub_endpoint_storage_container" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-03-31/iothub.json"

  name                       = RoutingStorageContainerProperties.name
  resource_group_name        = RoutingStorageContainerProperties.resourceGroup
  iothub_name                = resourceName
  connection_string          = RoutingStorageContainerProperties.connectionString
  batch_frequency_in_seconds = RoutingStorageContainerProperties.batchFrequencyInSeconds
  max_chunk_size_in_bytes    = RoutingStorageContainerProperties.maxChunkSizeInBytes
  container_name             = RoutingStorageContainerProperties.containerName
  encoding                   = RoutingStorageContainerProperties.encoding
  file_name_format           = RoutingStorageContainerProperties.fileNameFormat
}
