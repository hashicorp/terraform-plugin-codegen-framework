// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider_convert

import (
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/provider_generate"
)

type converter struct {
	spec spec.Specification
}

func NewConverter(spec spec.Specification) converter {
	return converter{
		spec: spec,
	}
}

func (c converter) ToGeneratorProviderSchema() (map[string]provider_generate.GeneratorProviderSchema, error) {
	providerSchemas := make(map[string]provider_generate.GeneratorProviderSchema, len(c.spec.DataSources))

	s, err := convertSchema(c.spec.Provider)
	if err != nil {
		return nil, err
	}

	providerSchemas[c.spec.Provider.Name] = s

	return providerSchemas, nil
}

func isRequired(cor specschema.OptionalRequired) bool {
	return cor == specschema.Required
}

func isOptional(cor specschema.OptionalRequired) bool {
	if cor == specschema.Optional || cor == specschema.ComputedOptional {
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
