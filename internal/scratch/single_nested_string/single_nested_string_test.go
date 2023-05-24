package single_nested_string

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestThingResourceModelFromCreateThingResponse(t *testing.T) {
	t.Parallel()

	description := "description"
	name := "name"

	testCases := map[string]struct {
		apiObject     apisdkCreateThingResponse
		expected      ThingResourceModel
		expectedDiags diag.Diagnostics
	}{
		"configuration-name": {
			apiObject: apisdkCreateThingResponse{
				Configuration: &apisdkThingConfiguration{
					Description: &description,
				},
				Name: &name,
			},
			expected: ThingResourceModel{
				Configuration: types.ObjectValueMust(
					map[string]attr.Type{
						"description": types.StringType,
					},
					map[string]attr.Value{
						"description": types.StringValue("description"),
					}),
				Name: types.StringValue("name"),
			},
		},
		"nil-configuration-name": {
			apiObject: apisdkCreateThingResponse{
				Configuration: nil,
				Name:          &name,
			},
			expected: ThingResourceModel{
				Configuration: types.ObjectNull(
					map[string]attr.Type{
						"description": types.StringType,
					},
				),
				Name: types.StringValue("name"),
			},
		},
		"nil-nil": {
			apiObject: apisdkCreateThingResponse{
				Configuration: nil,
				Name:          nil,
			},
			expected: ThingResourceModel{
				Configuration: types.ObjectNull(
					map[string]attr.Type{
						"description": types.StringType,
					},
				),
				Name: types.StringNull(),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, diags := ThingResourceModelFromCreateThingResponse(context.Background(), testCase.apiObject)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}

			if diff := cmp.Diff(diags, testCase.expectedDiags); diff != "" {
				t.Errorf("unexpected diagnostics difference: %s", diff)
			}
		})
	}
}

func TestToCreateThingRequest(t *testing.T) {
	t.Parallel()

	description := "description"
	name := "name"

	testCases := map[string]struct {
		thingResourceModel ThingResourceModel
		expected           apisdkCreateThingRequest
		expectedDiags      diag.Diagnostics
	}{
		"configuration-name": {
			thingResourceModel: ThingResourceModel{
				Configuration: types.ObjectValueMust(
					map[string]attr.Type{
						"description": types.StringType,
					},
					map[string]attr.Value{
						"description": types.StringValue("description"),
					}),
				Name: types.StringValue("name"),
			},
			expected: apisdkCreateThingRequest{
				Configuration: &apisdkThingConfiguration{
					Description: &description,
				},
				Name: &name,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, diags := testCase.thingResourceModel.ToCreateThingRequest(context.Background())

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}

			if diff := cmp.Diff(diags, testCase.expectedDiags); diff != "" {
				t.Errorf("unexpected diagnostics difference: %s", diff)
			}
		})
	}
}
