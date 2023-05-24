# Questions

## Associated External Type and Primitives

* Are there any instances in which a primitive type (i.e., bool, float, int, string) could have a 
    type defined for _associated_external_type_, and if so, what would the accompanying expand and 
    flatten functions look like?
  * At first glance, it seems as though primitives would just be pointers to Go types (e.g., *bool,
    *float, *int, *string).

```json
{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "bool_attribute",
            "bool": {
              "computed_optional_required": "computed",
              "associated_external_type": {
                "import": "example.com/apisdk",
                "type": "*apisdk.Bool"
              }
            }
          }
        ]
      }
    }
  ]
}
```

### Answer

* Supporting custom primitive external types will require that a corresponding custom framework
    type also be defined. The custom framework type would be responsible for the mapping to/from
    the custom external type. This will require supporting basetype interfaces 
    (e.g., [StringValuable](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework/types/basetypes#StringValuable)).
* Without a corresponding custom framework type, primitive attributes will only be able to support
    pointers and values (e.g., `bool` and `*bool`).

```json
          {
            "name": "bool_attribute",
            "bool": {
              "computed_optional_required": "computed",
              "associated_external_type": {
                "type": "*bool"
              }
            }
          }
```

## Associated External Types and Collection Types - List, Map & Set

* Are there any instances in which a collection type (i.e., list, map, set) could have a type 
    defined for _associated_external_type_, and if so, what would the accompanying expand and 
    flatten functions look like?
  * At first glance, it seems as though collection types would just be Go types (i.e., list => slice,
    map => map, set => slice). 

```json
{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "list_attribute",
            "list": {
              "computed_optional_required": "computed",
              "associated_external_type": {
                "import": "example.com/apisdk",
                "type": "*apisdk.List"
              }
            }
          }
        ]
      }
    }
  ]
}
```