# Intermediate Representation to Code Generator

## Running the Generator

The generator reads the intermediate representation (IR) from _stdin_ by default in order to 
facilitate the chaining together of CLI commands.

The following is a contrived example:

```shell
cat examples/ir.json | go run . all
```

An alternative is to use the `-input` flag to specify a file from which the IR can be read.

For example:

```shell
go run . all -input examples/ir.json
```

Both cases will process `ir.json`.

The generated code will be saved into the `generator/output` directory.

`ir.json` contains a simple intermediate representation (IR).

## Running the Tests

```shell
go test $(go list ./... | grep -v /output) -v -count=1
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
  `generator/output`.

Currently, the only command that has been implemented is `all`. This is a bit of a misnomer 
as the `all` command only generates schema for data sources, provider and resources at 
present.
