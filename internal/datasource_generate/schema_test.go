package datasource_generate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TestGeneratorDataSourceSchema_ModelObjectHelpersString(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorDataSourceSchema
		expected      []byte
		expectedError error
	}{
		"bool": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"bool": GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{
							Optional: true,
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
}
}`),
		},
		"list_nested": {
			input: GeneratorDataSourceSchema{
				Attributes: map[string]GeneratorAttribute{
					"bool": GeneratorBoolAttribute{
						BoolAttribute: schema.BoolAttribute{},
					},
					"list_nested": GeneratorListNestedAttribute{
						NestedObject: GeneratorNestedAttributeObject{
							Attributes: map[string]GeneratorAttribute{
								"bool_nested": GeneratorBoolAttribute{},
							},
						},
					},
				},
			},
			expected: []byte(`
func (m ExampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ExampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool": types.BoolType,
"list_nested": types.ListType{
ElemType: ListNestedModel{}.objectType(),
},
}
}

func (m ListNestedModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m ListNestedModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"bool_nested": types.BoolType,
}
}`),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelObjectHelpersTemplate("example")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
