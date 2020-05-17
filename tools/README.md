# Automation tools

This directory contains Go scripts for automating various operations. You can run `go run ./subdir` in this directory.

## apispec-rule-gen

This script generate rules from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). The repository has been added as a Git submodule and you will update the submodule version if you want to keep up with the latest.

The correspondence between API definitions and Terraform attributes is defined in the [mapping files](apispec-rule-gen/mappings). It also includes Terraform's schema file to check for invalid mappings.

### Update azure-rest-api-specs

```console
$ cd apispec-rule-gen/azure-rest-api-specs
$ git pull origin master
$ cd ../../
$ go run ./apispec-rule-gen
```

### Update terraform-provider-azurerm

```console
$ cd apispec-rule-gen/schema
$ tfenv install
# Edit provider.tf to update provider version
$ terraform show --json > schema.json
$ cd ../../
$ go run ./apispec-rule-gen
```
