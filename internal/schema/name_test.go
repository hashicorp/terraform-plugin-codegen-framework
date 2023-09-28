// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestName_CamelCase(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name     Name
		expected string
	}{
		"snake_case": {
			name:     "bool_attribute",
			expected: "BoolAttribute",
		},
		"lower_case_first": {
			name:     "boolAttribute",
			expected: "BoolAttribute",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.name.CamelCase()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestName_CamelCaseLCFirst(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name     Name
		expected string
	}{
		"snake_case": {
			name:     "bool_attribute",
			expected: "boolAttribute",
		},
		"lower_case_first": {
			name:     "boolAttribute",
			expected: "boolAttribute",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.name.CamelCaseLCFirst()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestName_String(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name     Name
		expected string
	}{
		"string": {
			name:     "boolAttribute",
			expected: "boolAttribute",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.name.String()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
