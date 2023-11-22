## 0.3.1 (November 22, 2023)

BUG FIXES:

* schema: Prevent compilation errors due to the generation of unused variables ([#93](https://github.com/hashicorp/terraform-plugin-codegen-framework/issues/93))

## 0.3.0 (November 14, 2023)

ENHANCEMENTS:

* Adds code generation for List, Map, Object, and Set attributes that have an associated external type ([#75](https://github.com/hashicorp/terraform-plugin-codegen-framework/issues/75))

BUG FIXES:

* Fix nested attribute name and generated custom value method name conflicts ([#81](https://github.com/hashicorp/terraform-plugin-codegen-framework/issues/81))

## 0.2.0 (October 24, 2023)

ENHANCEMENTS:

* Adds code generation for Bool, Float64, Int64, Number, and String attributes that have an associated external type ([#59](https://github.com/hashicorp/terraform-plugin-codegen-framework/issues/59))
* Adds usage of To/From methods for primitive attributes with an associated external type into To/From methods of nested attributes and blocks ([#73](https://github.com/hashicorp/terraform-plugin-codegen-framework/issues/73))

BUG FIXES:

* Allow Go reserved keywords to be used as attribute names in nested attributes ([#77](https://github.com/hashicorp/terraform-plugin-codegen-framework/issues/77))

## 0.1.0 (October 17, 2023)

NOTES:

* Initial release of `tfplugingen-framework` CLI for Terraform Provider Code Generation tech preview ([#61](https://github.com/hashicorp/terraform-plugin-codegen-framework/issues/61))

