mapping "azurerm_managed_application_definition" {
  import_path = "azure-rest-api-specs/specification/resources/resource-manager/Microsoft.Solutions/stable/2019-07-01/managedapplications.json"

  lock_level           = ApplicationLockLevel
  create_ui_definition = ApplicationDefinitionProperties.createUiDefinition
  display_name         = ApplicationDefinitionProperties.displayName
  description          = ApplicationDefinitionProperties.description
  package_enabled      = ApplicationDefinitionProperties.isEnabled
  main_template        = ApplicationDefinitionProperties.mainTemplate
  package_file_uri     = ApplicationDefinitionProperties.packageFileUri
}
