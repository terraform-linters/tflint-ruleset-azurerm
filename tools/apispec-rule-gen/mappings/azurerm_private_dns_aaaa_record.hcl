mapping "azurerm_private_dns_aaaa_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  TTL     = RecordSetProperties.ttl
  records = RecordSetProperties.aaaaRecords
}
