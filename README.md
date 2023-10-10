# Terraform Plugin Framework Code Generator

> _[Terraform Provider Code Generation](https://developer.hashicorp.com/terraform/plugin/code-generation) is currently in tech preview._

## Overview

Terraform Plugin Framework Code Generator is a CLI tool which converts a [Provider Code Specification](https://developer.hashicorp.com//terraform/plugin/code-generation/specification) into Go code for use in a Terraform [Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework) Provider.

The generator currently supports outputting:

 * **Schema**: With all framework functionality, such as validators and plan modifiers, and no limits on nesting.
 * **Data Model Types**: With conversion to external Go types, if provided in the specification, such as API SDK types.
 
Over time, it is anticipated that the Provider Code Specification and this generator will be further enhanced to support CRUD logic.

## Documentation

Full usage info and examples live on the HashiCorp developer site: https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator

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

The scaffold command generates starter code for a data source, provider, or resource to reduce initial development effort. The templates can be customized to match provider code conventions and automatically include API client configuration.

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
