package util

import (
	"encoding/json"
	"fmt"
	"os"

	high "github.com/pb33f/libopenapi/datamodel/high/v3"

	oas "github.com/getkin/kin-openapi/openapi3"
	docs "github.com/go-oas/docs"
)

type OpenAPI struct {
	Paths      map[string]*PathItem `json:"paths"`
	Info       *Info                `json:"info"`
	Components *Components          `json:"components"`
}

type Info struct {
	Name string `json:"name"`
}

type PathItem struct {
	Post *high.Operation `json:"post,omitempty"`
	Get  *high.Operation `json:"get,omitempty"`
}

type Components struct {
	Schemas map[string]OASchema `json:"schemas"`
}

type OASchema struct {
	Properties map[string]Property `json:"properties"`
}

type Property struct {
	docs.SchemaProperty
	Type                     string              `json:"type"`
	Format                   string              `json:"format,omitempty"`
	ComputedOptionalRequired string              `json:"computed_optional_required,omitempty"`
	Items                    Items               `json:"items,omitempty"`
	Properties               map[string]Property `json:"properties"`
}

type Items struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
}

func ExtractResourceName(jsonFilePath string) (string, error) {
	byteValue, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	var openAPI OpenAPI
	if err := json.Unmarshal(byteValue, &openAPI); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return openAPI.Info.Name, nil
}

func ExtractDto(jsonFilePath, dtoName string) ([]byte, error) {
	var openAPI oas.T

	byteValue, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	err = json.Unmarshal(byteValue, &openAPI)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	schema := openAPI.Components.Schemas[dtoName]
	jsonBytes, err := json.Marshal(schema)
	if err != nil {
		return nil, fmt.Errorf("error marshalling schema to JSON: %w", err)
	}

	return jsonBytes, nil
}
