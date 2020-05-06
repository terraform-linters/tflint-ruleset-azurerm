mapping "azurerm_mariadb_configuration" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name        = ConfigurationNameParameter
  server_name = ServerNameParameter
  value       = ConfigurationProperties.value
}

mapping "azurerm_mariadb_database" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name        = DatabaseNameParameter
  server_name = ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}

mapping "azurerm_mariadb_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name             = FirewallRuleNameParameter
  server_name      = ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}

mapping "azurerm_mariadb_server" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name                         = ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerPropertiesForDefaultCreate.administratorLogin
  administrator_login_password = ServerPropertiesForDefaultCreate.administratorLoginPassword
  version                      = ServerVersion
  ssl_enforcement              = SslEnforcement
}

mapping "azurerm_mariadb_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/mariadb.json"

  name        = virtualNetworkRuleNameParameter
  server_name = ServerNameParameter
  subnet_id   = VirtualNetworkRuleProperties.virtualNetworkSubnetId
}

mapping "azurerm_mysql_configuration" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name        = ConfigurationNameParameter
  server_name = ServerNameParameter
  value       = ConfigurationProperties.value
}

mapping "azurerm_mysql_database" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name        = DatabaseNameParameter
  server_name = ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}

mapping "azurerm_mysql_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name             = FirewallRuleNameParameter
  server_name      = ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}

mapping "azurerm_mysql_server" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name                         = ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerPropertiesForDefaultCreate.administratorLogin
  administrator_login_password = ServerPropertiesForDefaultCreate.administratorLoginPassword
  version                      = ServerVersion
  ssl_enforcement              = SslEnforcement
}

mapping "azurerm_mysql_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/mysql.json"

  name        = virtualNetworkRuleNameParameter
  server_name = ServerNameParameter
  subnet_id   = VirtualNetworkRuleProperties.virtualNetworkSubnetId
}

mapping "azurerm_postgresql_configuration" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name        = ConfigurationNameParameter
  server_name = ServerNameParameter
  value       = ConfigurationProperties.value
}

mapping "azurerm_postgresql_database" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name        = DatabaseNameParameter
  server_name = ServerNameParameter
  charset     = DatabaseProperties.charset
  collation   = DatabaseProperties.collation
}

mapping "azurerm_postgresql_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name             = FirewallRuleNameParameter
  server_name      = ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}

mapping "azurerm_postgresql_server" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name                         = ServerNameParameter
  sku_name                     = Sku.name
  administrator_login          = ServerPropertiesForDefaultCreate.administratorLogin
  administrator_login_password = ServerPropertiesForDefaultCreate.administratorLoginPassword
  version                      = ServerVersion
  ssl_enforcement              = SslEnforcement
}

mapping "azurerm_postgresql_virtual_network_rule" {
  import_path = "azure-rest-api-specs/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/postgresql.json"

  name        = virtualNetworkRuleNameParameter
  server_name = ServerNameParameter
  subnet_id   = VirtualNetworkRuleProperties.virtualNetworkSubnetId
}

mapping "azurerm_sql_database" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/databases.json"

  server_name                      = ServerNameParameter
  create_mode                      = DatabaseProperties.createMode
  source_database_id               = DatabaseProperties.sourceDatabaseId
  restore_point_in_time            = DatabaseProperties.restorePointInTime
  edition                          = DatabaseProperties.edition
  collation                        = DatabaseProperties.collation
  max_size_bytes                   = DatabaseProperties.maxSizeBytes
  requested_service_objective_id   = DatabaseProperties.requestedServiceObjectiveId
  requested_service_objective_name = DatabaseProperties.requestedServiceObjectiveName
  source_database_deletion_date    = DatabaseProperties.sourceDatabaseDeletionDate
  elastic_pool_name                = DatabaseProperties.elasticPoolName
  read_scale                       = any // DatabaseProperties.readScale
  zone_redundant                   = DatabaseProperties.zoneRedundant
}

