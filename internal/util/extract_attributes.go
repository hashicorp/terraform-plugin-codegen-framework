package util

import (
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
)

type RequestTypeWithMethodAndPath struct {
	spec.RequestType
	Method string `json:"method"`
	Path   string `json:"path"`
}

type RequestWithMethodAndPath struct {
	Create RequestTypeWithMethodAndPath    `json:"create,omitempty"`
	Read   RequestTypeWithMethodAndPath    `json:"read"`
	Update []*RequestTypeWithMethodAndPath `json:"update"`
	Delete RequestTypeWithMethodAndPath    `json:"delete"`
}

type RequestWithRefreshObjectName struct {
	Create RequestTypeWithMethodAndPath     `json:"create,omitempty"`
	Read   RequestTypeWithRefreshObjectName `json:"read"`
	Update []*RequestTypeWithMethodAndPath  `json:"update"`
	Delete RequestTypeWithMethodAndPath     `json:"delete"`
	Name   string                           `json:"name"`
	Id     string                           `json:"id"`
}

type RequestTypeWithRefreshObjectName struct {
	RequestTypeWithMethodAndPath
	Response string `json:"response"`
}

type RequestWithResponse struct {
	Requests []spec.Request
}

type NcloudProvider struct {
	provider.Provider
	Endpoint string `json:"endpoint,omitempty"`
}

type NcloudSpecification struct {
	spec.Specification
	Provider    *NcloudProvider                `json:"provider"`
	Requests    []RequestWithRefreshObjectName `json:"requests"`
	Resources   []Resource                     `json:"resources"`
	DataSources []DataSource                   `json:"datasources"`
}

type Resource struct {
	resource.Resource
	RefreshObjectName   string `json:"refresh_object_name"`
	ImportStateOverride string `json:"import_state_override"`
	Id                  string `json:"id"`
}

type DataSource struct {
	datasource.DataSource
	RefreshObjectName   string `json:"refresh_object_name"`
	ImportStateOverride string `json:"import_state_override"`
	Id                  string `json:"id"`
}

type Schema struct {
	Attributes resource.Attributes `json:"attributes"`
}

type Attribute struct {
	resource.Attribute
	Computed bool `json:"computed"`
	Optional bool `json:"optional"`
	Required bool `json:"required"`
}

type SingleNestedAttributeType struct {
	ComputedOptionalRequired string               `json:"computed_optional_required"`
	Attributes               []resource.Attribute `json:"attributes"`
}

type ListAttributeType struct {
	ComputedOptionalRequired string                   `json:"computed_optional_required"`
	ElementType              ListAttributeElementType `json:"element_type"`
}

type ListAttributeElementType struct {
	String interface{} `json:"string,omitempty"`
	Bool   interface{} `json:"bool,omitempty"`
	Int64  interface{} `json:"int64,omitempty"`
}

type ListNestedAttributeType struct {
	ComputedOptionalRequired string           `json:"computed_optional_required"`
	NestedObject             NestedObjectType `json:"nested_object"`
}

type NestedObjectType struct {
	Attributes []resource.Attribute `json:"attributes"`
}
