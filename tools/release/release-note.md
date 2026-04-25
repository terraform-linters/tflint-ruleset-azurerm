## What's Changed

### Breaking Changes
* Fix maintenance job failed by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/478
  * Removed the following rules
    * azurerm_container_registry_invalid_name
    * azurerm_container_registry_webhook_invalid_name
    * azurerm_container_registry_webhook_invalid_registry_name
    * azurerm_data_factory_dataset_mysql_invalid_linked_service_name
    * azurerm_data_factory_dataset_mysql_invalid_name
    * azurerm_data_factory_dataset_postgresql_invalid_linked_service_name
    * azurerm_data_factory_dataset_postgresql_invalid_name
    * azurerm_data_factory_dataset_sql_server_table_invalid_linked_service_name
    * azurerm_data_factory_dataset_sql_server_table_invalid_name
    * azurerm_data_factory_invalid_name
    * azurerm_data_factory_invalid_resource_group_name
    * azurerm_data_factory_linked_service_data_lake_storage_gen2_invalid_name
    * azurerm_data_factory_linked_service_mysql_invalid_name
    * azurerm_data_factory_linked_service_postgresql_invalid_name
    * azurerm_data_factory_linked_service_sql_server_invalid_name
    * azurerm_data_factory_pipeline_invalid_name
    * azurerm_kubernetes_cluster_invalid_name

### Enhancements
* feat(vm_size): Updates vm sizes to state at 2026-03-26 by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/480
* Added Retired Virtual Machine Size Rules for Linux and Windows VMs by @HenryGelderbloem in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/487

### Chores
* build(deps): bump github.com/zclconf/go-cty from 1.17.0 to 1.18.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/470
* build(deps): bump github.com/zclconf/go-cty from 1.17.0 to 1.18.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/471
* build(deps): bump actions/setup-go from 6.2.0 to 6.3.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/472
* build(deps): bump hashicorp/setup-terraform from 3.1.2 to 4.0.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/474
* build(deps): bump golang.org/x/oauth2 from 0.35.0 to 0.36.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/475
* deps: Bump Go version to 1.26 by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/476
* build(deps): bump actions/attest-build-provenance from 3.2.0 to 4.1.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/473
* Bump actions/checkout from 6.0.1 to 6.0.2 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/456
* chore(deps): bump google.golang.org/grpc from 1.75.1 to 1.79.3 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/477
* chore(deps): bump github.com/terraform-linters/tflint-plugin-sdk from 0.23.1 to 0.24.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/479
* dependabot: Set cooldown period by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/481
* chore(deps): bump actions/setup-go from 6.3.0 to 6.4.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/482
* chore(deps): bump github.com/hashicorp/go-version from 1.8.0 to 1.9.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/483
* chore(deps): bump peter-evans/create-pull-request from 8.1.0 to 8.1.1 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/488
* fix(rule generator): handle also CRLF instead of just LF by @HenryGelderbloem in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/486
* chore(deps): bump github.com/zclconf/go-cty from 1.18.0 to 1.18.1 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/492
* chore(deps): bump goreleaser/goreleaser-action from 7.0.0 to 7.1.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/493
* chore(deps): bump github.com/dave/dst from 0.27.3 to 0.27.4 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/495
* chore(deps): bump github.com/zclconf/go-cty from 1.18.0 to 1.18.1 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/494
* release: Migrate attest-build-provenance to attest by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/496

## New Contributors
* @HenryGelderbloem made their first contribution in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/486

**Full Changelog**: https://github.com/terraform-linters/tflint-ruleset-azurerm/compare/v0.31.1...v0.32.0
