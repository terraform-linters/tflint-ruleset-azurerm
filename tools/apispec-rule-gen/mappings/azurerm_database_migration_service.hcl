mapping "azurerm_database_migration_service" {
  import_path = "azure-rest-api-specs/specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2021-06-30/definitions/Services.json"

  subnet_id = DataMigrationServiceProperties.virtualSubnetId
  sku_name  = ServiceSku.name
}
