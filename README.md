# Intermediate Representation to Code Generator

## Running the Generator

### Build (local)
```shell
# Build the binary from source
# Creates binary named `tfplugingen-framework`
make build
```

### Input

The generator reads an [Intermediate Representation (IR)](https://github.com/hashicorp/terraform-plugin-codegen-spec) of a Terraform Provider. Input is read from **stdin** by default in order to facilitate the chaining together of CLI commands.

The following examples use an IR from the repo's integration tests [internal/cmd/testdata/custom_and_external/ir.json](./internal/cmd/testdata/custom_and_external/ir.json):

```shell
cat internal/cmd/testdata/custom_and_external/ir.json | ./tfplugingen-framework generate all
```

An alternative is to use the `--input` flag to specify a file from which the IR can be read.

For example:

```shell
./tfplugingen-framework generate all --input internal/cmd/testdata/custom_and_external/ir.json
```

### Commands
The IR JSON file contains `provider`, `resources`, and `datasources` definitions. These can all be processed together or individually with the following commands:

```shell
# Generates all code for provider, resources, and data-sources
./tfplugingen-framework generate all --input internal/cmd/testdata/custom_and_external/ir.json

# Generates all code for data-sources only.
./tfplugingen-framework generate data-sources --input internal/cmd/testdata/custom_and_external/ir.json

# Generates all code for provider only.
./tfplugingen-framework generate provider --input internal/cmd/testdata/custom_and_external/ir.json

# Generates all code for resources only.
./tfplugingen-framework generate resources --input internal/cmd/testdata/custom_and_external/ir.json
```

### Output

The generated code will default to the `./output` directory, but can also be specified with the `--output` parameter. Similarly, the name of the Go package in the generated code will default to `provider`, but can be specified with `--package`.
```shell
# Generates all code into a Go package named `generated` at the directory path `./internal/provider/generated`
./tfplugingen-framework generate all --input internal/cmd/testdata/custom_and_external/ir.json --output internal/provider/generated --package generated
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

