mapping "azurerm_redis_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json"

  start_ip = RedisFirewallRuleProperties.startIP
  end_ip   = RedisFirewallRuleProperties.endIP
}
