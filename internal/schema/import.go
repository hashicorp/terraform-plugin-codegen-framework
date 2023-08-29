// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schema

import (
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

const (
	AttrImport         = "github.com/hashicorp/terraform-plugin-framework/attr"
	BaseTypesImport    = "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ContextImport      = "context"
	DiagImport         = "github.com/hashicorp/terraform-plugin-framework/diag"
	FmtImport          = "fmt"
	PlanModifierImport = "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	StringsImport      = "strings"
	TfTypesImport      = "github.com/hashicorp/terraform-plugin-go/tftypes"
	TypesImport        = "github.com/hashicorp/terraform-plugin-framework/types"
	ValidatorImport    = "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type Imports struct {
	imports []code.Import
	paths   map[string]struct{}
}

func NewImports() *Imports {
	return &Imports{
		imports: []code.Import{},
		paths:   make(map[string]struct{}),
	}
}

func (i *Imports) Add(c ...code.Import) {
	for _, imp := range c {
		if _, ok := i.paths[imp.Path]; ok {
			continue
		}

		i.imports = append(i.imports, imp)
		i.paths[imp.Path] = struct{}{}
	}
}

func (i *Imports) All() []code.Import {
	return i.imports
}

func (i *Imports) Append(imps ...*Imports) {
	for _, imp := range imps {
		for _, c := range imp.imports {
			if _, ok := i.paths[c.Path]; ok {
				continue
			}

			i.imports = append(i.imports, c)
			i.paths[c.Path] = struct{}{}
		}
	}
}

func GetElementTypeImports(e specschema.ElementType) *Imports {
	imports := NewImports()

	switch {
	case e.Bool != nil:
		if e.Bool.CustomType != nil && e.Bool.CustomType.HasImport() {
			imports.Add(*e.Bool.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: TypesImport,
		})
		return imports
	case e.Float64 != nil:
		if e.Float64.CustomType != nil && e.Float64.CustomType.HasImport() {
			imports.Add(*e.Float64.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: TypesImport,
		})
		return imports
	case e.Int64 != nil:
		if e.Int64.CustomType != nil && e.Int64.CustomType.HasImport() {
			imports.Add(*e.Int64.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: TypesImport,
		})
		return imports
	case e.List != nil:
		imports.Add(GetElementTypeImports(e.List.ElementType).All()...)
		return imports
	case e.Map != nil:
		imports.Add(GetElementTypeImports(e.Map.ElementType).All()...)
		return imports
	case e.Number != nil:
		if e.Number.CustomType != nil && e.Number.CustomType.HasImport() {
			imports.Add(*e.Number.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: TypesImport,
		})
		return imports
	case e.Object != nil:
		imports.Add(GetAttrTypesImports(e.Object.CustomType, e.Object.AttributeTypes).All()...)
		return imports
	case e.Set != nil:
		imports.Add(GetElementTypeImports(e.Set.ElementType).All()...)
		return imports
	case e.String != nil:
		if e.String.CustomType != nil && e.String.CustomType.HasImport() {
			imports.Add(*e.String.CustomType.Import)
			return imports
		}
		imports.Add(code.Import{
			Path: TypesImport,
		})
		return imports
	default:
		return imports
	}
}

func AttrImports() *Imports {
	imports := NewImports()

	imports.Add(code.Import{
		Path: AttrImport,
	})

	return imports
}

func GetAttrTypesImports(customType *specschema.CustomType, attrTypes []specschema.ObjectAttributeType) *Imports {
	imports := NewImports()

	if customType != nil && customType.HasImport() {
		imports.Add(*customType.Import)
	}

	if len(attrTypes) == 0 {
		return imports
	}

	for _, v := range attrTypes {
		switch {
		case v.Bool != nil:
			if v.Bool.CustomType != nil && v.Bool.CustomType.HasImport() {
				imports.Add(*v.Bool.CustomType.Import)
				continue
			}
			imports.Add(code.Import{
				Path: AttrImport,
			})
			imports.Add(code.Import{
				Path: TypesImport,
			})
		case v.Float64 != nil:
			if v.Float64.CustomType != nil && v.Float64.CustomType.HasImport() {
				imports.Add(*v.Float64.CustomType.Import)
				continue
			}
			imports.Add(code.Import{
				Path: AttrImport,
			})
			imports.Add(code.Import{
				Path: TypesImport,
			})
		case v.Int64 != nil:
			if v.Int64.CustomType != nil && v.Int64.CustomType.HasImport() {
				imports.Add(*v.Int64.CustomType.Import)
				continue
			}
			imports.Add(code.Import{
				Path: AttrImport,
			})
			imports.Add(code.Import{
				Path: TypesImport,
			})
		case v.List != nil:
			if v.List.CustomType != nil && v.List.CustomType.HasImport() {
				imports.Add(*v.List.CustomType.Import)
			}
			imports.Add(GetElementTypeImports(v.List.ElementType).All()...)
		case v.Map != nil:
			if v.Map.CustomType != nil && v.Map.CustomType.HasImport() {
				imports.Add(*v.Map.CustomType.Import)
			}
			imports.Add(GetElementTypeImports(v.Map.ElementType).All()...)
		case v.Number != nil:
			if v.Number.CustomType != nil && v.Number.CustomType.HasImport() {
				imports.Add(*v.Number.CustomType.Import)
				continue
			}
			imports.Add(code.Import{
				Path: AttrImport,
			})
			imports.Add(code.Import{
				Path: TypesImport,
			})
		case v.Object != nil:
			imports.Add(GetAttrTypesImports(v.Object.CustomType, v.Object.AttributeTypes).All()...)
		case v.Set != nil:
			if v.Set.CustomType != nil && v.Set.CustomType.HasImport() {
				imports.Add(*v.Set.CustomType.Import)
			}
			imports.Add(GetElementTypeImports(v.Set.ElementType).All()...)
		case v.String != nil:
			if v.String.CustomType != nil && v.String.CustomType.HasImport() {
				imports.Add(*v.String.CustomType.Import)
				continue
			}
			imports.Add(code.Import{
				Path: AttrImport,
			})
			imports.Add(code.Import{
				Path: TypesImport,
			})
		}
	}

	return imports
}

func CustomTypeImports(c *specschema.CustomType) *Imports {
	imports := NewImports()

	if c != nil {
		if c.HasImport() {
			imports.Add(*c.Import)
		}
	} else {
		imports.Add(code.Import{
			Path: TypesImport,
		})
	}

	return imports
}

func CustomDefaultImports(c *specschema.CustomDefault) *Imports {
	imports := NewImports()

	if c == nil {
		return imports
	}

	if !c.HasImport() {
		return imports
	}

	for _, i := range c.Imports {
		if len(i.Path) > 0 {
			imports.Add(i)
		}
	}

	return imports
}

func CustomPlanModifierImports(c *specschema.CustomPlanModifier) *Imports {
	imports := NewImports()

	if c == nil {
		return imports
	}

	if !c.HasImport() {
		return imports
	}

	for _, i := range c.Imports {
		if len(i.Path) > 0 {
			imports.Add(code.Import{
				Path: PlanModifierImport,
			})

			imports.Add(i)
		}
	}

	return imports
}

func CustomValidatorImports(cv *specschema.CustomValidator) *Imports {
	imports := NewImports()

	if cv == nil {
		return imports
	}

	if !cv.HasImport() {
		return imports
	}

	for _, i := range cv.Imports {
		if len(i.Path) > 0 {
			imports.Add(code.Import{
				Path: ValidatorImport,
			})

			imports.Add(i)
		}
	}

	return imports
}
