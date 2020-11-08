mapping "azurerm_cdn_endpoint" {
  import_path = "azure-rest-api-specs/specification/cdn/resource-manager/Microsoft.Cdn/stable/2019-04-15/cdn.json"

  resource_group_name           = resourceGroupNameParameter
  is_http_allowed               = EndpointPropertiesUpdateParameters.isHttpAllowed
  is_https_allowed              = EndpointPropertiesUpdateParameters.isHttpsAllowed
  content_types_to_compress     = EndpointPropertiesUpdateParameters.contentTypesToCompress
  is_compression_enabled        = EndpointPropertiesUpdateParameters.isCompressionEnabled
  querystring_caching_behaviour = QueryStringCachingBehavior
  optimization_type             = OptimizationType
  origin_host_header            = EndpointPropertiesUpdateParameters.originHostHeader
  origin_path                   = EndpointPropertiesUpdateParameters.originPath
  probe_path                    = EndpointPropertiesUpdateParameters.probePath
}
