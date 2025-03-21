mapping "azurerm_batch_certificate" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/BatchManagement.json"

  account_name        = AccountNameParameter
  resource_group_name = ResourceGroupNameParameter
  certificate         = CertificateCreateOrUpdateProperties.data
  format              = CertificateBaseProperties.format
  password            = CertificateCreateOrUpdateProperties.password
  thumbprint          = CertificateBaseProperties.thumbprintAlgorithm
}
