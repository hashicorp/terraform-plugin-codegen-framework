package validate

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
)

func TestNames(t *testing.T) {
	testCases := map[string]struct {
		input         spec.Specification
		expectedError error
	}{
		"data-source-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
					},
					{
						Name: "example",
					},
				},
			},
			expectedError: fmt.Errorf(`data source name "example" is duplicated`),
		},
		"data-source-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
					},
					{
						Name: "different",
					},
				},
			},
		},
		"data-source-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "first_attr",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute name "first_attr" is duplicated`),
		},
		"data-source-attribute-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "second_attr",
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "first_block",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block name "first_block" is duplicated`),
		},
		"data-source-attribute-and-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first",
								},
							},
							Blocks: []datasource.Block{
								{
									Name: "first",
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "second_block",
								},
							},
						},
					},
				},
			},
		},
		"data-source-list-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									ListNested: &datasource.ListNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"data-source-list-nested-attribute-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									ListNested: &datasource.ListNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-list-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									ListNested: &datasource.ListNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-map-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									MapNested: &datasource.MapNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"data-source-map-nested-attribute-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									MapNested: &datasource.MapNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-map-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									MapNested: &datasource.MapNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-set-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SetNested: &datasource.SetNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"data-source-set-nested-attribute-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SetNested: &datasource.SetNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-set-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SetNested: &datasource.SetNestedAttribute{
										NestedObject: datasource.NestedAttributeObject{
											Attributes: []datasource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-single-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "nested_attr",
											},
											{
												Name: "nested_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"data-source-single-nested-attribute-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "nested_first_attr",
											},
											{
												Name: "nested_second_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-attribute-and-single-nested-attribute-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &datasource.SingleNestedAttribute{
										Attributes: []datasource.Attribute{
											{
												Name: "first_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-list-nested-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block "first_block" block name "nested_block" is duplicated`),
		},
		"data-source-list-nested-block-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-and-list-nested-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									ListNested: &datasource.ListNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-set-nested-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block "first_block" block name "nested_block" is duplicated`),
		},
		"data-source-set-nested-block-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-and-set-nested-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SetNested: &datasource.SetNestedBlock{
										NestedObject: datasource.NestedBlockObject{
											Blocks: []datasource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-single-nested-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "nested_block",
											},
											{
												Name: "nested_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" block "first_block" block name "nested_block" is duplicated`),
		},
		"data-source-single-nested-block-names-unique": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "nested_first_block",
											},
											{
												Name: "nested_second_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-block-and-single-nested-block-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Blocks: []datasource.Block{
								{
									Name: "first_block",
									SingleNested: &datasource.SingleNestedBlock{
										Blocks: []datasource.Block{
											{
												Name: "first_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"data-source-object-attribute-type-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
											},
											{
												Name: "obj_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" object attribute type name "obj_attr" is duplicated`),
		},
		"data-source-object-object-attribute-type-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "first_attr",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: []schema.ObjectAttributeType{
													{
														Name: "obj_obj_attr",
													},
													{
														Name: "obj_obj_attr",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`data source "example" attribute "first_attr" object attribute type "obj_attr" object attribute type name "obj_obj_attr" is duplicated`),
		},
		"data-source-object-and-object-object-attribute-type-names-duplicated": {
			input: spec.Specification{
				DataSources: []datasource.DataSource{
					{
						Name: "example",
						Schema: &datasource.Schema{
							Attributes: []datasource.Attribute{
								{
									Name: "obj_attr",
									Object: &datasource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: []schema.ObjectAttributeType{
													{
														Name: "obj_attr",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
					},
					{
						Name: "example",
					},
				},
			},
			expectedError: fmt.Errorf(`resource name "example" is duplicated`),
		},
		"resource-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
					},
					{
						Name: "different",
					},
				},
			},
		},
		"resource-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "first_attr",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute name "first_attr" is duplicated`),
		},
		"resource-attribute-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
								},
								{
									Name: "second_attr",
								},
							},
						},
					},
				},
			},
		},
		"resource-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "first_block",
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block name "first_block" is duplicated`),
		},
		"resource-attribute-and-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first",
								},
							},
							Blocks: []resource.Block{
								{
									Name: "first",
								},
							},
						},
					},
				},
			},
		},
		"resource-block-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
								},
								{
									Name: "second_block",
								},
							},
						},
					},
				},
			},
		},
		"resource-list-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									ListNested: &resource.ListNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"resource-list-nested-attribute-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									ListNested: &resource.ListNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-list-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									ListNested: &resource.ListNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-map-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									MapNested: &resource.MapNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"resource-map-nested-attribute-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									MapNested: &resource.MapNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-map-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									MapNested: &resource.MapNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-set-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SetNested: &resource.SetNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_attr",
												},
												{
													Name: "nested_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"resource-set-nested-attribute-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SetNested: &resource.SetNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "nested_first_attr",
												},
												{
													Name: "nested_second_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-set-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SetNested: &resource.SetNestedAttribute{
										NestedObject: resource.NestedAttributeObject{
											Attributes: []resource.Attribute{
												{
													Name: "first_attr",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-single-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "nested_attr",
											},
											{
												Name: "nested_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" attribute name "nested_attr" is duplicated`),
		},
		"resource-single-nested-attribute-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "nested_first_attr",
											},
											{
												Name: "nested_second_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-attribute-and-single-nested-attribute-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									SingleNested: &resource.SingleNestedAttribute{
										Attributes: []resource.Attribute{
											{
												Name: "first_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-list-nested-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block "first_block" block name "nested_block" is duplicated`),
		},
		"resource-list-nested-block-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-block-and-list-nested-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									ListNested: &resource.ListNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-set-nested-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_block",
												},
												{
													Name: "nested_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block "first_block" block name "nested_block" is duplicated`),
		},
		"resource-set-nested-block-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "nested_first_block",
												},
												{
													Name: "nested_second_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-block-and-set-nested-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SetNested: &resource.SetNestedBlock{
										NestedObject: resource.NestedBlockObject{
											Blocks: []resource.Block{
												{
													Name: "first_block",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-single-nested-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "nested_block",
											},
											{
												Name: "nested_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" block "first_block" block name "nested_block" is duplicated`),
		},
		"resource-single-nested-block-names-unique": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "nested_first_block",
											},
											{
												Name: "nested_second_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-block-and-single-nested-block-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Blocks: []resource.Block{
								{
									Name: "first_block",
									SingleNested: &resource.SingleNestedBlock{
										Blocks: []resource.Block{
											{
												Name: "first_block",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"resource-object-attribute-type-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
											},
											{
												Name: "obj_attr",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" object attribute type name "obj_attr" is duplicated`),
		},
		"resource-object-object-attribute-type-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "first_attr",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: []schema.ObjectAttributeType{
													{
														Name: "obj_obj_attr",
													},
													{
														Name: "obj_obj_attr",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedError: fmt.Errorf(`resource "example" attribute "first_attr" object attribute type "obj_attr" object attribute type name "obj_obj_attr" is duplicated`),
		},
		"resource-object-and-object-object-attribute-type-names-duplicated": {
			input: spec.Specification{
				Resources: []resource.Resource{
					{
						Name: "example",
						Schema: &resource.Schema{
							Attributes: []resource.Attribute{
								{
									Name: "obj_attr",
									Object: &resource.ObjectAttribute{
										AttributeTypes: []schema.ObjectAttributeType{
											{
												Name: "obj_attr",
												Object: []schema.ObjectAttributeType{
													{
														Name: "obj_attr",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := SchemaNames(context.Background(), testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}
		})
	}
}

var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}

	return x.Error() == y.Error()
})
