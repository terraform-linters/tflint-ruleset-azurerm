## What's Changed

Support for Cosign signatures has been removed from this release. The `checksums.txt.keyless.sig` and `checksums.txt.pem` will not be included in the release.
These files are not used in normal use cases, so in most cases this will not affect you, but if you are affected, you can use Artifact Attestations instead.

### Breaking Changes
* Bump github.com/terraform-linters/tflint-plugin-sdk from 0.22.0 to 0.23.1 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/426
  * Requires TFLint v0.46+
* Bump API versions by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/433
  * Removed the following rules
    * azurerm_netapp_account_invalid_name
    * azurerm_netapp_pool_invalid_account_name
    * azurerm_netapp_pool_invalid_name
    * azurerm_netapp_snapshot_invalid_account_name
    * azurerm_netapp_snapshot_invalid_pool_name
    * azurerm_netapp_snapshot_invalid_volume_name
    * azurerm_netapp_volume_invalid_account_name
    * azurerm_netapp_volume_invalid_name
    * azurerm_netapp_volume_invalid_pool_name

### Enhancements
* feat(vm_size): update vm sizes to state at 2025-08-09 by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/408
* feat: prevent destroy on resources that contain data to prevent accidantal data loss by @pregress in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/427
  * Added a new rule: `azurerm_resources_missing_prevent_destroy`
* feat(vm_size): Updates vm sizes to state at 2025-11-05 by @JannoTjarks in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/429

### Chores
* Bump actions/checkout from 4.2.2 to 5.0.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/409
* Bump github.com/zclconf/go-cty from 1.16.3 to 1.16.4 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/411
* Bump github.com/zclconf/go-cty from 1.16.3 to 1.16.4 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/412
* dependabot: allow actions writes by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/413
* Bump goreleaser/goreleaser-action from 6.3.0 to 6.4.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/410
* Fix E2E tests to take into account the newly added JSON fields by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/421
* Bump actions/attest-build-provenance from 2.4.0 to 3.0.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/414
* Bump github.com/zclconf/go-cty from 1.16.4 to 1.17.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/416
* Bump golang.org/x/oauth2 from 0.30.0 to 0.31.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/417
* Bump github.com/zclconf/go-cty from 1.16.4 to 1.17.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/418
* Bump actions/setup-go from 5.5.0 to 6.0.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/419
* Bump sigstore/cosign-installer from 3.9.2 to 3.10.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/422
* Bump golang.org/x/oauth2 from 0.31.0 to 0.32.0 in /tools by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/424
* Bump sigstore/cosign-installer from 3.10.0 to 4.0.0 by @dependabot[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/425
* Bump Go version to 1.25 by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/432
* 🤖 MicrosoftDocs/azure-docs changes by @github-actions[bot] in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/420
* Drop support for Cosign signatures by @wata727 in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/434

## New Contributors
* @pregress made their first contribution in https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/427

**Full Changelog**: https://github.com/terraform-linters/tflint-ruleset-azurerm/compare/v0.29.0...v0.30.0
