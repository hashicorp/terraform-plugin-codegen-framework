// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource_convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/datasource_generate"
)

type converter struct {
	spec spec.Specification
}

func NewConverter(spec spec.Specification) converter {
	return converter{
		spec: spec,
	}
}

func (c converter) ToGeneratorDataSourceSchema() (map[string]datasource_generate.GeneratorDataSourceSchema, error) {
	dataSourceSchemas := make(map[string]datasource_generate.GeneratorDataSourceSchema, len(c.spec.DataSources))

	for _, v := range c.spec.DataSources {
		s, err := convertSchema(v)
		if err != nil {
			return nil, err
		}

		dataSourceSchemas[v.Name] = s
	}

	return dataSourceSchemas, nil
}

func convertElementType(e specschema.ElementType) (attr.Type, error) {
	switch {
	case e.Bool != nil:
		return types.BoolType, nil
	case e.Float64 != nil:
		return types.Float64Type, nil
	case e.Int64 != nil:
		return types.Int64Type, nil
	case e.List != nil:
		elemType, err := convertElementType(e.List.ElementType)
		if err != nil {
			return nil, err
		}
		return types.ListType{
			ElemType: elemType,
		}, nil
	case e.Map != nil:
		elemType, err := convertElementType(e.Map.ElementType)
		if err != nil {
			return nil, err
		}
		return types.MapType{
			ElemType: elemType,
		}, nil
	case e.Number != nil:
		return types.NumberType, nil
	case e.Object != nil:
		attrType, err := convertAttrTypes(e.Object)
		if err != nil {
			return nil, err
		}
		return types.ObjectType{
			AttrTypes: attrType,
		}, nil
	case e.Set != nil:
		elemType, err := convertElementType(e.Set.ElementType)
		if err != nil {
			return nil, err
		}
		return types.SetType{
			ElemType: elemType,
		}, nil
	case e.String != nil:
		return types.StringType, nil
	}

	return nil, fmt.Errorf("element type is not defined: %+v", e)
}

func convertAttrTypes(o []specschema.ObjectAttributeType) (map[string]attr.Type, error) {
	attrTypes := make(map[string]attr.Type, len(o))

	for _, v := range o {
		switch {
		case v.Bool != nil:
			attrTypes[v.Name] = types.BoolType
		case v.Float64 != nil:
			attrTypes[v.Name] = types.Float64Type
		case v.Int64 != nil:
			attrTypes[v.Name] = types.Int64Type
		case v.List != nil:
			elemType, err := convertElementType(v.List.ElementType)
			if err != nil {
				return nil, err
			}
			attrTypes[v.Name] = types.ListType{
				ElemType: elemType,
			}
		case v.Map != nil:
			elemType, err := convertElementType(v.Map.ElementType)
			if err != nil {
				return nil, err
			}
			attrTypes[v.Name] = types.MapType{
				ElemType: elemType,
			}
		case v.Number != nil:
			attrTypes[v.Name] = types.NumberType
		case v.Object != nil:
			aTypes, err := convertAttrTypes(v.Object)
			if err != nil {
				return nil, err
			}
			aTypes[v.Name] = types.ObjectType{
				AttrTypes: aTypes,
			}
		case v.Set != nil:
			elemType, err := convertElementType(v.Set.ElementType)
			if err != nil {
				return nil, err
			}
			attrTypes[v.Name] = types.MapType{
				ElemType: elemType,
			}
		case v.String != nil:
			attrTypes[v.Name] = types.StringType
		default:
			return nil, fmt.Errorf("attribute type not defined: %+v", v)
		}
	}

	return attrTypes, nil
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
