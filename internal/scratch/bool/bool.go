package bool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Following examples are used to illustrate what should be created for a model
// depending upon the Intermediate Representation IR.

// IR for simple data source not using associated_external_type

/*
{
  "datasources": [
    {
      "name": "example",
      "schema": {
        "attributes": [
          {
            "name": "bool_attribute",
            "bool": {
              "computed_optional_required": "computed"
              "associated_external_type": {
                "import": "example.com/apisdk",
                "type" "*apisdk.ExampleBoolAttribute"
              }
            }
          }
        ]
      }
    }
  ]
}
*/

// Schema from IR for simple data source not using associated_external_type

//func (e *exampleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
//	resp.Schema = schema.Schema{
//		Attributes: map[string]schema.Attribute{
//			"bool_attribute": schema.BoolAttribute{
//				Computed: true,
//			},
//		},
//	}
//}

// apisdkCreateExampleResponse just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkCreateExampleResponse struct {
	BoolAttribute *apisdkExampleBoolAttribute
}

// apisdkCreateExampleRequest just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkCreateExampleRequest struct {
	BoolAttribute *apisdkExampleBoolAttribute
}

// apisdkThingConfiguration just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkExampleBoolAttribute *bool

type ExampleDataSourceModel struct {
	BoolAttribute types.Bool `tfsdk:"bool_attribute"`
}

func (m ExampleDataSourceModel) attributeType(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: m.objectAttributeTypes(ctx),
	}
}

func (m ExampleDataSourceModel) objectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"bool_attribute": types.BoolType,
	}
}

func (m ExampleDataSourceModel) objectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.objectAttributeTypes(ctx),
	)
}

func (m ExampleDataSourceModel) objectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.objectAttributeTypes(ctx),
		data,
	)
}

func ExampleDataSourceModelFromCreateExampleResponse(ctx context.Context, apiObject apisdkCreateExampleResponse) (ExampleDataSourceModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var model ExampleDataSourceModel

	// START: each attribute/field

	exampleBoolAttribute, exampleBoolAttributeDiags := ExampleBoolAttributeBoolFromExampleBoolAttribute(ctx, apiObject.BoolAttribute)

	diags.Append(exampleBoolAttributeDiags...)

	if diags.HasError() {
		return model, diags
	}

	model.BoolAttribute = exampleBoolAttribute

	// END: each attribute/field

	return model, diags
}

func ExampleBoolAttributeBoolFromExampleBoolAttribute(ctx context.Context, apiObject *apisdkExampleBoolAttribute) (types.Bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	if apiObject == nil {
		return types.BoolNull(), diags
	}

	return types.BoolPointerValue(*apiObject), diags
}

func (m ExampleDataSourceModel) ExampleDataSourceModelToCreateExampleRequest(ctx context.Context) (apisdkCreateExampleRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	apiObject := apisdkCreateExampleRequest{}

	// START: each attribute/field

	val := m.BoolAttribute.ValueBool()
	boolAttribute := apisdkExampleBoolAttribute(&val)

	// Set external field
	apiObject.BoolAttribute = &boolAttribute

	// END: each attribute/field

	return apiObject, diags
}
