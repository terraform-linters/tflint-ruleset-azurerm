mapping "azurerm_private_dns_a_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/PrivateDns/stable/2024-06-01/privatedns.json"

  TTL     = RecordSetProperties.ttl
  records = RecordSetProperties.aRecords
}
