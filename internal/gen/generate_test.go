package gen

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"

	"github/hashicorp/terraform-provider-code-generator/internal/format"
)

func TestDataSourcesModels(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ir             string
		expectedSchema string
		expectedError  bool
	}{
		"datasource-bool": {
			ir:             "datasource-bool-ir.json",
			expectedSchema: "datasource-bool-models.txt",
		},
		"datasource-list": {
			ir:             "datasource-list-ir.json",
			expectedSchema: "datasource-list-models.txt",
		},
		"datasource-single-nested-attribute": {
			ir:             "datasource-single-nested-attribute-ir.json",
			expectedSchema: "datasource-single-nested-attribute-models.txt",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			irBytes, err := os.ReadFile(filepath.Join("fixtures", testCase.ir))
			if err != nil {
				t.Errorf("cannot read file: %s", err)
			}

			// unmarshal JSON
			ir := spec.Specification{}

			err = json.Unmarshal(irBytes, &ir)
			if err != nil {
				t.Error(err)
			}

			dataSourcesModelsGenerator := DataSourcesModelsGenerator{
				Templates: []string{
					"../templates/model/datasource_model.gotmpl",
					"../templates/model/attributes.gotmpl",
					"../templates/model/bool_attribute.gotmpl",
					"../templates/model/list_attribute.gotmpl",
					"../templates/model/single_nested_attribute.gotmpl",
					"../templates/model/single_nested_model.gotmpl",
				},
			}

			got, err := dataSourcesModelsGenerator.Process(ir)

			formattedGot, err := format.Format(got)
			if err != nil {
				t.Error(err)
			}

			expectedBytes, err := os.ReadFile(filepath.Join("fixtures", testCase.expectedSchema))
			if err != nil {
				t.Errorf("cannot read file: %s", err)
			}

			if diff := cmp.Diff(formattedGot["example"], expectedBytes); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}

			//if diff := cmp.Diff(diags, testCase.expectedDiags); diff != "" {
			//	t.Errorf("unexpected diagnostics difference: %s", diff)
			//}
		})
	}
}

func TestDataSourcesHelpers(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ir             string
		expectedSchema string
		expectedError  bool
	}{
		"datasource-bool": {
			ir:             "datasource-bool-ir.json",
			expectedSchema: "datasource-bool-helpers.txt",
		},
		"datasource-list": {
			ir:             "datasource-list-ir.json",
			expectedSchema: "datasource-list-helpers.txt",
		},
		"datasource-single-nested-attribute": {
			ir:             "datasource-single-nested-attribute-ir.json",
			expectedSchema: "datasource-single-nested-attribute-helpers.txt",
		},
		"datasource-single-nested-attribute-external-types": {
			ir:             "datasource-single-nested-attribute-external-types-ir.json",
			expectedSchema: "datasource-single-nested-attribute-external-types-helpers.txt",
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			irBytes, err := os.ReadFile(filepath.Join("fixtures", testCase.ir))
			if err != nil {
				t.Errorf("cannot read file: %s", err)
			}

			// unmarshal JSON
			ir := spec.Specification{}

			err = json.Unmarshal(irBytes, &ir)
			if err != nil {
				t.Error(err)
			}

			dataSourcesHelpersGenerator := DataSourcesHelpersGenerator{
				Templates: []string{
					"../templates/helper/datasource_helper.gotmpl",
					"../templates/helper/attributes.gotmpl",
					"../templates/helper/bool_attribute.gotmpl",
					"../templates/helper/list_attribute.gotmpl",
					"../templates/helper/elem_type.gotmpl",
					"../templates/helper/single_nested_attribute.gotmpl",
					"../templates/helper/single_nested_helper.gotmpl",
				},
			}

			got, err := dataSourcesHelpersGenerator.Process(ir)

			formattedGot, err := format.Format(got)
			if err != nil {
				t.Error(err)
			}

			expectedBytes, err := os.ReadFile(filepath.Join("fixtures", testCase.expectedSchema))
			if err != nil {
				t.Errorf("cannot read file: %s", err)
			}

			if diff := cmp.Diff(formattedGot["example"], expectedBytes); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}

			//if diff := cmp.Diff(diags, testCase.expectedDiags); diff != "" {
			//	t.Errorf("unexpected diagnostics difference: %s", diff)
			//}
		})
	}
}
