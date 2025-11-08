mapping "azurerm_iothub" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/IoTHub/preview/2022-04-30-preview/iothub.json"

  name                        = resourceName
  resource_group_name         = resourceGroupName
  event_hub_partition_count   = EventHubProperties.partitionCount
  event_hub_retention_in_days = EventHubProperties.retentionTimeInDays
}
