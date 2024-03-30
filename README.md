# TFLint Ruleset for terraform-provider-azurerm
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-azurerm/workflows/build/badge.svg?branch=master)](https://github.com/terraform-linters/tflint-ruleset-azurerm/actions)
[![GitHub release](https://img.shields.io/github/release/terraform-linters/tflint-ruleset-azurerm.svg)](https://github.com/terraform-linters/tflint-ruleset-azurerm/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

TFLint ruleset plugin for Terraform Provider for Azure (Resource Manager)

## Requirements

- TFLint v0.42+
- Go v1.22

## Installation

You can install the plugin by adding a config to `.tflint.hcl` and running `tflint --init`:

```hcl
plugin "azurerm" {
    enabled = true
    version = "0.25.1"
    source  = "github.com/terraform-linters/tflint-ruleset-azurerm"
}
```

## Rules

200+ rules are available. See the [documentation](docs/README.md).

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

Note that if you install the plugin with make install, you must omit the `version` and `source` attributes in `.tflint.hcl`:

```hcl
plugin "azurerm" {
    enabled = true
}
```

## Add a new rule

If you are interested in adding a new rule to this ruleset, you can use the generator. Run the following command:

```
$ go run ./rules/generator
```

Follow the instructions to edit the generated files and open a new pull request.