mapping "azurerm_mssql_database" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/databases.json"

  name                        = DatabaseNameParameter
  auto_pause_delay_in_minutes = DatabaseProperties.autoPauseDelay
  create_mode                 = DatabaseProperties.createMode
  creation_source_database_id = DatabaseProperties.sourceDatabaseId
  collation                   = DatabaseProperties.collation
  elastic_pool_id             = DatabaseProperties.elasticPoolId
  license_type                = DatabaseProperties.licenseType
  min_capacity                = DatabaseProperties.minCapacity
  restore_point_in_time       = DatabaseProperties.restorePointInTime
  read_replica_count          = DatabaseProperties.readReplicaCount
  read_scale                  = any // DatabaseProperties.readScale
  sample_name                 = DatabaseProperties.sampleName
  sku_name                    = Sku.name
  zone_redundant              = DatabaseProperties.zoneRedundant
}

mapping "azurerm_sql_active_directory_administrator" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/serverAzureADAdministrators.json"

  server_name = ServerNameParameter
  login       = ServerAdministratorProperties.login
  tenant_id   = ServerAdministratorProperties.tenantId
}

mapping "azurerm_sql_elasticpool" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/elasticPools.json"

  server_name = ServerNameParameter
  edition     = ElasticPoolProperties.edition
  dtu         = ElasticPoolProperties.dtu
  db_dtu_min  = ElasticPoolProperties.databaseDtuMin
  db_dtu_max  = ElasticPoolProperties.databaseDtuMax
}

mapping "azurerm_mssql_elasticpool" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/elasticPools.json"

  resource_group_name = ResourceGroupParameter
  server_name         = ServerNameParameter
  max_size_bytes      = ElasticPoolProperties.maxSizeBytes
  zone_redundant      = ElasticPoolProperties.zoneRedundant
}

mapping "azurerm_sql_failover_group" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/preview/2015-05-01-preview/failoverGroups.json"

  resource_group_name = ResourceGroupParameter
  server_name         = ServerNameParameter
  databases           = FailoverGroupProperties.databases
}

mapping "azurerm_sql_firewall_rule" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/firewallRules.json"

  server_name      = ServerNameParameter
  start_ip_address = FirewallRuleProperties.startIpAddress
  end_ip_address   = FirewallRuleProperties.endIpAddress
}

mapping "azurerm_sql_server" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/servers.json"

  name                         = ServerNameParameter
  version                      = ServerProperties.version
  administrator_login          = ServerProperties.administratorLogin
  administrator_login_password = ServerProperties.administratorLoginPassword
}

mapping "azurerm_mssql_server_security_alert_policy" {
  import_path = "azure-rest-api-specs/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/databaseSecurityAlertPolicies.json"

  state                      = DatabaseSecurityAlertPolicyProperties.state
  disabled_alerts            = DatabaseSecurityAlertPolicyProperties.disabledAlerts
  email_addresses            = DatabaseSecurityAlertPolicyProperties.emailAddresses
  retention_days             = DatabaseSecurityAlertPolicyProperties.retentionDays
  storage_account_access_key = DatabaseSecurityAlertPolicyProperties.storageAccountAccessKey
  storage_endpoint           = DatabaseSecurityAlertPolicyProperties.storageEndpoint
}

mapping "azurerm_mssql_virtual_machine" {
  import_path = "azure-rest-api-specs/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/sqlvm.json"

  sql_license_type                 = SqlVirtualMachineProperties.sqlServerLicenseType
  r_services_enabled               = AdditionalFeaturesServerConfigurations.isRServicesEnabled
  sql_connectivity_port            = SqlConnectivityUpdateSettings.port
  sql_connectivity_type            = SqlConnectivityUpdateSettings.connectivityType
  sql_connectivity_update_password = SqlConnectivityUpdateSettings.sqlAuthUpdatePassword
  sql_connectivity_update_username = SqlConnectivityUpdateSettings.sqlAuthUpdateUserName
}
