mapping "azurerm_private_dns_srv_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json"

  ttl = RecordSetProperties.ttl
}
