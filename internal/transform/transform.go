package transform

import "encoding/json"

func Unmarshal(data []byte) (IntermediateRepresentation, error) {
	var r IntermediateRepresentation
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *IntermediateRepresentation) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type IntermediateRepresentation struct {
	DataSources []DataSource `json:"datasources,omitempty"`
	Provider    Provider     `json:"provider"`
	Resources   []Resource   `json:"resources,omitempty"`
}

type DataSource struct {
	Name   string           `json:"name"`
	Schema DataSourceSchema `json:"schema"`
}

type DataSourceSchema struct {
	Attributes []DataSourceAttribute `json:"attributes,omitempty"`
	Blocks     []DataSourceBlock     `json:"blocks,omitempty"`
}

type DataSourceAttribute struct {
	Name         string                           `json:"name"`
	Bool         *DataSourceBoolAttribute         `json:"bool,omitempty"`
	List         *DataSourceListAttribute         `json:"list,omitempty"`
	ListNested   *DataSourceListNestedAttribute   `json:"list_nested,omitempty"`
	Object       *DataSourceObjectAttribute       `json:"object,omitempty"`
	SingleNested *DataSourceSingleNestedAttribute `json:"single_nested,omitempty"`
}

type DataSourceBoolAttribute struct {
	ComputedOptionalRequired ComputedOptionalRequired `json:"computed_optional_required"`
	CustomType               *CustomType              `json:"custom_type,omitempty"`
	DeprecationMessage       *string                  `json:"deprecation_message,omitempty"`
	Description              *string                  `json:"description,omitempty"`
	Sensitive                *bool                    `json:"sensitive,omitempty"`
	Validators               []Validators             `json:"validators,omitempty"`
}

type DataSourceListAttribute struct {
	ComputedOptionalRequired ComputedOptionalRequired `json:"computed_optional_required"`
	CustomType               *CustomType              `json:"custom_type,omitempty"`
	DeprecationMessage       *string                  `json:"deprecation_message,omitempty"`
	Description              *string                  `json:"description,omitempty"`
	ElementType              ElementType              `json:"element_type"`
	Sensitive                *bool                    `json:"sensitive,omitempty"`
	Validators               []Validators             `json:"validators,omitempty"`
}

type DataSourceListNestedAttribute struct {
	ComputedOptionalRequired ComputedOptionalRequired        `json:"computed_optional_required"`
	CustomType               *CustomType                     `json:"custom_type,omitempty"`
	DeprecationMessage       *string                         `json:"deprecation_message,omitempty"`
	Description              *string                         `json:"description,omitempty"`
	NestedObject             DataSourceAttributeNestedObject `json:"nested_object"`
	Sensitive                *bool                           `json:"sensitive,omitempty"`
	Validators               []Validators                    `json:"validators,omitempty"`
}

type DataSourceAttributeNestedObject struct {
	Attributes []DataSourceAttribute `json:"attributes"`
	CustomType *CustomType           `json:"custom_type,omitempty"`
	Validators []Validators          `json:"validators,omitempty"`
}

type DataSourceObjectAttribute struct {
	ComputedOptionalRequired ComputedOptionalRequired `json:"computed_optional_required"`
	CustomType               *CustomType              `json:"custom_type,omitempty"`
	DeprecationMessage       *string                  `json:"deprecation_message,omitempty"`
	Description              *string                  `json:"description,omitempty"`
	ObjectType               []ObjectElement          `json:"object_type"`
	Sensitive                *bool                    `json:"sensitive,omitempty"`
	Validators               []Validators             `json:"validators,omitempty"`
}

type DataSourceSingleNestedAttribute struct {
	Attributes               []DataSourceAttribute    `json:"attributes"`
	ComputedOptionalRequired ComputedOptionalRequired `json:"computed_optional_required"`
	CustomType               *CustomType              `json:"custom_type,omitempty"`
	DeprecationMessage       *string                  `json:"deprecation_message,omitempty"`
	Description              *string                  `json:"description,omitempty"`
	Sensitive                *bool                    `json:"sensitive,omitempty"`
	Validators               []Validators             `json:"validators,omitempty"`
}

type CustomType struct {
	Import    string `json:"import"`
	Type      string `json:"type"`
	ValueType string `json:"value_type"`
}

type ElementType struct {
	Bool   *BoolElement    `json:"bool,omitempty"`
	List   *ListElement    `json:"list,omitempty"`
	Map    *MapElement     `json:"map,omitempty"`
	Object []ObjectElement `json:"object,omitempty"`
	String *StringElement  `json:"string,omitempty"`
}

type BoolElement struct {
}

type ListElement struct {
	Bool   *BoolElement    `json:"bool,omitempty"`
	List   *ListElement    `json:"list,omitempty"`
	Map    *MapElement     `json:"map,omitempty"`
	Object []ObjectElement `json:"object,omitempty"`
	String *StringElement  `json:"string,omitempty"`
}

