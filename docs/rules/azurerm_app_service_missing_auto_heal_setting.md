# azurerm_app_service_missing_auto_heal_setting

Disallow missing auto_heal_setting configuration in site_config block for Azure App Service resources.

This rule applies to the following Azure App Service resource types:
- `azurerm_linux_web_app`
- `azurerm_linux_web_app_slot`
- `azurerm_windows_web_app`
- `azurerm_windows_web_app_slot`

## Configuration

```hcl
rule "azurerm_app_service_missing_auto_heal_setting" {
  enabled = true
}
```

## Example

### Terraform Configuration

```hcl
# Non-compliant: Linux Web App with site_config but no auto_heal_setting
resource "azurerm_linux_web_app" "example" {
  name                = "example-app"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  service_plan_id     = azurerm_service_plan.example.id

  site_config {
    always_on = true
  }
}

# Compliant: Linux Web App with auto_heal_setting configured
resource "azurerm_linux_web_app" "example" {
  name                = "example-app"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  service_plan_id     = azurerm_service_plan.example.id

  site_config {
    always_on = true
    
    auto_heal_setting {
      action {
        action_type = "Recycle"
      }
      trigger {
        status_code {
          count             = 5
          interval          = "00:01:00"
          status_code_range = "500-599"
        }
      }
    }
  }
}

# Compliant: Windows Web App Slot with comprehensive auto_heal_setting
resource "azurerm_windows_web_app_slot" "example" {
  name           = "example-slot"
  app_service_id = azurerm_windows_web_app.example.id

  site_config {
    auto_heal_setting {
      action {
        action_type                    = "Recycle"
        minimum_process_execution_time = "00:01:00"
      }
      trigger {
        status_code {
          count             = 5
          interval          = "00:01:00"
          status_code_range = "500-599"
        }
        requests {
          count    = 100
          interval = "00:01:00"
        }
        slow_request {
          count      = 10
          interval   = "00:02:00"
          time_taken = "00:00:45"
        }
      }
    }
  }
}
```

## Why

Configuring `auto_heal_setting` in the `site_config` block is a best practice for production Azure App Service resources. Auto-healing helps improve application resilience by automatically recycling or restarting the app when specific conditions are met, such as:

- High number of HTTP errors (status codes in the 500-599 range)
- Excessive request volume
- Slow response times
- Memory threshold breaches

By proactively detecting and responding to unhealthy states, auto-healing can prevent prolonged outages and improve overall application availability. This rule ensures that App Service resources have auto-healing configured to maintain production resilience.

For more information about building robust apps for the cloud with auto-heal, see the [Azure App Service documentation on Auto Heal](https://azure.github.io/AppService/2020/05/15/Robust-Apps-for-the-cloud.html#auto-heal).

## How to Fix

Add an `auto_heal_setting` block within the `site_config` block:

```hcl
site_config {
  auto_heal_setting {
    action {
      action_type = "Recycle"
    }
    trigger {
      status_code {
        count             = 5
        interval          = "00:01:00"
        status_code_range = "500-599"
      }
    }
  }
}
```
You can combine multiple triggers to create comprehensive auto-healing policies.
