## 0.23.0 (2023-04-23)

### Enhancements

- [#265](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/265): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.15.0 to 0.16.1

### Chores

- [#262](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/262) [#267](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/267): Bump terraform-provider-azurerm to v3.53.0 from v3.48.0
- [#264](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/264): Bump peter-evans/create-pull-request from 4 to 5
- [#268](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/268): Follow up of the EnsureNoError deprecation

## 0.22.0 (2023-03-21)

### Enhancements

- [#258](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/258): Add support for HBv4 and HX series

### BugFixes

- [#252](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/252): Fix `azurerm_postgresql_server` rule by correcting import path

### Chores

- [#238](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/238): Bump golang.org/x/net from 0.3.0 to 0.7.0
- [#237](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/237) [#239](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/239) [#257](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/257) [#259](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/259): Bump terraform-provider-azurerm to v3.48.0 from v3.43.0
- [#242](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/242): Bump golang.org/x/text from 0.3.7 to 0.3.8 in /tools
- [#245](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/245): Bump sigstore/cosign-installer from 2 to 3
- [#247](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/247): Bump github.com/hashicorp/hcl/v2 from 2.16.0 to 2.16.2 in /tools
- [#248](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/248): Bump github.com/hashicorp/hcl/v2 from 2.16.0 to 2.16.2
- [#254](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/254): Bump github.com/zclconf/go-cty from 1.12.1 to 1.13.1 in /tools
- [#255](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/255): Bump github.com/zclconf/go-cty from 1.12.1 to 1.13.1
- [#256](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/256): Bump actions/setup-go from 3 to 4
- [#260](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/260): Fix signing for Cosign v2
- [#261](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/261): go 1.20

## 0.21.0 (2023-02-12)

### Breaking Changes

- [#235](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/235): Remove `azurerm_container_group_invalid_ip_address_type` rule

### Enhancements

- [#234](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/234): Bump terraform-provider-azurerm to v3.43.0 from v3.35.0

### Chores

- [#226](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/226): Bump goreleaser/goreleaser-action from 3 to 4
- [#227](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/227): Pass GITHUB_TOKEN to e2e test workflow
- [#228](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/228): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.14.0 to 0.15.0
- [#231](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/231): Bump github.com/hashicorp/hcl/v2 from 2.15.0 to 2.16.0
- [#232](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/232): Bump github.com/hashicorp/hcl/v2 from 2.15.0 to 2.16.0 in /tools
- [#236](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/236): Update snapshot

## 0.20.0 (2022-12-12)

### Breaking Changes

- [#213](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/213) [#223](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/223) [#224](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/224): Bump terraform-provider-azurerm to v3.35.0 from v3.28.0
  - Removed `azurerm_linux_virtual_machine_scale_set_invalid_upgrade_mode` rule
  - Removed `azurerm_virtual_machine_scale_set_invalid_upgrade_policy_mode` rule
  - Removed `azurerm_windows_virtual_machine_scale_set_invalid_upgrade_mode` rule

### Chores

- [#217](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/217): Add signatures for keyless signing
- [#219](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/219): Bump github.com/hashicorp/hcl/v2 from 2.14.1 to 2.15.0
- [#220](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/220): Bump github.com/hashicorp/hcl/v2 from 2.14.1 to 2.15.0 in /tools

## 0.19.0 (2022-10-27)

### Breaking Changes

- [#202](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/202) [#214](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/214): Bump terraform-provider-azurerm to v3.28.0 from v3.21.1
  - Removed `azurerm_managed_disk_invalid_create_option` and `azurerm_snapshot_invalid_create_option` rules

### Enhancements

- [#210](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/210): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.12.0 to 0.14.0

### Chores

- [#205](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/205): Bump github.com/hashicorp/hcl/v2 from 2.14.0 to 2.14.1 in /tools
- [#208](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/208): Bump github.com/dave/dst from 0.27.0 to 0.27.2
- [#211](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/211): Bump github.com/zclconf/go-cty from 1.11.0 to 1.11.1
- [#212](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/212): Bump github.com/zclconf/go-cty from 1.11.0 to 1.11.1 in /tools

## 0.18.0 (2022-09-08)

The minimum supported version of TFLint has changed in this version. TFLint v0.40.0+ is required for this plugin to work.

### Breaking Changes

- [#198](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/198): Bump tflint-plugin-sdk to v0.12.0

### Enhancements

- [#186](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/186) [#187](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/187): autogenerated maintenace
- [#194](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/194): Adding rule for azurerm resource missing tags
- [#199](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/199): Add support for Dpdsv5/Dpldsv5/Dpsv5/Dplsv5/Epdsv5/Epsv5 series
- [#200](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/200): Bump terraform-provider-azurerm to v3.21.1 from v3.16.0
- [#201](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/201): Update outdated VM sizes for Kubernetes cluster node pool

### Chores

- [#185](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/185): go 1.19
- [#191](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/191): Bump github.com/zclconf/go-cty from 1.10.0 to 1.11.0 in /tools
- [#195](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/195): Bump github.com/hashicorp/hcl/v2 from 2.13.0 to 2.14.0 in /tools
- [#196](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/196): Bump github.com/hashicorp/hcl/v2 from 2.13.0 to 2.14.0
- [#197](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/197): build: Use `go-version-file` instead of `go-version`

## 0.17.1 (2022-07-31)

### Enhancements

- [#180](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/180): autogenerated maintenance
- [#184](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/184): Bump terraform-provider-azurerm to v3.16.0 from v3.10.0

### Chores

- [#182](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/182): Bump github.com/hashicorp/hcl/v2 from 2.12.0 to 2.13.0
- [#183](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/183): Bump github.com/hashicorp/hcl/v2 from 2.12.0 to 2.13.0 in /tools

## 0.17.0 (2022-06-18)

### Enhancements

- [#174](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/174): autogenerated maintenance
- [#176](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/176): Add support for Lsv3/Lasv3 series
- [#178](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/178): Bump terraform-provider-azurerm to v3.10.0 from v3.4.0

### BugFixes

- [#179](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/179): Accept 8.0 as a correct version on MySQL server

### Chores

- [#170](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/170): Bump goreleaser/goreleaser-action from 2 to 3
- [#172](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/172): automating maintenance with Github actions
- [#175](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/175): Bump github.com/dave/dst from 0.26.2 to 0.27.0
- [#177](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/177): Tweak autogenerated maintenance frequency

## 0.16.0 (2022-05-05)

### Enhancements

- [#168](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/168): Bump terraform-provider-azurerm to v3.4.0 from v3.0.2
- [#169](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/169): Add support for Ebdsv5/Ebsv5 series

### Chores

- [#159](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/159): chores: Remove snaker
- [#160](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/160): Fix rule template for rule generator
- [#163](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/163): Bump actions/setup-go from 2 to 3
- [#164](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/164): Bump github.com/hashicorp/hcl/v2 from 2.11.1 to 2.12.0
- [#165](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/165): Bump github.com/hashicorp/hcl/v2 from 2.11.1 to 2.12.0 in /tools
- [#166](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/166): Bump github.com/google/go-cmp from 0.5.7 to 0.5.8
- [#167](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/167): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.10.0 to 0.11.0

## 0.15.0 (2022-03-30)

The minimum supported version of TFLint has changed in this version. TFLint v0.35.0+ is required for this plugin to work.

### Breaking Changes

- [#150](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/150) [#156](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/156) [#158](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/158): Bump terraform-provider-azurerm to v3.0.2 from v2.84.0
- [#155](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/155): Bump tflint-plugin-sdk for gRPC-based new plugin system

### Chores

- [#146](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/146): Bump github.com/zclconf/go-cty from 1.9.1 to 1.10.0 in /tools
- [#148](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/148): Bump github.com/hashicorp/hcl/v2 from 2.10.1 to 2.11.1 in /tools
- [#149](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/149): Bump github.com/hashicorp/hcl/v2 from 2.10.1 to 2.11.1
- [#151](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/151): Bump github.com/google/go-cmp from 0.5.6 to 0.5.7
- [#152](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/152): Bump actions/github-script from 5 to 6
- [#153](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/153): Bump actions/checkout from 2 to 3
- [#154](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/154): build: go 1.18

## 0.14.0 (2021-11-07)

### Breaking Changes

- [#144](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/144): Bump terraform-provider-azurerm to v2.84.0 from v2.80.0
  - Remove `azurerm_mysql_server_invalid_ssl_enforcement` rule
- [#145](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/145): build: Remove unsupported build targets

### Enhancements

- [#143](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/143): Update machine sizes

### Chores

- [#141](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/141): Fix github-script to v5

## 0.13.2 (2021-10-12)

### Enhancements

- [#140](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/140): Bump terraform-provider-azurerm to v2.80.0 from v2.76.0

### Chores

- [#139](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/139): Bump actions/github-script from 4 to 5

## 0.13.1 (2021-09-12)

### BugFixes

- [#137](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/137): build: Update GoReleaser version
  - v0.13.0 release doesn't include darwin/arm64 build. This change fixes the issue.

## 0.13.0 (2021-09-12)

### Enhancements

- [#136](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/136): Bump terraform-provider-azurerm to v2.76.0 from v2.71.0

### Chores

- [#130](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/130): Bump actions/github-script from 4.0.2 to 4.1
- [#131](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/131): Bump github.com/zclconf/go-cty from 1.9.0 to 1.9.1 in /tools
- [#132](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/132): Bump actions/setup-go from 2.1.3 to 2.1.4
- [#134](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/134): Update doc snapshot
- [#135](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/135): build: Go 1.17

## 0.12.0 (2021-08-08)

### Enhancements

- [#122](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/122): Add `azurerm_kubernetes_cluster_default_node_pool_invalid_vm_size` rule
- [#128](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/128): Bump terraform-provider-azurerm to v2.71.0 from v2.66.0

### Chores

- [#123](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/123): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.9.0 to 0.9.1
- [#124](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/124): Add rule generator
- [#125](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/125): Bump github.com/zclconf/go-cty from 1.8.0 to 1.9.0 in /tools
- [#126](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/126): Bump github.com/hashicorp/hcl/v2 from 2.10.0 to 2.10.1 in /tools
- [#127](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/127): Bump github.com/hashicorp/hcl/v2 from 2.10.0 to 2.10.1

## 0.11.0 (2021-07-05)

The minimum supported version of TFLint has changed in this version. TFLint v0.30.0+ is required for this plugin to work.

### Breaking Changes

- [#118](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/118): Bump tflint-plugin-sdk to v0.9.0

### Enhancements

- [#119](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/119): rule: Add support for FX series and update other sizes
- [#120](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/120): Bump terraform-provider-azurerm to v2.66.0 from v2.62.0

## 0.10.1 (2021-06-12)

### Chores

- [#116](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/116): build: Add support for darwin/arm64 build

## 0.10.0 (2021-06-05)

### Enhancements

- [#114](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/114): Bump terraform-provider-azurerm to v2.62.0 from v2.57.0

### Chores

- [#113](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/113): Bump github.com/google/go-cmp from 0.5.5 to 0.5.6
- [#115](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/115): Add notes about auto installation

## 0.9.1 (2021-05-02)

### Enhancements

- [#105](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/105) [#112](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/112): Bump terraform-provider-azurerm to v2.57.0 from v2.50.0
- [#111](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/111): rules: Add support for HBv3 series and update other sizes

### Chores

- [#100](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/100): Bump github.com/google/go-cmp from 0.5.4 to 0.5.5
- [#102](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/102) [#108](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/108): Bump github.com/hashicorp/hcl/v2 from 2.9.0 to 2.10.0 in /tools
- [#103](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/103) [#109](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/109): Bump github.com/hashicorp/hcl/v2 from 2.9.0 to 2.10.0
- [#106](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/106): Bump github.com/terraform-linters/tflint-plugin-sdk from 0.8.1 to 0.8.2
- [#110](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/110): Bump actions/github-script from v3 to v4.0.2

## 0.9.0 (2021-03-07)

### Breaking Changes

- [#99](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/99): Bump terraform-provider-azurerm to v2.50.0 from v2.45.1
  - Remove `azurerm_log_analytics_workspace_invalid_retention_in_days` rule

### Chores

- [#96](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/96): Upgrade to Go 1.16
- [#97](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/97): Bump github.com/hashicorp/hcl/v2 from 2.8.2 to 2.9.0 in /tools
- [#98](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/98): Bump github.com/hashicorp/hcl/v2 from 2.8.2 to 2.9.0

## 0.8.2 (2021-02-02)

- [#95](https://github.com/terraform-linters/tflint-ruleset-azurerm/pull/95): Bump tflint-plugin-sdk to v0.8.1

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
