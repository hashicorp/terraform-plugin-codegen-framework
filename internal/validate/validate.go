package validate

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func JSON(input []byte) error {
	if !json.Valid(input) {
		return errors.New("invalid JSON")
	}

	return nil
}

// Schema
// TODO: Check for duplicate keys at same nesting level in JSON.
// TODO: Handle schema when supplied as URL
func Schema(inputFile, schemaFile string) []error {
	var errs []error

	documentLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", inputFile))
	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", schemaFile))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			errs = append(errs, fmt.Errorf("%s", desc.String()))
		}
	}

	return errs
}
