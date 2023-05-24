package list_nested

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// IR for resource with list nested attribute using associated_external_type

/*
{
  "name": "thing",
  "schema": {
    "attributes": [
      {
        "name": "configuration",
        "list_nested": {
          "nested_object": {
            "attributes": [
              {
                "name": "description",
                "string": {
                  "computed_optional_required": "optional"
                }
              }
            ]
          },
          "computed_optional_required": "optional",
          "associated_external_type": {
            "import": "example.com/apisdk",
            "type": "*apisdk.ThingConfiguration"
          }
        }
      },
      {
        "name": "name",
        "string": {
          "computed_optional_required": "required"
        }
      }
    ]
  }
}
*/

// Schema from IR for resource with list nested attribute using associated_external_type

//func (e *thingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
//	resp.Schema = schema.Schema{
//		Attributes: map[string]schema.Attribute{
//			"configuration": schema.ListNestedAttribute{
//				NestedObject: schema.NestedAttributeObject{
//					Attributes: map[string]schema.Attribute{
//						"Description": schema.StringAttribute{
//							Optional: true,
//						},
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

// apisdkCreateThingRequest just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkCreateThingRequest struct {
	Configuration []*apisdkThingConfiguration
	Name          *string
}

// apisdkCreateThingResponse just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkCreateThingResponse struct {
	Configuration []*apisdkThingConfiguration
	Name          *string
}

// apisdkThingConfiguration just added here, so it's available in the same package but would
// not be created during code generation.
type apisdkThingConfiguration struct {
	Description *string
}

// Model from IR for data source with list nested attribute using associated_external_type
// 		All the following code is expected to be generated (i.e., model and all funcs/methods).

