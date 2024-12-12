# Ncloud Terraform Plugin Framework Code Generator

> Based on [Terraform Provider Code Generation](https://developer.hashicorp.com/terraform/plugin/code-generation), this is an experimental version for ncloud specific terraform code generation.
> All the features are should be referenced with [Terraform Plugin Codegen Framework](https://github.com/hashicorp/terraform-plugin-codegen-framework)

> _[Terraform Provider Code Generation](https://developer.hashicorp.com/terraform/plugin/code-generation) is currently in tech preview. If you have general questions or feedback about provider code generation, please create a new topic in the [Plugin Development Community Forum](https://discuss.hashicorp.com/c/terraform-providers/tf-plugin-sdk)._

## Overview

Terraform Plugin Framework Code Generator is a CLI tool which converts a [Provider Code Specification](https://developer.hashicorp.com//terraform/plugin/code-generation/specification) into Go code for use in a Terraform [Plugin Framework](https://developer.hashicorp.com/terraform/plugin/framework) Provider. Additionally, scaffolding commands with customizable templates are available which generate boilerplate provider code for new data sources and resources to reduce development time.

The generator currently supports outputting:

 * **[Schema](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/schemas)**: With all framework functionality, such as validators and plan modifiers, and no limits on nesting.
 * **[Data Model Types](https://developer.hashicorp.com/terraform/plugin/framework/handling-data/accessing-values#get-the-entire-configuration-plan-or-state)**: With conversion to external Go types, if provided in the specification, such as API SDK types.
 
Over time, it is anticipated that the Provider Code Specification and this generator will be further enhanced to support CRUD logic.

## Usage

### Installation

You install a copy of the binary manually from the [releases](https://github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/releases) tab, or install via the Go toolchain:

```shell
go install github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest
```

### Generate Command

[The generate command uses a specification](https://github.com/NaverCloudPlatform/terraform-plugin-codegen-openapi) as input and generates Terraform Provider code as output.

For example:

```shell
tfplugingen-framework generate all \
    --input specification.json \
    --output internal/provider
```

## How to write down config.yaml (Ncloud Specific)

### Provider

* `name` (`string`): (Required) Name of service. will be set as an package name.
* `endpoint` (`string`): (Required) Default endpoint of service.

### Resources

* `refresh_object_name` (`string`): (Optional) Name of schema object that represents resource. Will be used in refreshing logic. If not provided, default value is **READ response object**.
  
* `id` (`string`): (Required) How to access unique id of resource from **CREATE response object**.
  
* `import_state_override` (`string`): (Optional) Required attribute information to get state from import operation(READ).
  
* `Create, Read, Delete`: Commonly required attributes are as belows.
  * `path` (`string`): (Required) Path of CREATE operation.
  * `method` (`string`): (Required) Method of CREATE operation.
  
* `Update`: Update is type of array with objects. Need to write down path, method information in first element.

### Example of `config.yml`

```yaml
provider:
  name: apigw
  endpoint: "https://apigateway.apigw.ntruss.com/api/v1"
resources:
  product:
    refresh_object_name: PostProductResponse
    id: product.product_id
    create:
      path: /products
      method: POST
    read:
      path: /products/{product-id}
      method: GET
    update:
      - path: /products/{product-id}
        method: PATCH
    delete:
      path: /products/{product-id}
      method: DELETE
  api_keys:
    refresh_object_name: ApiKeyDto
    id: api_key.api_key_id
    create:
      path: /api-keys
      method: POST
    read:
      path: /api-keys/{api-key-id}
      method: GET
    update:
      - path: /api-keys/{api-key-id}
        method: PUT
    delete:
      path: /api-keys/{api-key-id}
      method: DELETE
  resource:
    refresh_object_name: ResourceDto
    id: resource.resourceId
    import_state_override: product-id.api-id.resource-id
    create:
      path: /products/{product-id}/apis/{api-id}/resources
      method: POST
    read:
      path: /products/{product-id}/apis/{api-id}/resources
      method: GET
    update:
      - path: /products/{product-id}/apis/{api-id}/resources/{resource-id}
        method: PATCH
    delete:
      path: /products/{product-id}/apis/{api-id}/resources/{resource-id}
      method: DELETE
data_sources:
  product:
    refresh_object_name: PostProductResponse
    id: product.product_id
    read:
      path: /products/{product-id}
      method: GET
```

## License

Refer to [Mozilla Public License v2.0](./LICENSE).

## Experimental Status

By using the software in this repository (the "Software"), you acknowledge that: (1) the Software is still in development, may change, and has not been released as a commercial product by HashiCorp and is not currently supported in any way by HashiCorp; (2) the Software is provided on an "as-is" basis, and may include bugs, errors, or other issues; (3) the Software is NOT INTENDED FOR PRODUCTION USE, use of the Software may result in unexpected results, loss of data, or other unexpected results, and HashiCorp disclaims any and all liability resulting from use of the Software; and (4) HashiCorp reserves all rights to make all decisions about the features, functionality and commercial release (or non-release) of the Software, at any time and without any obligation or liability whatsoever.
