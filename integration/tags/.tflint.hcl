plugin "azurerm" {
  enabled = true
}

rule "azurerm_resource_missing_tags" {
  enabled = true
  exclude = ["azurerm_virtual_machine"]
  tags = ["Environment", "Department"]
}