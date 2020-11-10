
# azurerm_storage_account_invalid_name

Warns about invalid storage account names.

Allowed values are:
- Mininum length 3.
- Maximum length 24.	
- Lowercase letters and numbers.

Note: Storage account names are globally unique, this rule only validates the name format, but not it's uniqueness. 
## Example

```hcl

resource "azurerm_storage_account" "example" {
  name  = "AA123"
}

```

```
$ tflint
1 issue(s) found:

Error: "AA123" does not match valid pattern ^[a-z0-9]{3,24}$ (azurerm_storage_account_invalid_name)

  on main.tf line 82:
  82:   name                     = "AA123"

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.5.1/docs/rules/azurerm_storage_account_invalid_name.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/resource-name-rules#microsoftstorage
