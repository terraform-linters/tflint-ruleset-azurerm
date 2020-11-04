mapping "azurerm_batch_application" {
  import_path = "azure-rest-api-specs/specification/batch/resource-manager/Microsoft.Batch/stable/2019-08-01/BatchManagement.json"

  name                = ApplicationNameParameter
  resource_group_name = ResourceGroupNameParameter
  account_name        = AccountNameParameter
  allow_updates       = ApplicationProperties.allowUpdates
  default_version     = ApplicationProperties.defaultVersion
  display_name        = ApplicationProperties.displayName
}