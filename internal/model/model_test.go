package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestObjectHelper_String(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    ObjectHelper
		expected string
	}{
		"example": {
			input: ObjectHelper{
				Name: "example",
				AttrTypes: map[string]string{
					"key_name": "types.StringType",
				},
			},
			expected: `func (m exampleModel) objectType() types.ObjectType {
return types.ObjectType{AttrTypes: m.objectAttributeTypes()}
}

func (m exampleModel) objectAttributeTypes() map[string]attr.Type {
return map[string]attr.Type{
"key_name": types.StringType,
}
}`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.String()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
