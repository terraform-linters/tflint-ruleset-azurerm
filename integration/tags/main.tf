resource "azurerm_virtual_machine" "main" {
  name                = "main-vm"
  location            = "West US 2"
  resource_group_name = "example"
}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "West Europe"
  tags = { foo = "bar" }
}

resource "azurerm_service_plan" "example" {
  name                = "example-appserviceplan"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  os_type             = "Linux"
  sku_name            = "S1"

  tags = {
    Environment = "Production"
  }
}

resource "azurerm_linux_web_app" "example" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  service_plan_id     = azurerm_service_plan.example.id

  site_config {
    auto_heal_setting {
      action {
        action_type = "Recycle"
      }
      trigger {
        requests {
          count    = 100
          interval = "00:01:00"
        }
      }
    }
  }

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }

  tags = {
    Environment = "Production",
    Department = "IT"
  }
}
