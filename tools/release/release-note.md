## What's Changed
### Breaking Changes
* fix: repairs the maintenance jobs apispec-rule-gen and api-version-bumper by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/453
  * Removed the following rules
    * azurerm_batch_account_invalid_name
    * azurerm_batch_account_invalid_pool_allocation_mode
    * azurerm_batch_application_invalid_account_name
    * azurerm_batch_application_invalid_name
    * azurerm_batch_certificate_invalid_account_name
    * azurerm_batch_certificate_invalid_format
    * azurerm_batch_pool_invalid_account_name
    * azurerm_batch_pool_invalid_name
    * azurerm_traffic_manager_profile_invalid_profile_status
    * azurerm_traffic_manager_profile_invalid_traffic_routing_method

### Enhancements
* Feat appinsights hidden tags on app services by @pregress in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/430
* feat: add rule azurerm_linux_virtual_machine_invalid_name by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/436
* Feat appservice autoheal by @pregress in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/441
* Feat/run azurerm supported tags updater by @TomBurdge in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/458
* Update virtual machine sizes by @kuglimon in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/461

### Chores
* Bump golang.org/x/oauth2 from 0.32.0 to 0.33.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/435
* docs: add more explanation to azurerm_windows_virtual_machine_invalid_name.md by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/437
* Bump actions/checkout from 5.0.0 to 6.0.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/438
* Bump actions/setup-go from 6.0.0 to 6.1.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/440
* Bump github.com/hashicorp/go-version from 1.7.0 to 1.8.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/442
* Bump peter-evans/create-pull-request from 7.0.8 to 7.0.9 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/439
* Bump peter-evans/create-pull-request from 7.0.9 to 7.0.11 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/443
* Bump golang.org/x/oauth2 from 0.33.0 to 0.34.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/447
* Bump peter-evans/create-pull-request from 7.0.11 to 8.0.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/446
* Bump actions/attest-build-provenance from 3.0.0 to 3.1.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/448
* Bump actions/checkout from 6.0.0 to 6.0.1 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/444
* ci: user_assigned_identity_resource exception is not needed anymore since submodule commit f12a154 by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/450
* chore: updates api-version-bumper/terraform-provider-azurerm to commit f1aec10 by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/451
* Bump actions/setup-go from 6.1.0 to 6.2.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/452
* Bump peter-evans/create-pull-request from 8.0.0 to 8.1.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/455
* Bump actions/attest-build-provenance from 3.1.0 to 3.2.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/457
* Bump golang.org/x/oauth2 from 0.34.0 to 0.35.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/460

## New Contributors
* @TomBurdge made their first contribution in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/458
* @kuglimon made their first contribution in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/461

**Full Changelog**: https://github.com/terraform-linters/tflint-ruleset-azurerm/compare/v0.30.0...v0.31.0
