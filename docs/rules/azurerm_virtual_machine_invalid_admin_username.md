# azurerm_linux_virtual_machine_invalid_size

Warns about admin user values that appear to be invalid based on https://docs.microsoft.com/en-us/azure/virtual-machines/linux/faq#what-are-the-username-requirements-when-creating-a-vm

Invalid Admin Usernames include:
- azureuser
- 1
- 123
- a
- actuser
- adm
- admin
- admin1
- admin2
- administrator
- aspnet
- backup
- console
- david
- guest
- john
- owner
- root
- server
- sql
- support_388945a0
- support
- sys
- test
- test1
- test2
- test3
- user
- user1
- user2
- user3
- user4
- user5
- video

Username validation is lower cased to ensure that linting rules are applied regardless of casing for the admin_username field.

Note, the value "AzureUser" will result in an linting error even though is not explicitly mentioned in the Virtual Machine docs. This is because of it's common and default usage, it becomes an easy to guess username. It's recommended to select a unique user name.

## Example

```hcl
resource "azurerm_virtual_machine" "test" {
	os_profile {
		admin_username = "123"
	}
}
```

```
$ tflint
1 issue(s) found:

Error: "123" is not a valid VM Admin username (azurerm_virtual_machine_invalid_admin_username)

  on main.tf line 75:
  75:     admin_username = "123"

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.5.1/docs/rules/azurerm_virtual_machine_invalid_admin_username.md

```

## Why

Requests containing not allowed usernames will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the not allowed usernames with a valid value that is not in the list above.

## Source

https://docs.microsoft.com/en-us/azure/virtual-machines/linux/faq#what-are-the-username-requirements-when-creating-a-vm
