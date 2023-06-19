// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource_convert

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/resource_generate"
)

type converter struct {
	spec spec.Specification
}

func NewConverter(spec spec.Specification) converter {
	return converter{
		spec: spec,
	}
}

func (c converter) ToGeneratorResourceSchema() (map[string]resource_generate.GeneratorResourceSchema, error) {
	resourceSchemas := make(map[string]resource_generate.GeneratorResourceSchema, len(c.spec.Resources))

	for _, v := range c.spec.Resources {
		s, err := convertSchema(v)
		if err != nil {
			return nil, err
		}

		resourceSchemas[v.Name] = s
	}

	return resourceSchemas, nil
}

func isRequired(cor specschema.ComputedOptionalRequired) bool {
	return cor == specschema.Required
}

func isOptional(cor specschema.ComputedOptionalRequired) bool {
	if cor == specschema.Optional || cor == specschema.ComputedOptional {
		return true
	}

	return false
}

func isComputed(cor specschema.ComputedOptionalRequired) bool {
	if cor == specschema.Computed || cor == specschema.ComputedOptional {
		return true
	}

	return false
}

func isSensitive(s *bool) bool {
	if s == nil {
		return false
	}

	return *s
}

func description(d *string) string {
	if d == nil {
		return ""
	}

	return *d
}

func deprecationMessage(dm *string) string {
	if dm == nil {
		return ""
	}

	return *dm
}
