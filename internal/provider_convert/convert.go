package provider_convert

import (
	"fmt"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github/hashicorp/terraform-provider-code-generator/internal/provider_generate"
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
