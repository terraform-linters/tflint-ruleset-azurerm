mapping "azurerm_iothub_dps_certificate" {
  import_path = "azure-rest-api-specs/specification/iothub/resource-manager/Microsoft.Devices/stable/2020-03-01/iothub.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iot_dps_name        = resourceName
  certificate_content = CertificateProperties.certificate
}
