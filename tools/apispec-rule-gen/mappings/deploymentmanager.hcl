mapping "azurerm_template_deployment" {
  import_path = "azure-rest-api-specs/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/deploymentmanager.json"

  resource_group_name = ServiceUnitProperties.targetResourceGroup
  deployment_mode     = ServiceUnitProperties.deploymentMode
}