type ThingResourceModel struct {
	Configuration types.List   `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
}

func (m ThingResourceModel) objectAttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"configuration": types.ListType{
			ElemType: m.configurationListElementType(ctx),
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

func (m ThingResourceModel) configurationListElementType(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: ThingConfigurationModel{}.objectAttributeTypes(ctx),
	}
}

func (m ThingResourceModel) configurationListNull(ctx context.Context) types.List {
	return types.ListNull(
		m.configurationListElementType(ctx),
	)
}

func (m ThingResourceModel) configurationListValueFrom(ctx context.Context, data any) (types.List, diag.Diagnostics) {
	return types.ListValueFrom(
		ctx,
		m.configurationListElementType(ctx),
		data,
	)
}

type ThingConfigurationModel struct {
	Description types.String `tfsdk:"description"`
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

// ThingResourceModelFromCreateThingResponse should transform the API response to ThingResourceModel.
//
// Assuming that the response looks as follows:
//
//	type CreateThingResponse struct {
//		Configuration []*ThingConfiguration
//		Name          *string
//	}
//
//	type ThingConfiguration struct {
//		Description *string
//	}
func ThingResourceModelFromCreateThingResponse(ctx context.Context, apiObject apisdkCreateThingResponse) (ThingResourceModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var m ThingResourceModel

	// START: each attribute/field

	thingConfigurationModelList, thingConfigurationModelListDiags := ThingConfigurationListFromThingConfigurations(ctx, apiObject.Configuration)

	diags.Append(thingConfigurationModelListDiags...)

	if diags.HasError() {
		return m, diags
	}

	m.Configuration = thingConfigurationModelList

	if apiObject.Name != nil {
		m.Name = types.StringValue(*apiObject.Name)
	}

	// END: each attribute/field

	return m, diags
}

func ThingConfigurationListFromThingConfigurations(ctx context.Context, thingConfigurationSlice []*apisdkThingConfiguration) (types.List, diag.Diagnostics) {
	var diags diag.Diagnostics
	var m ThingResourceModel
	list := m.configurationListNull(ctx)

	for _, v := range thingConfigurationSlice {
		elem, d := ThingConfigurationObjectFromThingConfiguration(ctx, v)

		diags.Append(d...)

		if diags.HasError() {
			return list, diags
		}

		listElems := list.Elements()
		listElems = append(listElems, elem)

		list, d = m.configurationListValueFrom(ctx, listElems)
	}

	return list, diags
}

func ThingConfigurationObjectFromThingConfiguration(ctx context.Context, apiObject *apisdkThingConfiguration) (types.Object, diag.Diagnostics) {
	var diags diag.Diagnostics
	var m ThingConfigurationModel

	if apiObject == nil {
		return m.objectNull(ctx), diags
	}

	m.Description = types.StringPointerValue(apiObject.Description)

	return m.objectValueFrom(ctx, m)
}

// TODO: Should the calls to ElementsAs, generating the slice of models and generating the slice of *apisdkThingConfiguration
// be split out into separate functions?

func (m ThingResourceModel) ToCreateThingRequest(ctx context.Context) (apisdkCreateThingRequest, diag.Diagnostics) {
	var diags diag.Diagnostics
	apiObject := apisdkCreateThingRequest{}

	// START: each attribute/field

	// For descending nesting levels, convert to model
	configurationModels, d := m.ThingConfigurationModelsFromList(ctx)

	diags.Append(d...)

	if diags.HasError() {
		return apiObject, diags
	}

	// Delegate external type conversion to model
	apiThingConfigurations, d := ThingConfigurationsFromThingConfigurationModels(ctx, configurationModels)

	diags.Append(d...)

	if diags.HasError() {
		return apiObject, diags
	}

	// Set external field
	apiObject.Configuration = apiThingConfigurations

	// For values, just do the conversion
	apiObject.Name = m.Name.ValueStringPointer()

	// END: each attribute/field

	return apiObject, diags
}

func (m ThingResourceModel) ThingConfigurationModelsFromList(ctx context.Context) ([]ThingConfigurationModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var thingConfigurationModels []ThingConfigurationModel
	var objs []types.Object

	d := m.Configuration.ElementsAs(ctx, &objs, false)

	diags.Append(d...)

	if diags.HasError() {
		return thingConfigurationModels, diags
	}

	for _, v := range objs {
		tfThingConfigurationModel, tfThingConfigurationModelDiags := ThingConfigurationModelFromObject(ctx, v)

		diags.Append(tfThingConfigurationModelDiags...)

		if diags.HasError() {
			return thingConfigurationModels, diags
		}

		thingConfigurationModels = append(thingConfigurationModels, tfThingConfigurationModel)
	}

	return thingConfigurationModels, diags
}

func ThingConfigurationModelFromObject(ctx context.Context, tfObject types.Object) (ThingConfigurationModel, diag.Diagnostics) {
	var model ThingConfigurationModel

	diags := tfObject.As(ctx, &model, basetypes.ObjectAsOptions{})

	return model, diags
}

func ThingConfigurationsFromThingConfigurationModels(ctx context.Context, thingConfigurationModels []ThingConfigurationModel) ([]*apisdkThingConfiguration, diag.Diagnostics) {
	var diags diag.Diagnostics
	var apiThingConfigurations []*apisdkThingConfiguration

	for _, v := range thingConfigurationModels {

		apiThingConfiguration, apiThingConfigurationDiags := v.ToThingConfiguration(ctx)

		diags.Append(apiThingConfigurationDiags...)

		if diags.HasError() {
			return apiThingConfigurations, diags
		}

		apiThingConfigurations = append(apiThingConfigurations, apiThingConfiguration)
	}

	return apiThingConfigurations, diags

}

func (m ThingConfigurationModel) ToThingConfiguration(ctx context.Context) (*apisdkThingConfiguration, diag.Diagnostics) {
	var diags diag.Diagnostics

	apiObject := &apisdkThingConfiguration{
		Description: m.Description.ValueStringPointer(),
	}

	return apiObject, diags
}
