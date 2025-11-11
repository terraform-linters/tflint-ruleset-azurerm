# azurerm_windows_virtual_machine_invalid_name

Warns about values that appear to be invalid based on [Naming rules and restrictions for Azure resources](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-name-rules#microsoftcompute).

In this rule, the string must match the regular expression `^[a-zA-Z0-9]{0,1}[a-zA-Z0-9-]{0,13}[a-zA-Z0-9]$`.

## Example

```hcl
resource "azurerm_windows_virtual_machine" "foo" {
  name = "dummy-" // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "dummy-" does not match valid pattern ^[a-zA-Z0-9]{0,1}[a-zA-Z0-9-]{0,13}[a-zA-Z0-9]$ (azurerm_windows_virtual_machine_invalid_name)

  on template.tf line 2:
  2:   name = "dummy-" // invalid value

```

## Why

There are hard limitations regarding the Naming of a Windows Server Hostname. E.g.: Windows does not allow any hostname which is longer than 15 characters.
Requests containing invalid values will return an error when calling the API by `terraform apply`.

More details are available in the official Microsoft documentation: https://learn.microsoft.com/de-de/troubleshoot/windows-server/active-directory/naming-conventions-for-computer-domain-site-ou#dns-host-names

## How To Fix

Replace the warned value with a valid value. In this example a valid name could be "dummy-1", because it matches the valid pattern ^[a-zA-Z0-9]{0,1}[a-zA-Z0-9-]{0,13}[a-zA-Z0-9]$.
