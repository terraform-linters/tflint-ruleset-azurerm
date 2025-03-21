mapping "azurerm_private_dns_aaaa_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/privatedns.json"

  TTL     = RecordSetProperties.ttl
  records = RecordSetProperties.aaaaRecords
}
