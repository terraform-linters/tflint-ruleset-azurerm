
# azurerm_subnet_invalid_address_prefixes

Warns about invalid CIDR in address_prefixes.

Every item in the address_prefixes list must be a [valid CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

## Example

```hcl

resource "azurerm_subnet" "test" {
	address_prefixes     = ["invalid_cidr"]
}

```

```
$ tflint
1 issue(s) found:

Error: "invalid_cidr" is not a valid CIDR (azurerm_subnet_invalid_address_prefixes)

  on main.tf line 86:
  86:   address_prefixes     = ["invalid_cidr"]

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.5.1/docs/rules/azurerm_subnet_invalid_address_prefixes.md
```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

* https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
* https://docs.microsoft.com/en-us/azure/virtual-network/virtual-networks-faq
