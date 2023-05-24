# Intermediate Representation to Code Generator

## Running the Generator

The generator reads the intermediate representation (IR) from _stdin_ by default in order to 
facilitate the chaining together of CLI commands.

The following is a contrived example:

```shell
cat examples/ir.json | go run . schema
```

An alternative is to use the `-input` flag to specify a file from which the IR can be read.

For example:

```shell
go run . schema -input examples/ir.json
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

Currently, the only command that has been implemented is `schema`. This is a bit of a misnomer 
as the `schema` command generates schema, models and model helper functions but _only_ for data
sources currently.

## Further Considerations

### Input

* Do we want to have the default for input for both the IR and the IR schema be stdin?

### Validate

* Do we want to improve the detail of the error messages that are generated when the IR is
  not valid JSON or is this considered outside the scope of the IR2FCG?

### Generate

* Go code generation is leveraging [text/template](https://pkg.go.dev/text/template). Currently,
  testing of the code generation is using files which contain a textual representation of the 
  expected generated code. This is quite difficult to visually inspect as no syntax highlighting
  is provided by an IDE. Perhaps there's a better/easier way to handle this process?
* The templates that are used for generating Go code are `*.gotmpl` files which do not benefit 
  from Go syntax highlighting via an IDE. Again, perhaps there's a better/easier way to deal
  with the templates.