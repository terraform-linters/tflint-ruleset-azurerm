mapping "azurerm_dns_caa_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
  ttl  = RecordSetProperties.TTL
}
