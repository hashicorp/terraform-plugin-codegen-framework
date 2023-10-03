# terraform-plugin-codegen-framework

_Experimental: This code is under active development and not intended for production usage._

## Overview

Terraform Plugin Code Generation Framework is a binary which converts a [specification](https://github.com/hashicorp/terraform-plugin-codegen-spec) into Go code for use in a Terraform Plugin Provider.

The generator will output a schema, and the associated types for use in the CRUD operations performed by a Terraform provider. Additionally, scaffolding output can be generated to provide a skeleton of a Terraform provider.

## Usage

### Installation

```shell
go install github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest
```

### Generate Command

The generate command uses a specification](https://github.com/hashicorp/terraform-plugin-codegen-spec) as input and generates Terraform Provider code as output.

For example:

```shell
tfplugingen-framework generate all \
    --input specification.json \
    --output internal/provider
```

Refer to the [documentation](https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator#generate-command) for further details.

### Scaffold Command

The scaffold command generates skeleton code for a data source, provider, or resource.

For example:

```shell
tfplugingen-framework scaffold data-source \
    --name example \
    --force \
    --output-dir internal/provider
```

Refer to the [documentation](https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator#scaffold-command) for further details.

## License

Refer to [Mozilla Public License v2.0](./LICENSE).
