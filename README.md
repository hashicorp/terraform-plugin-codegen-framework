# Terraform Plugin Framework Code Generator

> _[Terraform Provider Code Generation](https://developer.hashicorp.com/terraform/plugin/code-generation) is currently in tech preview. If you have general questions or feedback about provider code generation, please create a new topic in the [Plugin Development Community Forum](https://discuss.hashicorp.com/c/terraform-providers/tf-plugin-sdk)._

## Overview

Terraform Plugin Framework Code Generator is a CLI tool which converts a [Provider Code Specification](https://developer.hashicorp.com//terraform/plugin/code-generation/specification) into Go code for use in a Terraform [Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework) Provider. Additionally, scaffolding commands with customizable templates are available which generate boilerplate provider code for new data sources and resources to reduce development time.

The generator currently supports outputting:

 * **[Schema](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/schemas)**: With all framework functionality, such as validators and plan modifiers, and no limits on nesting.
 * **[Data Model Types](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/accessing-values#get-the-entire-configuration-plan-or-state)**: With conversion to external Go types, if provided in the specification, such as API SDK types.
 
Over time, it is anticipated that the Provider Code Specification and this generator will be further enhanced to support CRUD logic.

## Documentation

Full usage info and examples live on the HashiCorp developer site: https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator

## Usage

### Installation

You install a copy of the binary manually from the [releases](https://github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/releases) tab, or install via the Go toolchain:

```shell
go install github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest
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

## Experimental Status

By using the software in this repository (the "Software"), you acknowledge that: (1) the Software is still in development, may change, and has not been released as a commercial product by HashiCorp and is not currently supported in any way by HashiCorp; (2) the Software is provided on an "as-is" basis, and may include bugs, errors, or other issues; (3) the Software is NOT INTENDED FOR PRODUCTION USE, use of the Software may result in unexpected results, loss of data, or other unexpected results, and HashiCorp disclaims any and all liability resulting from use of the Software; and (4) HashiCorp reserves all rights to make all decisions about the features, functionality and commercial release (or non-release) of the Software, at any time and without any obligation or liability whatsoever.
