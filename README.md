# TFLint Ruleset for terraform-provider-azurerm
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-azurerm/workflows/build/badge.svg?branch=master)](https://github.com/terraform-linters/tflint-ruleset-azurerm/actions)
[![GitHub release](https://img.shields.io/github/release/terraform-linters/tflint-ruleset-azurerm.svg)](https://github.com/terraform-linters/tflint-ruleset-azurerm/releases/latest)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

TFLint ruleset plugin for Terraform Provider for Azure (Resource Manager)

## Requirements

- TFLint v0.19+
- Go v1.14

## Installation

Download the plugin and place it in `~/.tflint.d/plugins/tflint-ruleset-azurerm` (or `./.tflint.d/plugins/tflint-ruleset-azurerm`). When using the plugin, configure as follows in `.tflint.hcl`:

```hcl
plugin "azurerm" {
    enabled = true
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
