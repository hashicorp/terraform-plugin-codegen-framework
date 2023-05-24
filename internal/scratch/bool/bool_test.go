package bool

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestThingResourceModelFromCreateThingResponse(t *testing.T) {
	t.Parallel()

	tr := true
	a := apisdkExampleBoolAttribute(&tr)

	testCases := map[string]struct {
		apiObject     apisdkCreateExampleResponse
		expected      ExampleDataSourceModel
		expectedDiags diag.Diagnostics
	}{
		"success": {
			apiObject: apisdkCreateExampleResponse{
				BoolAttribute: &a,
			},
			expected: ExampleDataSourceModel{
				BoolAttribute: types.BoolValue(true),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, diags := ExampleDataSourceModelFromCreateExampleResponse(context.Background(), testCase.apiObject)

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

	tr := true
	a := apisdkExampleBoolAttribute(&tr)

	testCases := map[string]struct {
		model         ExampleDataSourceModel
		expected      apisdkCreateExampleRequest
		expectedDiags diag.Diagnostics
	}{
		"success": {
			model: ExampleDataSourceModel{
				BoolAttribute: types.BoolValue(true),
			},
			expected: apisdkCreateExampleRequest{
				BoolAttribute: &a,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, diags := testCase.model.ExampleDataSourceModelToCreateExampleRequest(context.Background())

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}

			if diff := cmp.Diff(diags, testCase.expectedDiags); diff != "" {
				t.Errorf("unexpected diagnostics difference: %s", diff)
			}
		})
	}
}
