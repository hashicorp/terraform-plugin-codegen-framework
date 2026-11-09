// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func listNested(attributes generatorschema.GeneratorAttributes) GeneratorListNestedAttribute {
	return GeneratorListNestedAttribute{
		NestedObject: GeneratorNestedAttributeObject{
			Attributes: attributes,
		},
	}
}

func typeNames(attributes generatorschema.GeneratorAttributes, prefix string, out map[string]string) map[string]string {
	if out == nil {
		out = map[string]string{}
	}

	for k, v := range attributes {
		n, ok := v.(generatorschema.GoTypeNamer)

		if !ok {
			continue
		}

		out[prefix+k] = n.GoTypeName(k).ToPascalCase()

		if a, ok := v.(generatorschema.Attributes); ok {
			typeNames(a.GetAttributes(), prefix+k+".", out)
		}
	}

	return out
}

func TestGeneratorSchema_ResolveTypeNames(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         generatorschema.GeneratorAttributes
		expected      map[string]string
		expectedError string
	}{
		"no-collision": {
			input: generatorschema.GeneratorAttributes{
				"tags": listNested(generatorschema.GeneratorAttributes{
					"key": GeneratorStringAttribute{},
				}),
				"processors": listNested(generatorschema.GeneratorAttributes{
					"labels": listNested(generatorschema.GeneratorAttributes{
						"key": GeneratorStringAttribute{},
					}),
				}),
			},
			expected: map[string]string{
				"tags":              "Tags",
				"processors":        "Processors",
				"processors.labels": "Labels",
			},
		},
		"root-and-nested": {
			input: generatorschema.GeneratorAttributes{
				"tags": listNested(generatorschema.GeneratorAttributes{
					"key": GeneratorStringAttribute{},
				}),
				"processors": listNested(generatorschema.GeneratorAttributes{
					"tags": listNested(generatorschema.GeneratorAttributes{
						"key": GeneratorStringAttribute{},
					}),
				}),
			},
			expected: map[string]string{
				"tags":            "Tags",
				"processors":      "Processors",
				"processors.tags": "ProcessorsTags",
			},
		},
		"siblings": {
			input: generatorschema.GeneratorAttributes{
				"inputs": listNested(generatorschema.GeneratorAttributes{
					"tags": listNested(generatorschema.GeneratorAttributes{
						"key": GeneratorStringAttribute{},
					}),
				}),
				"outputs": listNested(generatorschema.GeneratorAttributes{
					"tags": listNested(generatorschema.GeneratorAttributes{
						"key": GeneratorStringAttribute{},
					}),
				}),
			},
			expected: map[string]string{
				"inputs":       "Inputs",
				"inputs.tags":  "InputsTags",
				"outputs":      "Outputs",
				"outputs.tags": "OutputsTags",
			},
		},
		"deep": {
			input: generatorschema.GeneratorAttributes{
				"inputs": listNested(generatorschema.GeneratorAttributes{
					"settings": listNested(generatorschema.GeneratorAttributes{
						"tags": listNested(generatorschema.GeneratorAttributes{
							"key": GeneratorStringAttribute{},
						}),
					}),
				}),
				"outputs": listNested(generatorschema.GeneratorAttributes{
					"settings": listNested(generatorschema.GeneratorAttributes{
						"tags": listNested(generatorschema.GeneratorAttributes{
							"key": GeneratorStringAttribute{},
						}),
					}),
				}),
			},
			expected: map[string]string{
				"inputs":                "Inputs",
				"inputs.settings":       "InputsSettings",
				"inputs.settings.tags":  "InputsSettingsTags",
				"outputs":               "Outputs",
				"outputs.settings":      "OutputsSettings",
				"outputs.settings.tags": "OutputsSettingsTags",
			},
		},
		"unresolvable": {
			input: generatorschema.GeneratorAttributes{
				"tags": listNested(generatorschema.GeneratorAttributes{
					"key": GeneratorStringAttribute{},
				}),
				"processors_tags": listNested(generatorschema.GeneratorAttributes{
					"key": GeneratorStringAttribute{},
				}),
				"processors": listNested(generatorschema.GeneratorAttributes{
					"tags": listNested(generatorschema.GeneratorAttributes{
						"key": GeneratorStringAttribute{},
					}),
				}),
			},
			expectedError: "ProcessorsTagsType/ProcessorsTagsValue for processors.tags, processors_tags",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := generatorschema.GeneratorSchema{Attributes: testCase.input}.ResolveTypeNames()

			if testCase.expectedError != "" {
				if err == nil || !strings.Contains(err.Error(), testCase.expectedError) {
					t.Fatalf("expected error containing %q, got %v", testCase.expectedError, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(typeNames(testCase.input, "", nil), testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSchema_ResolveTypeNames_DifferentStructures(t *testing.T) {
	t.Parallel()

	g := generatorschema.GeneratorSchema{
		Attributes: generatorschema.GeneratorAttributes{
			"inputs": listNested(generatorschema.GeneratorAttributes{
				"tags": listNested(generatorschema.GeneratorAttributes{
					"key": GeneratorStringAttribute{},
				}),
			}),
			"outputs": listNested(generatorschema.GeneratorAttributes{
				"tags": listNested(generatorschema.GeneratorAttributes{
					"count": GeneratorInt64Attribute{},
				}),
			}),
		},
	}

	if err := g.ResolveTypeNames(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	b, err := g.CustomTypeValueBytes()

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	for _, want := range []string{
		"type InputsTagsValue struct",
		"type OutputsTagsValue struct",
		"ElemType: InputsTagsValue{}.Type(ctx)",
		"ElemType: OutputsTagsValue{}.Type(ctx)",
	} {
		if !bytes.Contains(b, []byte(want)) {
			t.Errorf("expected generated code to contain %q", want)
		}
	}

	if bytes.Contains(b, []byte("type TagsValue struct")) {
		t.Errorf("unexpected TagsValue type in generated code")
	}
}
