mapping "azurerm_redis_cache" {
  import_path = "azure-rest-api-specs/specification/redis/resource-manager/Microsoft.Cache/stable/2018-03-01/redis.json"

  location                  = RedisCreateParameters.location
  capacity                  = Sku.capacity
  family                    = Sku.family
  sku_name                  = Sku.name
  enable_non_ssl_port       = RedisCommonProperties.enableNonSslPort
  minimum_tls_version       = RedisCommonProperties.minimumTlsVersion
  private_static_ip_address = RedisCreateProperties.staticIP
  shard_count               = RedisCommonProperties.shardCount
  subnet_id                 = RedisCreateProperties.subnetId
  zones                     = RedisCreateParameters.zones
}
