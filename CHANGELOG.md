## 0.8.1 (2021-02-01)

- [#94](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/94): Remove rules which checks about overflow integer

## 0.8.0 (2021-01-31)

The minimum supported version of TFLint has changed in this version. TFLint v0.24.0+ is required for this plugin to work.

### Breaking Changes

- [#93](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/93): Bump tflint-plugin-sdk to v0.8.0

### Enhancements

- [#92](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/92): Bump terraform-provider-azurerm to v2.45.1 from v2.41.0

### Chores

- [#89](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/89): Bump github.com/hashicorp/hcl/v2 from 2.8.1 to 2.8.2 in /tools
- [#91](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/91): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.7.0 to 0.7.1

## 0.7.0 (2021-01-04)

The minimum supported version of TFLint has changed in this version. TFLint v0.23.0+ is required for this plugin to work.

### Breaking Changes

- [#88](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/88): Bump tflint-plugin-sdk to v0.7.0

### Enhancements

- [#87](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/87): Bump terraform-provider-azurerm to v2.41.0 from v2.37.0

### Chores

- [#80](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/80): Bump github.com/google/go-cmp from 0.5.3 to 0.5.4
- [#85](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/85): Bump github.com/hashicorp/hcl/v2 from 2.7.1 to 2.8.1 in /tools
- [#86](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/86): Bump github.com/hashicorp/hcl/v2 from 2.7.1 to 2.8.1

## 0.6.0 (2020-11-24)

The minimum supported version of TFLint has changed in this version. TFLint v0.21.0+ is required for this plugin to work.

### Breaking Changes

- [#79](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/79): Bump tflint-plugin-sdk to v0.6.0
  - Added support for JSON configuration syntax

### Enhancements

- [#78](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/78): Bump terraform-provider-azurerm to v2.37.0 from v2.32.0

### Chores

- [#68](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/68): Added input parameters for generation paths. Also debugging configs for VS Code
- [#71](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/71): Refactor of Mapping Files
- [#72](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/72): Add a make command to update the api spec submodule
- [#74](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/74): Bump github.com/google/go-cmp from 0.5.2 to 0.5.3
- [#75](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/75): Stop using set-env commands
- [#76](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/76): Bump github.com/hashicorp/hcl/v2 from 2.6.0 to 2.7.1 in /tools
- [#77](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/77): Bump github.com/hashicorp/hcl/v2 from 2.6.0 to 2.7.1

## 0.5.1 (2020-10-19)

### Enhancements

- [#65](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/65): Bump terraform-provider-azurerm to v2.32.0 from v2.27.0

### Chores

- [#63](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/63): Bump actions/setup-go from v2.1.2 to v2.1.3

## 0.5.0 (2020-09-14)

The minimum supported version of TFLint has changed in this version. TFLint v0.20.0+ is required for this plugin to work. Also, this release is built with Go v1.15. As a result, darwin/386 build will no longer available from the release.

### Breaking Changes

- [#61](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/61): Bump tflint-plugin-sdk to v0.5.0

### Chores

- [#53](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/53): chore(deps): bump go to v1.15
- [#54](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/54): Update GitHub Actions by Dependabot
- [#55](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/55): Bump actions/setup-go from v1 to v2.1.2
- [#56](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/56): Bump actions/github-script from v2 to v3
- [#57](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/57): Bump github.com/google/go-cmp from 0.5.1 to 0.5.2
- [#59](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/59): Update doc snapshot
- [#60](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/60): Bump terraform-provider-azurerm to v2.27.0 from v2.23.0
- [#62](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/62): Bump goreleaser version

## 0.4.0 (2020-08-17)

The minimum supported version of TFLint has changed in this version. TFLint v0.19.0+ is required for this plugin to work.

### Breaking Changes

- [#50](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/50): Bump tflint-plugin-sdk to v0.4.0

### Enhancements

- [#47](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/47): Update VM machine sizes
- [#48](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/48): Bump terraform-provider-azurerm to v2.23.0 from v2.19.0

### Chores

- [#43](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/43): Update azure-docs snapshots
- [#46](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/46): Bump github.com/google/go-cmp from 0.5.0 to 0.5.1

## 0.3.0 (2020-07-19)

The minimum supported version of TFLint has changed in this version. TFLint v0.18.0+ is required for this plugin to work.

### Breaking Changes

- [#41](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/41): Bump tflint-plugin-sdk to v0.3.0

### Enhancements

- [#31](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/31) [#40](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/40): Bump terraform-provider-azurerm to v2.19.0 from v2.16.0

### BugFixes

- [#32](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/32) [#36](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/36): Add newly added vm sizes to rules

### Chores

- [#33](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/33): Add workflow to watch azure-docs content changes

## 0.2.0 (2020-06-28)

The minimum supported version of TFLint has changed in this version. TFLint v0.17.0+ is required for this plugin to work.

### Breaking Changes

- [#20](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/20) [#25](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/25): Bump tflint-plugin-sdk from v0.1.0 to v0.2.0

### Chores

- [#21](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/21): Bump github.com/hashicorp/hcl/v2 from 2.5.1 to 2.6.0 in /tools
- [#22](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/22): Bump github.com/hashicorp/hcl/v2 from 2.5.1 to 2.6.0
- [#23](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/23): Bump github.com/google/go-cmp from 0.4.1 to 0.5.0
- [#24](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/24): Create Dependabot config file
- [#27](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/27): Bump terraform-provider-azurerm to v2.16.0 from v2.10.0

## 0.1.0 (2020-05-18)

Initial release 🎉
