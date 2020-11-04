mapping "azurerm_stream_analytics_job" {
  import_path = "azure-rest-api-specs/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2016-03-01/streamingjobs.json"

  name                                     = StreamingJobNameParameter
  resource_group_name                      = ResourceGroupNameParameter
  location                                 = Resource.location
  compatibility_level                      = CompatibilityLevel
  data_locale                              = StreamingJobProperties.dataLocale
  events_late_arrival_max_delay_in_seconds = StreamingJobProperties.eventsLateArrivalMaxDelayInSeconds
  events_out_of_order_max_delay_in_seconds = StreamingJobProperties.eventsOutOfOrderMaxDelayInSeconds
  events_out_of_order_policy               = EventsOutOfOrderPolicy
  output_error_policy                      = OutputErrorPolicy
}