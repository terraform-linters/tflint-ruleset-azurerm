mapping "azurerm_iothub" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name                        = resourceName
  resource_group_name         = resourceGroupName
  event_hub_partition_count   = EventHubProperties.partitionCount
  event_hub_retention_in_days = EventHubProperties.retentionTimeInDays
}

mapping "azurerm_iothub_consumer_group" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name                = resourceName
  iothub_name         = resourceName
  resource_group_name = resourceGroupName
}

mapping "azurerm_iothub_dps" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
}

mapping "azurerm_iothub_dps_certificate" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iot_dps_name        = resourceName
  certificate_content = CertificateProperties.certificate
}

mapping "azurerm_iothub_dps_shared_access_policy" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iothub_dps_name     = resourceName
}

mapping "azurerm_iothub_endpoint_eventhub" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name              = resourceName
  connection_string = RoutingEventHubProperties.connectionString
}

mapping "azurerm_iothub_endpoint_servicebus_queue" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name              = resourceName
  connection_string = RoutingServiceBusQueueEndpointProperties.connectionString
} 

mapping "azurerm_iothub_endpoint_servicebus_topic" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name              = resourceName
  connection_string = RoutingServiceBusTopicEndpointProperties.connectionString
}

mapping "azurerm_iothub_endpoint_storage_container" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

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

mapping "azurerm_iothub_fallback_route" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  resource_group_name = resourceGroupName
  iothub_name         = resourceName
  enabled             = FallbackRouteProperties.isEnabled
  endpoint_names      = FallbackRouteProperties.endpointNames
  condition           = FallbackRouteProperties.condition
}

mapping "azurerm_iothub_route" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name                = RouteProperties.name
  resource_group_name = resourceGroupName
  iothub_name         = resourceName
  source              = RouteProperties.source
  condition           = RouteProperties.condition
  endpoint_names      = RouteProperties.endpointNames
  enabled             = RouteProperties.isEnabled
}

mapping "azurerm_iothub_shared_access_policy" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/preview/2019-03-22-preview/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iothub_name         = resourceName
}
