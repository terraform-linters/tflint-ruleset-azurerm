# azurerm_resources_missing_prevent_destroy

Require `lifecycle { prevent_destroy = true }` for Azure resources that contain data to protect them from accidental deletion.

## Configuration

```hcl
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
  resources = [] # (Optional) Override default resources list
  exclude = []   # (Optional) Exclude some resource types from prevent_destroy checks
}
```

## Default Resources

By default, this rule applies to the following Azure resources that typically contain data:

- `azurerm_backup_container_storage_account`, `azurerm_backup_policy_file_share`, `azurerm_backup_policy_vm`, `azurerm_backup_protected_vm`
- `azurerm_cosmosdb_account`, `azurerm_cosmosdb_cassandra_cluster`, `azurerm_cosmosdb_cassandra_table`, `azurerm_cosmosdb_mongo_database`, `azurerm_cosmosdb_postgresql_cluster`, `azurerm_cosmosdb_sql_container`, `azurerm_cosmosdb_sql_database`, `azurerm_cosmosdb_table`
- `azurerm_databricks_workspace`
- `azurerm_eventhub`, `azurerm_eventhub_namespace`
- `azurerm_iothub`
- `azurerm_key_vault`, `azurerm_key_vault_certificate`, `azurerm_key_vault_key`, `azurerm_key_vault_secret`
- `azurerm_mssql_database`, `azurerm_mssql_server`, `azurerm_mssql_virtual_machine`
- `azurerm_mysql_flexible_database`, `azurerm_mysql_flexible_server`
- `azurerm_postgresql_flexible_server`, `azurerm_postgresql_flexible_server_database`, `azurerm_postgresql_server`
- `azurerm_servicebus_namespace`, `azurerm_servicebus_queue`, `azurerm_servicebus_topic`
- `azurerm_storage_account`, `azurerm_storage_blob`, `azurerm_storage_container`, `azurerm_storage_queue`, `azurerm_storage_share`, `azurerm_storage_share_directory`, `azurerm_storage_share_file`, `azurerm_storage_table`

## Example

```hcl
resource "azurerm_storage_account" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
  account_tier        = "Standard"
  account_replication_type = "LRS"
}
```

```
$ tflint
1 issue(s) found:

Warning: Resource is missing lifecycle { prevent_destroy = true }. This resource contains data that should be protected from accidental deletion. (azurerm_resources_missing_prevent_destroy)

  on main.tf line 1:
   1: resource "azurerm_storage_account" "example" {

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.17.1/docs/rules/azurerm_resources_missing_prevent_destroy.md

```

## Why

Data-containing Azure resources like storage accounts, databases, and key vaults should be protected from accidental deletion using Terraform's `prevent_destroy` lifecycle rule. This helps prevent costly data loss incidents during infrastructure changes.

## How To Fix

Add a `lifecycle` block with `prevent_destroy = true` to protect the resource:

```hcl
resource "azurerm_storage_account" "example" {
  name                = "example"
  location            = "West Europe"
  resource_group_name = "example"
  account_tier        = "Standard"
  account_replication_type = "LRS"

  lifecycle {
    prevent_destroy = true
  }
}
```

## Advanced Configuration

### Custom Resource List

You can override the default list of resources to check:

```hcl
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
  resources = ["azurerm_virtual_machine", "azurerm_managed_disk"]
}
```

### Excluding Resources

You can exclude specific resource types from the check:

```hcl
rule "azurerm_resources_missing_prevent_destroy" {
  enabled = true
  exclude = ["azurerm_storage_blob"] # Don't check storage blobs
}
```