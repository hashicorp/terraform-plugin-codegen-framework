package single_nested_string_without_external_types

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// IR for resource with single nested attribute without associated_external_type

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
          "computed_optional_required": "optional"
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

// Model from IR for data source with single nested attribute using associated_external_type
// 		All the following code is expected to be generated (i.e., model and all funcs/methods).

type ThingResourceModel struct {
	Configuration types.Object `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
}

// objectAttributeTypes returns the map[string]attr.Type for fields at this level (i.e., Configuration and Name).
// Retrieval of AttrTypes for the Configuration object is delegated to ThingConfigurationModel{}.objectAttributeTypes().
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

type ThingConfigurationModel struct {
	Description types.String `tfsdk:"description"`
}

// objectAttributeTypes returns the map[string]attr.Type for fields at this level (i.e., Description).
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
