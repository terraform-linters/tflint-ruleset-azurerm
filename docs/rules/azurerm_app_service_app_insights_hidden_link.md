# azurerm_app_service_app_insights_hidden_link

Disallow missing lifecycle ignore_changes for Application Insights hidden-link tags.

This rule applies to all Azure App Service resource types (Web Apps and Function Apps):
- `azurerm_linux_web_app`
- `azurerm_linux_web_app_slot`
- `azurerm_windows_web_app`
- `azurerm_windows_web_app_slot`
- `azurerm_linux_function_app`
- `azurerm_linux_function_app_slot`
- `azurerm_windows_function_app`
- `azurerm_windows_function_app_slot`

## Configuration

```hcl
rule "azurerm_app_service_app_insights_hidden_link" {
  enabled = true
}
```

## Example

### Terraform Configuration

```hcl
# Non-compliant: Linux Web App with Application Insights configured without ignore_changes
resource "azurerm_linux_web_app" "example" {
  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
}

# Non-compliant: Windows Web App Slot with Application Insights configured without ignore_changes
resource "azurerm_windows_web_app_slot" "example" {
  name           = "example-slot"
  app_service_id = azurerm_windows_web_app.example.id

  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example-key"
  }
}

# Non-compliant: Windows Function App with Application Insights configured without ignore_changes
resource "azurerm_windows_function_app" "example" {
  name                = "example-func"
  resource_group_name = "example-rg"
  location            = "West Europe"
  service_plan_id     = azurerm_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example-key"
  }
}

# Compliant: Linux Function App with proper ignore_changes
resource "azurerm_linux_function_app" "example" {
  name                = "example-func"
  resource_group_name = "example-rg"
  location            = "West Europe"
  service_plan_id     = azurerm_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

  app_settings = {
    "APPLICATIONINSIGHTS_CONNECTION_STRING" = "example"
  }
  
  lifecycle {
    ignore_changes = [
      "tags[\"hidden-link: /app-insights-conn-string\"]",
      "tags[\"hidden-link: /app-insights-instrumentation-key\"]",
      "tags[\"hidden-link: /app-insights-resource-id\"]",
    ]
  }
}

# Compliant: Windows Web App Slot with proper ignore_changes
resource "azurerm_windows_web_app_slot" "example" {
  name           = "example-slot"
  app_service_id = azurerm_windows_web_app.example.id

  app_settings = {
    "APPINSIGHTS_INSTRUMENTATIONKEY" = "example"
  }
  
  lifecycle {
    ignore_changes = [
      "tags[\"hidden-link: /app-insights-conn-string\"]",
      "tags[\"hidden-link: /app-insights-instrumentation-key\"]",
      "tags[\"hidden-link: /app-insights-resource-id\"]",
    ]
  }
}
```

## Why

When Application Insights is configured for Azure App Service resources (Web Apps, Web App Slots, Function Apps, or Function App Slots for both Linux and Windows), Azure automatically adds hidden-link tags to the resource. These tags can change during deployments and may cause unnecessary Terraform diffs. Using `lifecycle { ignore_changes }` for these specific tags prevents Terraform from attempting to manage these Azure-managed tags.

This rule checks for the presence of Application Insights configuration (either `APPLICATIONINSIGHTS_CONNECTION_STRING` or `APPINSIGHTS_INSTRUMENTATIONKEY` in app_settings) and ensures that the corresponding hidden-link tags are properly ignored.

## How to Fix

Add a `lifecycle` block with `ignore_changes` containing the Application Insights hidden-link tags:

```hcl
lifecycle {
  ignore_changes = [
    "tags[\"hidden-link: /app-insights-conn-string\"]",
    "tags[\"hidden-link: /app-insights-instrumentation-key\"]", 
    "tags[\"hidden-link: /app-insights-resource-id\"]",
  ]
}
```