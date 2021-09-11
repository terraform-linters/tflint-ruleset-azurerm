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
# Edit provider.tf
$ terraform init -upgrade
$ terraform providers schema -json > schema.json
$ cd ../../
$ go run ./apispec-rule-gen
```
