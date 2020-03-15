mapping "azurerm_private_dns_a_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  TTL     = RecordSetProperties.ttl
  records = RecordSetProperties.aRecords
}

mapping "azurerm_private_dns_aaaa_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  TTL     = RecordSetProperties.ttl
  records = RecordSetProperties.aaaaRecords
}

mapping "azurerm_private_dns_cname_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  TTL    = RecordSetProperties.ttl
  record = CnameRecord.cname
}

mapping "azurerm_private_dns_mx_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  ttl = RecordSetProperties.ttl
}

mapping "azurerm_private_dns_ptr_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  ttl     = RecordSetProperties.ttl
  records = RecordSetProperties.ptrRecords
}

mapping "azurerm_private_dns_srv_record" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  ttl = RecordSetProperties.ttl
}

mapping "azurerm_private_dns_txtrecord" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  ttl = RecordSetProperties.ttl
}

mapping "azurerm_private_dns_zone_virtual_network_link" {
  import_path = "azure-rest-api-specs/specification/privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json"

  registration_enabled = VirtualNetworkLinkProperties.registrationEnabled
}
