mapping "azurerm_dns_a_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name               = RecordSet.name
  TTL                = RecordSetProperties.TTL
  target_resource_id = SubResource.id
}

mapping "azurerm_dns_aaaa_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name               = RecordSet.name
  TTL                = RecordSetProperties.TTL
  target_resource_id = SubResource.id
}

mapping "azurerm_dns_caa_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
  ttl  = RecordSetProperties.TTL
}

mapping "azurerm_dns_cname_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name               = RecordSet.name
  TTL                = RecordSetProperties.TTL
  target_resource_id = SubResource.id
}

mapping "azurerm_dns_mx_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
  ttl  = RecordSetProperties.TTL
}

mapping "azurerm_dns_ns_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
  ttl  = RecordSetProperties.TTL
}

mapping "azurerm_dns_ptr_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
  ttl  = RecordSetProperties.TTL
}

mapping "azurerm_dns_srv_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
  ttl  = RecordSetProperties.TTL
}

mapping "azurerm_dns_txt_record" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
  ttl  = RecordSetProperties.TTL
}

mapping "azurerm_dns_zone" {
  import_path = "azure-rest-api-specs/specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json"

  name = RecordSet.name
}
