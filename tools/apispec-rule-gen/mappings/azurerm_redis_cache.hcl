mapping "azurerm_redis_cache" {
  import_path = "azure-rest-api-specs/specification/redis/resource-manager/Microsoft.Cache/Redis/stable/2024-11-01/redis.json"

  location                  = RedisCreateParameters.location
  capacity                  = Sku.capacity
  family                    = SkuFamily
  sku_name                  = SkuName
  enable_non_ssl_port       = RedisCommonProperties.enableNonSslPort
  minimum_tls_version       = TlsVersion
  private_static_ip_address = RedisCreateProperties.staticIP
  shard_count               = RedisCommonProperties.shardCount
  subnet_id                 = RedisCreateProperties.subnetId
  zones                     = RedisCreateParameters.zones
}
