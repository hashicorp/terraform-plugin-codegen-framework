// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestFrameworkIdentifier_Valid(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identifier schema.FrameworkIdentifier
		want       bool
	}{
		"invalid - empty": {
			identifier: "",
			want:       false,
		},
		"invalid - middle hyphen": {
			identifier: "fake-thing",
			want:       false,
		},
		"invalid - leading numeric": {
			identifier: "1_fake_thing",
			want:       false,
		},
		"invalid - uppercase": {
			identifier: "fake_Thing",
			want:       false,
		},
		"valid - lowercase alphabet": {
			identifier: "thing",
			want:       true,
		},
		"valid - leading underscore": {
			identifier: "_thing",
			want:       true,
		},
		"valid - middle underscore": {
			identifier: "fake_thing",
			want:       true,
		},
		"valid - alphanumeric": {
			identifier: "thing123",
			want:       true,
		},
		"valid - alphanumeric with underscores": {
			identifier: "fake_thing_123",
			want:       true,
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.identifier.Valid()
			if got != testCase.want {
				t.Fatalf("expected Valid() to return %t, got %t", testCase.want, got)
			}
		})
	}
}

func TestFrameworkIdentifier_ToCamelCase(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identifier schema.FrameworkIdentifier
		want       string
	}{
		"lowercase alphabet": {
			identifier: "thing",
			want:       "thing",
		},
		"leading underscore": {
			identifier: "_thing",
			want:       "thing",
		},
		"middle underscore": {
			identifier: "fake_thing",
			want:       "fakeThing",
		},
		"alphanumeric": {
			identifier: "thing123",
			want:       "thing123",
		},
		"valid - alphanumeric with underscores": {
			identifier: "fake_thing_123",
			want:       "fakeThing123",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.identifier.ToCamelCase()
			if got != testCase.want {
				t.Fatalf("expected ToCamelCase() to return %s, got %s", testCase.want, got)
			}
		})
	}
}

func TestFrameworkIdentifier_ToPrefixCamelCase(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identifier schema.FrameworkIdentifier
		want       string
	}{
		"lowercase alphabet": {
			identifier: "thing",
			want:       "thing",
		},
		"leading underscore": {
			identifier: "_thing",
			want:       "thing",
		},
		"middle underscore": {
			identifier: "fake_thing",
			want:       "fakeThing",
		},
		"alphanumeric": {
			identifier: "thing123",
			want:       "thing123",
		},
		"valid - alphanumeric with underscores": {
			identifier: "fake_thing_123",
			want:       "fakeThing123",
		},
		"attribute_types": {
			identifier: "attribute_types",
			want:       "nameAttributeTypes",
		},
		"equal": {
			identifier: "Equal",
			want:       "nameEqual",
		},
		"is_null": {
			identifier: "is_null",
			want:       "nameIsNull",
		},
		"is_unknown": {
			identifier: "is_unknown",
			want:       "nameIsUnknown",
		},
		"string": {
			identifier: "string",
			want:       "nameString",
		},
		"to_object_value": {
			identifier: "to_object_value",
			want:       "nameToObjectValue",
		},
		"to_terraform_value": {
			identifier: "to_terraform_value",
			want:       "nameToTerraformValue",
		},
		"type": {
			identifier: "type",
			want:       "nameType",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.identifier.ToPrefixCamelCase("name")
			if got != testCase.want {
				t.Fatalf("expected ToCamelCase() to return %s, got %s", testCase.want, got)
			}
		})
	}
}

func TestFrameworkIdentifier_ToPascalCase(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identifier schema.FrameworkIdentifier
		want       string
	}{
		"lowercase alphabet": {
			identifier: "thing",
			want:       "Thing",
		},
		"leading underscore": {
			identifier: "_thing",
			want:       "Thing",
		},
		"middle underscore": {
			identifier: "fake_thing",
			want:       "FakeThing",
		},
		"alphanumeric": {
			identifier: "thing123",
			want:       "Thing123",
		},
		"valid - alphanumeric with underscores": {
			identifier: "fake_thing_123",
			want:       "FakeThing123",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.identifier.ToPascalCase()
			if got != testCase.want {
				t.Fatalf("expected ToPascalCase() to return %s, got %s", testCase.want, got)
			}
		})
	}
}

func TestFrameworkIdentifier_ToPrefixPascalCase(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identifier schema.FrameworkIdentifier
		want       string
	}{
		"lowercase alphabet": {
			identifier: "thing",
			want:       "Thing",
		},
		"leading underscore": {
			identifier: "_thing",
			want:       "Thing",
		},
		"middle underscore": {
			identifier: "fake_thing",
			want:       "FakeThing",
		},
		"alphanumeric": {
			identifier: "thing123",
			want:       "Thing123",
		},
		"valid - alphanumeric with underscores": {
			identifier: "fake_thing_123",
			want:       "FakeThing123",
		},
		"attribute_types": {
			identifier: "attribute_types",
			want:       "NameAttributeTypes",
		},
		"equal": {
			identifier: "Equal",
			want:       "NameEqual",
		},
		"is_null": {
			identifier: "is_null",
			want:       "NameIsNull",
		},
		"is_unknown": {
			identifier: "is_unknown",
			want:       "NameIsUnknown",
		},
		"string": {
			identifier: "string",
			want:       "NameString",
		},
		"to_object_value": {
			identifier: "to_object_value",
			want:       "NameToObjectValue",
		},
		"to_terraform_value": {
			identifier: "to_terraform_value",
			want:       "NameToTerraformValue",
		},
		"type": {
			identifier: "type",
			want:       "NameType",
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.identifier.ToPrefixPascalCase("name")
			if got != testCase.want {
				t.Fatalf("expected ToPascalCase() to return %s, got %s", testCase.want, got)
			}
		})
	}
}
