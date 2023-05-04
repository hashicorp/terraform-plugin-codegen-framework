package validate

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
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
func Schema(inputFile string) error {
	document, err := os.ReadFile(inputFile)

	if err != nil {
		return err
	}

	return spec.Validate(context.TODO(), document)
}
