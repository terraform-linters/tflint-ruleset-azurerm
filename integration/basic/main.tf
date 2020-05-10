resource "azurerm_virtual_machine" "main" {
  name                = "main-vm"
  vm_size             = "Standard_DS1_v3"
  location            = "West US 2"
  resource_group_name = "example"
}
