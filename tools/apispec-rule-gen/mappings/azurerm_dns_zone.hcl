mapping "azurerm_dns_zone" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/Dns/stable/2018-05-01/dns.json"

  name = RecordSet.name
}
