mapping "azurerm_sql_database" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/databases.json"

  server_name                      = ServerNameParameter
  create_mode                      = DatabaseProperties.createMode
  source_database_id               = DatabaseProperties.sourceDatabaseId
  restore_point_in_time            = DatabaseProperties.restorePointInTime
  edition                          = any
  collation                        = DatabaseProperties.collation
  max_size_bytes                   = DatabaseProperties.maxSizeBytes
  requested_service_objective_id   = any
  requested_service_objective_name = any
  source_database_deletion_date    = DatabaseProperties.sourceDatabaseDeletionDate
  elastic_pool_name                = any
  read_scale                       = any
  zone_redundant                   = DatabaseProperties.zoneRedundant
}