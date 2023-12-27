package templating_test

import (
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/templating"
)

func TestProcessProviderTemplates(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		providerTemplateData map[string]templating.ProviderTemplateData
		templateDir          fstest.MapFS
		want                 map[string][]byte
	}{
		"simple": {
			providerTemplateData: map[string]templating.ProviderTemplateData{
				"petstore": {
					SnakeName:       "simple_pet",
					PascalName:      "SimplePet",
					CamelName:       "simplePet",
					Package:         "provider",
					SchemaFunc:      "SimplePetProviderSchema",
					SchemaModelType: "SimplePetModel",
				},
			},
			templateDir: fstest.MapFS{
				"provider.gotmpl": &fstest.MapFile{
					Data: []byte(`{{.SchemaFunc}}`),
				},
			},
			want: map[string][]byte{
				"provider_gen.go": []byte(`SimplePetProviderSchema`),
			},
		},
	}
	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			templator := templating.NewTemplator(testCase.templateDir)

			got, err := templator.ProcessProvider(testCase.providerTemplateData)
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
