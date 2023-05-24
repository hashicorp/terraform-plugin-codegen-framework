package provider_generate

import (
	"fmt"
	"strings"
	"text/template"

	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GeneratorObjectAttribute struct {
	schema.ObjectAttribute

	CustomType *specschema.CustomType
	Validators []specschema.ObjectValidator
}

func (g GeneratorObjectAttribute) Equal(ga GeneratorAttribute) bool {
	if _, ok := ga.(GeneratorObjectAttribute); !ok {
		return false
	}

	goa := ga.(GeneratorObjectAttribute)

	if !customTypeEqual(g.CustomType, goa.CustomType) {
		return false
	}

	if !g.validatorsEqual(g.Validators, goa.Validators) {
		return false
	}

	return g.ObjectAttribute.Equal(goa.ObjectAttribute)
}

func (g GeneratorObjectAttribute) ToString(name string) (string, error) {
	funcMap := template.FuncMap{
		"getAttrTypes": getAttrTypes,
	}

	t, err := template.New("object_attribute").Funcs(funcMap).Parse(objectAttributeGoTemplate)
	if err != nil {
		return "", err
	}

	if _, err = t.New("common_attribute").Parse(commonAttributeGoTemplate); err != nil {
		return "", err
	}

	var buf strings.Builder

	attrib := map[string]GeneratorObjectAttribute{
		name: g,
	}

	err = t.Execute(&buf, attrib)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g GeneratorObjectAttribute) validatorsEqual(x, y []specschema.ObjectValidator) bool {
	if x == nil && y == nil {
		return true
	}

	if x == nil && y != nil {
		return false
	}

	if x != nil && y == nil {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	//TODO: Sort before comparing.
	for k, v := range x {
		if v.Custom == nil && y[k].Custom != nil {
			return false
		}

		if v.Custom != nil && y[k].Custom == nil {
			return false
		}

		if v.Custom != nil && y[k].Custom != nil {
			if *v.Custom.Import != *y[k].Custom.Import {
				return false
			}
		}

		if v.Custom.SchemaDefinition != y[k].Custom.SchemaDefinition {
			return false
		}
	}

	return true
}

func getAttrTypes(attrTypes map[string]attr.Type) string {
	var aTypes strings.Builder

	for k, v := range attrTypes {
		switch t := v.(type) {
		case basetypes.BoolType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.BoolType,", k))
		case basetypes.Float64Type:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.Float64Type,", k))
		case basetypes.Int64Type:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.Int64Type,", k))
		case types.ListType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.ListType{\nElemType: %s,\n},", k, getElementType(t.ElementType())))
		case types.MapType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.MapType{\nElemType: %s,\n},", k, getElementType(t.ElementType())))
		case basetypes.NumberType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.NumberType,", k))
		case types.ObjectType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.ObjectType{\nAttrTypes: map[string]attr.Type{\n%s\n},\n},", k, getAttrTypes(t.AttrTypes)))
		case types.SetType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.SetType{\nElemType: %s,\n},", k, getElementType(t.ElementType())))
		case basetypes.StringType:
			aTypes.WriteString(fmt.Sprintf("\"%s\": types.StringType,", k))
		}
	}

	return aTypes.String()
}
