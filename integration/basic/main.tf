resource "azurerm_linux_virtual_machine" "main" {
  name                = "main-vm"
  vm_size             = "Standard_D2ds_v4"
  location            = "westus2"
  resource_group_name = "example"
}
