# azurerm_linux_virtual_machine_invalid_name

Warns about values that appear to be invalid based on [Naming rules and restrictions for Azure resources](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-name-rules#microsoftcompute) for a Linux VM.

In this rule, the string must match the regular expression `^[a-zA-Z0-9][a-zA-Z0-9-]{0,62}$`.

## Example

```hcl
resource "azurerm_linux_virtual_machine" "foo" {
  name = "-dummy" // invalid value
}
```

```
$ tflint


1 issue(s) found:

Error: "-dummy" does not match valid pattern ^[a-zA-Z0-9][a-zA-Z0-9-]{0,62}$ (azurerm_linux_virtual_machine_invalid_name)

  on template.tf line 2:
  2:   name = "-dummy" // invalid value

```

## Why

There are hard limitations regarding the Naming of a Linux Server Hostname. E.g.: Linux does not allow any hostname which is longer than 63 characters. Requests containing invalid values will return an error when calling the API by terraform apply.

More details are available here: https://man7.org/linux/man-pages/man7/hostname.7.html

## How To Fix

Replace the warned value with a valid value. In this example a valid name could be "dummy", because it matches the valid pattern ^[a-zA-Z0-9][a-zA-Z0-9-]{0,62}$.
