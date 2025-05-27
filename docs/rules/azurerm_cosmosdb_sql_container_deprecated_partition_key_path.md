
# azurerm_cosmosdb_sql_container_deprecated_partition_key_path

### Description

This rule checks for the use of the deprecated `partition_key_path` attribute in `azurerm_cosmosdb_sql_container` resources. The `partition_key_path` attribute is deprecated and should be replaced with the `partition_key_paths` attribute, which supports a list of partition key paths. This update ensures compatibility with newer versions of the AzureRM provider and supports Cosmos DB's enhanced partitioning capabilities.

---

## Example

```hcl
resource "azurerm_cosmosdb_sql_container" "example" {
  partition_key_path = "/deprecated_key"
}
```

### Output when inspected by TFLint:

```bash
$ tflint
resource.tf:3:24 - Warning: `partition_key_path` is deprecated and should be replaced with `partition_key_paths`.
```

---

## Why

Using deprecated attributes like `partition_key_path` can lead to future compatibility issues with the AzureRM provider as it progresses to newer versions. The `partition_key_paths` attribute offers more flexibility by supporting multiple partition key paths, which is crucial for modern Cosmos DB designs that rely on enhanced partitioning strategies.

---

## How to Fix

Replace the deprecated `partition_key_path` attribute with the `partition_key_paths` attribute. Update your resource configuration as follows:

```hcl
resource "azurerm_cosmosdb_sql_container" "example" {
  partition_key_paths = ["/valid_key"]
}
```

This ensures your Terraform code is compliant with the latest version of the AzureRM provider and avoids potential errors when upgrading or using newer features of Cosmos DB.
