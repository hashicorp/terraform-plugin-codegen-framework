package single_nested_string

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// IR for resource with single nested attribute using associated_external_type

/*
{
  "name": "thing",
  "schema": {
    "attributes": [
      {
        "name": "configuration",
        "single_nested": {
          "attributes": [
            {
              "name": "description",
              "string": {
                "computed_optional_required": "optional",
              },
            }
          ],
          "computed_optional_required": "optional",
          "associated_external_type": {
            "import": "example.com/apisdk",
            "type" "*apisdk.ThingConfiguration"
          }
        }
      },
      {
        "name": "name",
        "string": {
          "computed_optional_required": "required"
        },
      }
    ]
  }
}
*/

// Schema from IR for resource with single nested attribute using associated_external_type

//func (e *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
//	resp.Schema = schema.Schema{
//		Attributes: map[string]schema.Attribute{
//			"configuration": schema.SingleNestedAttribute{
//				Attributes: map[string]schema.Attribute{
//					"Description": schema.StringAttribute{
//						Optional: true,
//					},
//				},
//				Optional: true,
//			},
//			"name": schema.StringAttribute{
//				Required: true,
//			},
//		},
//	}
//}

// apisdkCreateThingResponse just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkCreateThingResponse struct {
	Configuration *apisdkThingConfiguration
	Name          *string
}

// apisdkCreateThingRequest just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkCreateThingRequest struct {
	Configuration *apisdkThingConfiguration
	Name          *string
}

// apisdkThingConfiguration just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkThingConfiguration struct {
	Description *string
}

// Model from IR for data source with single nested attribute using associated_external_type
// 		All the following code is expected to be generated (i.e., model and all funcs/methods).

type ThingResourceModel struct {
	Configuration types.Object `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
}

func (m ThingResourceModel) objectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"configuration": types.ObjectType{
			AttrTypes: ThingConfigurationModel{}.objectAttributeTypes(ctx),
		},
		"name": types.StringType,
	}
}

func (m ThingResourceModel) objectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.objectAttributeTypes(ctx),
	)
}

func (m ThingResourceModel) objectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.objectAttributeTypes(ctx),
		data,
	)
}

// ThingResourceModelFromCreateThingResponse should transform the API response to ThingResourceModel.
//
// Assuming that the response looks as follows:
//
//	type CreateThingResponse struct {
//		Configuration *ThingConfiguration
//		Name          *string
//	}
//
//	type ThingConfiguration struct {
//		Description *string
//	}
func ThingResourceModelFromCreateThingResponse(ctx context.Context, apiObject apisdkCreateThingResponse) (ThingResourceModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var model ThingResourceModel

	// START: each attribute/field

	thingConfigurationModel, thingConfigurationModelDiags := ThingConfigurationObjectFromThingConfiguration(ctx, apiObject.Configuration)

	diags.Append(thingConfigurationModelDiags...)

	if diags.HasError() {
		return model, diags
	}

	model.Configuration = thingConfigurationModel

	if apiObject.Name != nil {
		model.Name = types.StringValue(*apiObject.Name)
	}

	// END: each attribute/field

	return model, diags
}

func ThingConfigurationObjectFromThingConfiguration(ctx context.Context, apiObject *apisdkThingConfiguration) (types.Object, diag.Diagnostics) {
	var diags diag.Diagnostics
	m := ThingConfigurationModel{}

	if apiObject == nil {
		return m.objectNull(ctx), diags
	}

	m.Description = types.StringPointerValue(apiObject.Description)

	return m.objectValueFrom(ctx, m)
}

// ToCreateThingRequest should transform the TF types used in the ThingResourceModel into a request object that can be used
// by apisdk in a CreateThingRequest.
//
// Given the IR and schema, CreateThingRequest should look as follows:
//
//	type CreateThingRequest struct {
//		Configuration *ThingConfiguration
//		Name          *string
//	}
//
//	type ThingConfiguration struct {
//		Description *string
//	}
func (m ThingResourceModel) ToCreateThingRequest(ctx context.Context) (apisdkCreateThingRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	apiObject := apisdkCreateThingRequest{}

	// START: each attribute/field

	// For descending nesting levels, convert to model
	// TODO: Handle list, map, set.
	// TODO: Handle nulls and unknowns
	tfThingConfigurationModel, tfThingConfigurationModelDiags := ThingConfigurationModelFromObject(ctx, m.Configuration)

	diags.Append(tfThingConfigurationModelDiags...)

	if diags.HasError() {
		return apiObject, diags
	}

	// Delegate external type conversion to model
	apiThingConfiguration, apiThingConfigurationDiags := tfThingConfigurationModel.ToThingConfiguration(ctx)

	diags.Append(apiThingConfigurationDiags...)

	if diags.HasError() {
		return apiObject, diags
	}

	// Set external field
	apiObject.Configuration = apiThingConfiguration

	// For values, just do the conversion
	apiObject.Name = m.Name.ValueStringPointer()

	// END: each attribute/field

	return apiObject, diags
}

func ThingConfigurationModelFromObject(ctx context.Context, tfObject types.Object) (ThingConfigurationModel, diag.Diagnostics) {
	var model ThingConfigurationModel

	diags := tfObject.As(ctx, &model, basetypes.ObjectAsOptions{})

	return model, diags
}

type ThingConfigurationModel struct {
	Description types.String `tfsdk:"description"`
}

func (m ThingConfigurationModel) ToThingConfiguration(ctx context.Context) (*apisdkThingConfiguration, diag.Diagnostics) {
	var diags diag.Diagnostics

	// TODO: Handle cases where ThingConfigurationModel contains lists, maps, objects or sets.
	apiObject := &apisdkThingConfiguration{
		Description: m.Description.ValueStringPointer(),
	}

	return apiObject, diags
}

func (m ThingConfigurationModel) objectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"description": types.StringType,
	}
}

func (m ThingConfigurationModel) objectNull(ctx context.Context) types.Object {
	return types.ObjectNull(
		m.objectAttributeTypes(ctx),
	)
}

func (m ThingConfigurationModel) objectValueFrom(ctx context.Context, data any) (types.Object, diag.Diagnostics) {
	return types.ObjectValueFrom(
		ctx,
		m.objectAttributeTypes(ctx),
		data,
	)
}
