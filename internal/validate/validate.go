package validate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
)

// IntermediateRepresentationValidator is a struct used to abstract away the details
// of the underlying validations.
type IntermediateRepresentationValidator struct {
}

// NewIntermediateRepresentationValidator returns an IntermediateRepresentationValidator struct.
func NewIntermediateRepresentationValidator() IntermediateRepresentationValidator {
	return IntermediateRepresentationValidator{}
}

// Validate validates that the input is valid JSON and that the input adheres to the JSON
// schema within the terraform-plugin-codegen-spec module.
func (v IntermediateRepresentationValidator) Validate(ctx context.Context, input []byte) error {
	err := v.validateJSON(input)
	if err != nil {
		return err
	}

	err = v.validateWithSchema(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

// validateJSON verifies that the supplied input is valid JSON.
func (v IntermediateRepresentationValidator) validateJSON(input []byte) error {
	if !json.Valid(input) {
		return errors.New("invalid JSON")
	}

	return nil
}

// validateWithSchema uses the spec.Validate function to verify that the input adheres to the JSON schema
// defined within the spec modules.
func (v IntermediateRepresentationValidator) validateWithSchema(ctx context.Context, input []byte) error {
	return spec.Validate(ctx, input)
}

// SpecValidator is a struct used to abstract away the details of the underlying validations.
type SpecValidator struct {
}

// NewSpecValidator returns a SpecValidator struct.
func NewSpecValidator() SpecValidator {
	return SpecValidator{}
}

// Validate validates that the input is valid JSON and that the input adheres to the JSON
// schema within the terraform-plugin-codegen-spec module.
func (v SpecValidator) Validate(ctx context.Context, input spec.Specification) error {
	return v.validateSchemaNames(ctx, input)
}

// validateSchemaNames determines whether any of the names used for data sources or
// resources are duplicated. SchemaNames also determines whether any of the names used
// for attributes or blocks at the same level of nesting within a schema are duplicated.
func (v SpecValidator) validateSchemaNames(ctx context.Context, s spec.Specification) error {
	duplicates := make(map[string]struct{})
	dataSourceNames := make(map[string]struct{})

	for _, v := range s.DataSources {
		if _, ok := dataSourceNames[v.Name]; ok {
			duplicates[fmt.Sprintf("data source name %q is duplicated", v.Name)] = struct{}{}
		}

		dataSourceNames[v.Name] = struct{}{}

		if v.Schema != nil {
			duplicateAttributes := duplicatedDatasourceAttributeNames(fmt.Sprintf("data source %q", v.Name), v.Schema.Attributes)

			for k, v := range duplicateAttributes {
				duplicates[k] = v
			}

			duplicateBlocks := duplicatedDatasourceBlockNames(fmt.Sprintf("data source %q", v.Name), v.Schema.Blocks)

			for k, v := range duplicateBlocks {
				duplicates[k] = v
			}
		}
	}

	resourceNames := make(map[string]struct{})

	for _, v := range s.Resources {
		if _, ok := resourceNames[v.Name]; ok {
			duplicates[fmt.Sprintf("resource name %q is duplicated", v.Name)] = struct{}{}
		}

		resourceNames[v.Name] = struct{}{}

		if v.Schema != nil {
			duplicateAttributes := duplicatedResourceAttributeNames(fmt.Sprintf("resource %q", v.Name), v.Schema.Attributes)

			for k, v := range duplicateAttributes {
				duplicates[k] = v
			}

			duplicateBlocks := duplicatedResourceBlockNames(fmt.Sprintf("resource %q", v.Name), v.Schema.Blocks)

			for k, v := range duplicateBlocks {
				duplicates[k] = v
			}
		}
	}

	if len(duplicates) > 0 {
		var dups []string

		for k := range duplicates {
			dups = append(dups, k)
		}

		return fmt.Errorf(strings.Join(dups, "\n"))
	}

	return nil
}

func duplicatedDatasourceAttributeNames(name string, a []datasource.Attribute) map[string]struct{} {
	duplicates := make(map[string]struct{})
	attrNames := make(map[string]struct{})

	for _, v := range a {
		if _, ok := attrNames[v.Name]; ok {
			duplicates[fmt.Sprintf("%s attribute name %q is duplicated", name, v.Name)] = struct{}{}
		}

		attrNames[v.Name] = struct{}{}

		attrName := fmt.Sprintf("%s attribute %q", name, v.Name)
		var nestedDuplicates map[string]struct{}

		switch {
		case v.ListNested != nil:
			nestedDuplicates = duplicatedDatasourceAttributeNames(attrName, v.ListNested.NestedObject.Attributes)
		case v.MapNested != nil:
			nestedDuplicates = duplicatedDatasourceAttributeNames(attrName, v.MapNested.NestedObject.Attributes)
		case v.Object != nil:
			nestedDuplicates = duplicatedObjectAttributeTypeNames(attrName, v.Object.AttributeTypes)
		case v.SetNested != nil:
			nestedDuplicates = duplicatedDatasourceAttributeNames(attrName, v.SetNested.NestedObject.Attributes)
		case v.SingleNested != nil:
			nestedDuplicates = duplicatedDatasourceAttributeNames(attrName, v.SingleNested.Attributes)
		}

		for k, v := range nestedDuplicates {
			duplicates[k] = v
		}
	}

	return duplicates
}

func duplicatedDatasourceBlockNames(name string, b []datasource.Block) map[string]struct{} {
	duplicates := make(map[string]struct{})
	blockNames := make(map[string]struct{})

	for _, v := range b {
		if _, ok := blockNames[v.Name]; ok {
			duplicates[fmt.Sprintf("%s block name %q is duplicated", name, v.Name)] = struct{}{}
		}

		blockNames[v.Name] = struct{}{}

		blockName := fmt.Sprintf("%s block %q", name, v.Name)
		var nestedBlockDuplicates map[string]struct{}

		switch {
		case v.ListNested != nil:
			nestedBlockDuplicates = duplicatedDatasourceBlockNames(blockName, v.ListNested.NestedObject.Blocks)
		case v.SetNested != nil:
			nestedBlockDuplicates = duplicatedDatasourceBlockNames(blockName, v.SetNested.NestedObject.Blocks)
		case v.SingleNested != nil:
			nestedBlockDuplicates = duplicatedDatasourceBlockNames(blockName, v.SingleNested.Blocks)
		}

		for k, v := range nestedBlockDuplicates {
			duplicates[k] = v
		}
	}

	return duplicates
}

func duplicatedResourceAttributeNames(name string, a []resource.Attribute) map[string]struct{} {
	duplicates := make(map[string]struct{})
	attrNames := make(map[string]struct{})

	for _, v := range a {
		if _, ok := attrNames[v.Name]; ok {
			duplicates[fmt.Sprintf("%s attribute name %q is duplicated", name, v.Name)] = struct{}{}
		}

		attrNames[v.Name] = struct{}{}

		attrName := fmt.Sprintf("%s attribute %q", name, v.Name)
		var nestedDuplicates map[string]struct{}

		switch {
		case v.ListNested != nil:
			nestedDuplicates = duplicatedResourceAttributeNames(attrName, v.ListNested.NestedObject.Attributes)
		case v.MapNested != nil:
			nestedDuplicates = duplicatedResourceAttributeNames(attrName, v.MapNested.NestedObject.Attributes)
		case v.Object != nil:
			nestedDuplicates = duplicatedObjectAttributeTypeNames(attrName, v.Object.AttributeTypes)
		case v.SetNested != nil:
			nestedDuplicates = duplicatedResourceAttributeNames(attrName, v.SetNested.NestedObject.Attributes)
		case v.SingleNested != nil:
			nestedDuplicates = duplicatedResourceAttributeNames(attrName, v.SingleNested.Attributes)
		}

		for k, v := range nestedDuplicates {
			duplicates[k] = v
		}
	}

	return duplicates
}

func duplicatedResourceBlockNames(name string, b []resource.Block) map[string]struct{} {
	duplicates := make(map[string]struct{})
	blockNames := make(map[string]struct{})

	for _, v := range b {
		if _, ok := blockNames[v.Name]; ok {
			duplicates[fmt.Sprintf("%s block name %q is duplicated", name, v.Name)] = struct{}{}
		}

		blockNames[v.Name] = struct{}{}

		blockName := fmt.Sprintf("%s block %q", name, v.Name)
		var nestedBlockDuplicates map[string]struct{}

		switch {
		case v.ListNested != nil:
			nestedBlockDuplicates = duplicatedResourceBlockNames(blockName, v.ListNested.NestedObject.Blocks)
		case v.SetNested != nil:
			nestedBlockDuplicates = duplicatedResourceBlockNames(blockName, v.SetNested.NestedObject.Blocks)
		case v.SingleNested != nil:
			nestedBlockDuplicates = duplicatedResourceBlockNames(blockName, v.SingleNested.Blocks)
		}

		for k, v := range nestedBlockDuplicates {
			duplicates[k] = v
		}
	}

	return duplicates
}

func duplicatedObjectAttributeTypeNames(name string, o []schema.ObjectAttributeType) map[string]struct{} {
	duplicates := make(map[string]struct{})
	attrTypeNames := make(map[string]struct{})

	for _, v := range o {
		if _, ok := attrTypeNames[v.Name]; ok {
			duplicates[fmt.Sprintf("%s object attribute type name %q is duplicated", name, v.Name)] = struct{}{}
		}

		attrTypeNames[v.Name] = struct{}{}

		attrTypeName := fmt.Sprintf("%s object attribute type %q", name, v.Name)
		var nestedAttrTypeDuplicates map[string]struct{}

		if v.Object != nil {
			nestedAttrTypeDuplicates = duplicatedObjectAttributeTypeNames(attrTypeName, v.Object)
		}

		for k, v := range nestedAttrTypeDuplicates {
			duplicates[k] = v
		}
	}

	return duplicates
}