type MapElement struct {
	Bool   *BoolElement    `json:"bool,omitempty"`
	List   *ListElement    `json:"list,omitempty"`
	Map    *MapElement     `json:"map,omitempty"`
	Object []ObjectElement `json:"object,omitempty"`
	String *StringElement  `json:"string,omitempty"`
}

type ObjectElement struct {
	Bool   *BoolElement    `json:"bool,omitempty"`
	List   *ListElement    `json:"list,omitempty"`
	Map    *MapElement     `json:"map,omitempty"`
	Name   string          `json:"name"`
	Object []ObjectElement `json:"object,omitempty"`
	String *StringElement  `json:"string,omitempty"`
}

type StringElement struct {
}

type DataSourceBlock struct {
	Name         string                       `json:"name"`
	ListNested   *DataSourceListNestedBlock   `json:"list_nested,omitempty"`
	SingleNested *DataSourceSingleNestedBlock `json:"single_nested,omitempty"`
}

type DataSourceListNestedBlock struct {
	CustomType         *CustomType                 `json:"custom_type,omitempty"`
	DeprecationMessage *string                     `json:"deprecation_message,omitempty"`
	Description        *string                     `json:"description,omitempty"`
	NestedObject       DataSourceBlockNestedObject `json:"nested_object"`
	Validators         []Validators                `json:"validators,omitempty"`
}

type DataSourceSingleNestedBlock struct {
	Attributes         []DataSourceAttribute `json:"attributes,omitempty"`
	Blocks             []DataSourceBlock     `json:"blocks,omitempty"`
	CustomType         *CustomType           `json:"custom_type,omitempty"`
	DeprecationMessage *string               `json:"deprecation_message,omitempty"`
	Description        *string               `json:"description,omitempty"`
	Validators         []Validators          `json:"validators,omitempty"`
}

type DataSourceBlockNestedObject struct {
	Attributes []DataSourceAttribute `json:"attributes,omitempty"`
	Blocks     []DataSourceBlock     `json:"blocks,omitempty"`
	CustomType *CustomType           `json:"custom_type,omitempty"`
	Validators []Validators          `json:"validators,omitempty"`
}

type Provider struct {
	Name   string          `json:"name"`
	Schema *ProviderSchema `json:"schema,omitempty"`
}

type ProviderSchema struct {
	Attributes []ProviderAttribute `json:"attributes,omitempty"`
	Blocks     []ProviderBlock     `json:"blocks"`
}

type ProviderAttribute struct {
}

type ProviderBlock struct {
}

type Resource struct {
	Name   string         `json:"name"`
	Schema ResourceSchema `json:"schema"`
}

type ResourceSchema struct {
	Attributes []ResourceAttribute `json:"attributes,omitempty"`
	Blocks     []ResourceBlock     `json:"blocks,omitempty"`
}

type ResourceAttribute struct {
	Name string                 `json:"name"`
	Bool *ResourceBoolAttribute `json:"bool,omitempty"`
}

type ResourceBoolAttribute struct {
	ComputedOptionalRequired ComputedOptionalRequired `json:"computed_optional_required"`
	CustomType               *CustomType              `json:"custom_type,omitempty"`
	Default                  *BoolDefault             `json:"default,omitempty"`
	DeprecationMessage       *string                  `json:"deprecation_message,omitempty"`
	Description              *string                  `json:"description,omitempty"`
	PlanModifiers            []PlanModifier           `json:"plan_modifiers,omitempty"`
	Sensitive                *bool                    `json:"sensitive,omitempty"`
	Validators               []Validators             `json:"validators,omitempty"`
}

type ResourceBlock struct {
}

type Validators struct {
	Custom *CustomValidator `json:"custom,omitempty"`
}

type CustomValidator struct {
	Import           string `json:"import"`
	SchemaDefinition string `json:"schema_definition"`
}

type BoolDefault struct {
	Custom *CustomDefault `json:"custom,omitempty"`
	Static *bool          `json:"static"`
}

type CustomDefault struct {
	Import           string `json:"import"`
	SchemaDefinition string `json:"schema_definition"`
}

type PlanModifier struct {
	Custom             *CustomPlanModifier `json:"custom,omitempty"`
	RequiresReplace    *RequiresReplace    `json:"requires_replace,omitempty"`
	UseStateForUnknown *UseStateForUnknown `json:"use_state_for_unknown,omitempty"`
}

type CustomPlanModifier struct {
	Import           string `json:"import"`
	SchemaDefinition string `json:"schema_definition"`
}

type RequiresReplace struct {
}

type UseStateForUnknown struct {
}

type ComputedOptionalRequired string

const (
	Computed         ComputedOptionalRequired = "computed"
	ComputedOptional ComputedOptionalRequired = "computed_optional"
	Optional         ComputedOptionalRequired = "optional"
	Required         ComputedOptionalRequired = "required"
)
