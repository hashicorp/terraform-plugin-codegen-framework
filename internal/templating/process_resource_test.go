package templating_test

import (
	"log/slog"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templating"
)

func TestProcessResourceTemplates(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		resourceTemplateData map[string]templating.ResourceTemplateData
		templateDir          fstest.MapFS
		want                 map[string][]byte
	}{
		"simple": {
			resourceTemplateData: map[string]templating.ResourceTemplateData{
				"simple_pet": {
					SnakeName:       "simple_pet",
					PascalName:      "SimplePet",
					CamelName:       "simplePet",
					Package:         "provider",
					SchemaFunc:      "SimplePetResourceSchema",
					SchemaModelType: "SimplePetModel",
				},
			},
			templateDir: fstest.MapFS{
				"simple_pet_resource.gotmpl": &fstest.MapFile{
					Data: []byte(`{{.SchemaFunc}}`),
				},
			},
			want: map[string][]byte{
				"simple_pet_resource_gen.go": []byte(`SimplePetResourceSchema`),
			},
		},
		"defaults": {
			resourceTemplateData: map[string]templating.ResourceTemplateData{
				"simple_pet": {
					SnakeName:       "simple_pet",
					PascalName:      "SimplePet",
					CamelName:       "simplePet",
					Package:         "provider",
					SchemaFunc:      "SimplePetResourceSchema",
					SchemaModelType: "SimplePetModel",
				},
				"simple_order": {
					SnakeName:       "simple_order",
					PascalName:      "SimpleOrder",
					CamelName:       "simpleOrder",
					Package:         "provider",
					SchemaFunc:      "SimpleOrderResourceSchema",
					SchemaModelType: "SimpleOrderModel",
				},
			},
			templateDir: fstest.MapFS{
				"resource_default.gotmpl": &fstest.MapFile{
					Data: []byte(`{{.SchemaFunc}}`),
				},
			},
			want: map[string][]byte{
				"simple_pet_resource_gen.go":   []byte(`SimplePetResourceSchema`),
				"simple_order_resource_gen.go": []byte(`SimpleOrderResourceSchema`),
			},
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			templator := templating.NewTemplator(slog.Default(), testCase.templateDir)

			got, err := templator.ProcessResources(testCase.resourceTemplateData)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}

			if len(got) != len(testCase.want) {
				t.Fatalf("unexpected number of files: got %d, wanted %d", len(got), len(testCase.want))
			}

			for name, wantBytes := range testCase.want {
				gotBytes, ok := got[name]
				if !ok {
					t.Errorf("did not find expected file: %s", name)
					continue
				}

				if diff := cmp.Diff(string(gotBytes), string(wantBytes)); diff != "" {
					t.Errorf("unexpected difference in %s: %s", name, diff)
				}
			}
		})
	}
}
