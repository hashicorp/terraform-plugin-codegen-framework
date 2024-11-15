package util

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-spec/resource"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-spec/spec"
)

type RequestWithDTOName struct {
	Create spec.RequestType       `json:"create,omitempty"`
	Read   RequestTypeWithDTOName `json:"read"`
	Update []*spec.RequestType    `json:"update"`
	Delete spec.RequestType       `json:"delete"`
}

type RequestTypeWithDTOName struct {
	spec.RequestType
	Response string `json:"response"`
}

type RequestWithResponse struct {
	Requests []spec.Request
}

type CodeSpec struct {
	Provider    map[string]interface{} `json:"provider"`
	Resources   []Resource             `json:"resources"`
	DataSources []Resource             `json:"datasources"`
	Requests    []RequestWithDTOName   `json:"requests"`
	Version     string                 `json:"version"`
}

type Resource struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
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

func ExtractAttribute(file string) (resource.Attributes, string, string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var result CodeSpec
	if err := json.Unmarshal(byteValue, &result); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	resources := result.Resources
	if len(resources) == 0 {
		log.Fatal("'resources' slice is empty")
	}

	// TODO - 여러 리소스 처리하기
	resourceName := resources[0].Name
	dataSourceName := result.DataSources[0].Name
	rawAttributes := resources[0].Schema.Attributes

	return rawAttributes, resourceName, dataSourceName
}

func ExtractRequest(file string) RequestWithDTOName {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open JSON req file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var result CodeSpec
	if err := json.Unmarshal(byteValue, &result); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if len(result.Requests) == 0 {
		log.Fatalf("No requests found in the JSON file")
	}

	return result.Requests[0]
}
