# Automation tools

This directory contains Go scripts for automating various operations. You can run `go run ./subdir` in this directory.

## apispec-rule-gen

This script generate rules from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). The repository has been added as a Git submodule and you will update the submodule version if you want to keep up with the latest.

The correspondence between API definitions and Terraform attributes is defined in the [mapping files](apispec-rule-gen/mappings). It also includes Terraform's schema file to check for invalid mappings.

### Debugging with VS Code

* Open the root of this repo in [VS Code](https://code.visualstudio.com/).
* Install the [Go Extension for VS Code](https://github.com/golang/vscode-go) via the Extensions icon on the left hand menu.
* Press F5 to start the debugger or click on 'Run' -> 'Start Debugging'

The launch.json file provides default values for the tool base path, rule generation target path and doc generation target path. You can override these if needed.

```json
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/tools/apispec-rule-gen/",
      "env": {},
      "args": [
        "-base-path=.",
        "-rules-path=../../rules",
        "-docs-path=../../docs"
      ]
    }
  ]
}
```

### Update azure-rest-api-specs

```console
$ cd apispec-rule-gen/azure-rest-api-specs
$ git pull origin main
$ cd ../../
$ go run ./apispec-rule-gen
```

### Update terraform-provider-azurerm

```console
$ cd apispec-rule-gen/schema
$ tfenv install
$ terraform init -upgrade
$ terraform providers schema -json > schema.json
$ cd ../../
$ go run ./apispec-rule-gen
```

## api-version-bumper

This script parses the [terraform-provider-azurerm](https://github.com/hashicorp/terraform-provider-azurerm) codebase and updates the mapping files with the identified API versions. The repository has been added as a Git submodule and you will update the submodule version if you want to keep up with the latest.

This tool is meant to be used together with apispec-rule-gen, and a typical workflow is to update API versions in mapping files and then generate rules with apispec-rule-gen.

### Update terraform-provider-azurerm

```console
$ cd api-version-bumper/terraform-provider-azurerm
$ git pull origin main
$ cd ../../
$ go run ./api-version-bumper
```

##  release

This script automates the release process. Run `go run ./release` and follow the instructions to modify and commit the files required for releasing and push them to the remote. The actual build and publish process is triggerd by GitHub Actions and is not the responsibility of this script.
