# Intermediate Representation to Code Generator

## Running the Generator

### Build (local)
```shell
# Build the binary from source
# Creates binary named `terraform-plugin-codegen-framework`
make build
```

### Input

The generator reads an [Intermediate Representation (IR)](https://github.com/hashicorp/terraform-plugin-codegen-spec) of a Terraform Provider. Input is read from **stdin** by default in order to facilitate the chaining together of CLI commands.

The following is a contrived example:

```shell
cat examples/ir.json | ./terraform-plugin-codegen-framework generate all
```

An alternative is to use the `--input` flag to specify a file from which the IR can be read.

For example:

```shell
./terraform-plugin-codegen-framework generate all --input examples/ir.json
```

### Commands
The IR JSON file contains `provider`, `resources`, and `datasources` definitions. These can all be processed together or individually with the following commands:

```shell
# Generates all code for provider, resources, and data-sources
./terraform-plugin-codegen-framework generate all --input examples/ir.json

# Generates all code for data-sources only.
./terraform-plugin-codegen-framework generate data-sources --input examples/ir.json

# Generates all code for provider only.
./terraform-plugin-codegen-framework generate provider --input examples/ir.json

# Generates all code for resources only.
./terraform-plugin-codegen-framework generate resources --input examples/ir.json
```

### Output

The generated code will default to the `./output` directory, but can also be specified with the `--output` parameter. Similarly, the name of the Go package in the generated code will default to `provider`, but can be specified with `--package`.
```shell
# Generates all code into a Go package named `generated` at the directory path `./internal/provider/generated`
./terraform-plugin-codegen-framework generate all --input examples/ir.json --output internal/provider/generated --package generated
```

## Running the Tests

```shell
make test
```

## Overview

The Intermediate Representation to Code Generator (IR2CG) uses 
[mitchell/cli](https://github.com/mitchellh/cli) to implement a command-line interface that
provides processing of an IR and conversion of the IR into Go code that can be used within
a Terraform provider built using the 
[Terraform Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework). 

The general flow is as follows:

* Parse command-line flags for use as configuration during processing.
* Read the IR.
* Validate that the IR is JSON.
* Validate the IR against the IR Schema.
* Unmarshall the IR onto Go types.
* Generate Go code for schema, models and model helper functions.
* Format the generated Go code.
* Write the formatted Go code, one file per data source, provider or resource, into
  `./output`.

