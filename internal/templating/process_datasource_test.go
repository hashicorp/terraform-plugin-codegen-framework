package templating_test

import (
	"log/slog"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templating"
)

func TestProcessDataSourceTemplates(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		datasourceTemplateData map[string]templating.DataSourceTemplateData
		templateDir            fstest.MapFS
		want                   map[string][]byte
	}{
		"simple": {
			datasourceTemplateData: map[string]templating.DataSourceTemplateData{
				"simple_pet": {
					SnakeName:       "simple_pet",
					PascalName:      "SimplePet",
					CamelName:       "simplePet",
					Package:         "provider",
					SchemaFunc:      "SimplePetDataSourceSchema",
					SchemaModelType: "SimplePetModel",
				},
			},
			templateDir: fstest.MapFS{
				"simple_pet_datasource.gotmpl": &fstest.MapFile{
					Data: []byte(`{{.SchemaFunc}}`),
				},
			},
			want: map[string][]byte{
				"simple_pet_datasource_gen.go": []byte(`SimplePetDataSourceSchema`),
			},
		},
		"defaults": {
			datasourceTemplateData: map[string]templating.DataSourceTemplateData{
				"simple_pet": {
					SnakeName:       "simple_pet",
					PascalName:      "SimplePet",
					CamelName:       "simplePet",
					Package:         "provider",
					SchemaFunc:      "SimplePetDataSourceSchema",
					SchemaModelType: "SimplePetModel",
				},
				"simple_order": {
					SnakeName:       "simple_order",
					PascalName:      "SimpleOrder",
					CamelName:       "simpleOrder",
					Package:         "provider",
					SchemaFunc:      "SimpleOrderDataSourceSchema",
					SchemaModelType: "SimpleOrderModel",
				},
			},
			templateDir: fstest.MapFS{
				"datasource_default.gotmpl": &fstest.MapFile{
					Data: []byte(`{{.SchemaFunc}}`),
				},
			},
			want: map[string][]byte{
				"simple_pet_datasource_gen.go":   []byte(`SimplePetDataSourceSchema`),
				"simple_order_datasource_gen.go": []byte(`SimpleOrderDataSourceSchema`),
			},
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			templator := templating.NewTemplator(slog.Default(), testCase.templateDir)

			got, err := templator.ProcessDataSources(testCase.datasourceTemplateData)
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
