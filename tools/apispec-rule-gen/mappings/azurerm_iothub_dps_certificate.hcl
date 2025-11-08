mapping "azurerm_iothub_dps_certificate" {
  import_path = "azure-rest-api-specs/specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/DeviceProvisioningServices/stable/2022-02-05/iotdps.json"

  name                = resourceName
  resource_group_name = resourceGroupName
  iot_dps_name        = resourceName
  certificate_content = CertificateProperties.certificate
}
