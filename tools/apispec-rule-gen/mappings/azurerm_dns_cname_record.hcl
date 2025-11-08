mapping "azurerm_dns_cname_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/Dns/stable/2018-05-01/dns.json"

  name               = RecordSet.name
  TTL                = RecordSetProperties.TTL
  target_resource_id = SubResource.id
}
