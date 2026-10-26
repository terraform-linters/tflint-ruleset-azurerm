# azurerm_resource_provider_not_registered

Require that the Azure resource providers used by azurerm resources are registered by their provider configuration.

## Configuration

```hcl
rule "azurerm_resource_provider_not_registered" {
  enabled = true

  # (Optional) Mode to assume when resource_provider_registrations is not set. Defaults to "none", the azurerm 5.x default.
  # Set this when you configure the mode with the ARM_RESOURCE_PROVIDER_REGISTRATIONS environment variable instead.
  default_registrations = "none"

  # (Optional) Namespaces that are registered outside Terraform, for example centrally by a platform team.
  # Resources that use them are not reported. Matched case-insensitively.
  ignored_namespaces = ["microsoft.insights"]
}
```

## Example

```hcl
provider "azurerm" {
  resource_provider_registrations = "core"
  resource_providers_to_register  = ["Microsoft.Dashboardd"]

  features {}
}

resource "azurerm_dashboard_grafana" "this" {
  grafana_major_version = 11
  location              = "westeurope"
  name                  = "dg-example"
  resource_group_name   = "rg-example"
}
```

```
$ tflint
1 issue(s) found:

Warning: The "Microsoft.Dashboard" resource provider required by "azurerm_dashboard_grafana" is not registered by provider "azurerm" (resource_provider_registrations = "core"). Add it to resource_providers_to_register. (azurerm_resource_provider_not_registered)

  on main.tf line 8:
   8: resource "azurerm_dashboard_grafana" "this" {

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.32.0/docs/rules/azurerm_resource_provider_not_registered.md
```

## Why

Since azurerm 5.0, `resource_provider_registrations` defaults to `none`, and the built-in sets (`core`, `extended`, `all`, `legacy`) are no longer updated. Resources whose resource provider is not registered in the subscription pass `terraform plan`, but fail during `terraform apply` with `MissingSubscriptionRegistration`. A mistyped or wrongly cased namespace in `resource_providers_to_register` is ignored by the provider with only a warning in the debug log.

The rule looks up the namespaces each resource uses, based on the `API Providers` section of the azurerm resource documentation. It reports namespaces that are not registered by the `resource_provider_registrations` mode or `resource_providers_to_register` of the provider configuration the resource uses. Namespaces that Azure registers in every subscription, such as `Microsoft.Resources` and `Microsoft.Authorization`, are ignored.

The `API Providers` sections also list namespaces that a resource only reads from, such as `Microsoft.OperationalInsights` for `azurerm_container_app_environment`, whose Log Analytics workspace may be in another subscription. The rule may therefore report namespaces that the subscription does not strictly need. Use `ignored_namespaces` for those.

The azurerm provider compares namespaces with the exact casing Azure returns, which is not always the documented casing: Azure returns `microsoft.insights`, not `Microsoft.Insights`. Entries in `resource_providers_to_register` with a different casing are skipped by the provider with only a warning, so the rule reports them on the provider configuration and suggests the casing Azure uses. Resources that use such a namespace are not reported again. Check the casing with `az provider show --namespace <namespace> --query namespace`.

Resources in child modules are checked too, as Terraform runs them with the provider configurations of the root module. The modules are read from disk: local modules by their path, and other modules from `.terraform/modules` (or `$TF_DATA_DIR/modules`), so run `terraform init` first. Provider configurations passed with the `providers` argument are followed, and the default `azurerm` configuration is inherited implicitly. Issues for resources in a child module are reported on the `module` block in the root module, with the module address in the message:

```
Warning: The "microsoft.insights" resource provider required by "azurerm_monitor_diagnostic_setting" in module.base.module.vnet is not registered by provider "azurerm" (resource_provider_registrations = "core"). Add it to resource_providers_to_register. (azurerm_resource_provider_not_registered)

  on main.tf line 12:
  12: module "base" {
```

Modules are always read, regardless of the TFLint `call_module_type` setting.

List every namespace the configuration uses, even if another configuration targeting the same subscription already registers it. The azurerm provider skips namespaces that are already registered and never unregisters them, so overlapping lists are harmless. Note that with any mode other than `none`, or any non-empty `resource_providers_to_register`, the provider lists the resource providers of the subscription on every run and registers the missing ones, which requires permission to do so. If registration is managed outside Terraform, keep `none` and add those namespaces to `ignored_namespaces` instead.

The rule does not report anything when:

- the module has no `azurerm` provider configuration, such as a reusable module that is linted on its own;
- `resource_provider_registrations` or `resource_providers_to_register` cannot be evaluated, for example when it refers to a module output;
- the resource type has no `API Providers` section in the azurerm documentation;
- the resource is in a child module that is not installed, has a `source` that is not a static string, or declares its own `azurerm` provider configuration;
- the resource uses a provider configuration that is not passed to its module.

## How To Fix

Add the reported namespaces to `resource_providers_to_register` of the provider configuration the resource uses, with the casing from the message. If a namespace is registered outside Terraform, add it to `ignored_namespaces`.
