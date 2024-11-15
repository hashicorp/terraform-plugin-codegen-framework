package resource_api_key

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/common"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/conn"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

func ApiKeyResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"api_key_description": schema.StringAttribute{
						Computed:            true,
						Description:         "Api Key Description",
						MarkdownDescription: "Api Key Description",
					},
					"api_key_id": schema.StringAttribute{
						Computed:            true,
						Description:         "Api Key Id",
						MarkdownDescription: "Api Key Id",
					},
					"api_key_name": schema.StringAttribute{
						Computed:            true,
						Description:         "Api Key Name",
						MarkdownDescription: "Api Key Name",
					},
					"domain_code": schema.StringAttribute{
						Computed:            true,
						Description:         "Domain Code",
						MarkdownDescription: "Domain Code",
					},
					"is_enabled": schema.BoolAttribute{
						Computed:            true,
						Description:         "Is Enabled",
						MarkdownDescription: "Is Enabled",
					},
					"mod_time": schema.StringAttribute{
						Computed:            true,
						Description:         "Mod Time",
						MarkdownDescription: "Mod Time",
					},
					"modifier": schema.StringAttribute{
						Computed:            true,
						Description:         "Modifier",
						MarkdownDescription: "Modifier",
					},
					"primary_key": schema.StringAttribute{
						Computed:            true,
						Description:         "Primary Key",
						MarkdownDescription: "Primary Key",
					},
					"secondary_key": schema.StringAttribute{
						Computed:            true,
						Description:         "Secondary Key",
						MarkdownDescription: "Secondary Key",
					},
					"tenant_id": schema.StringAttribute{
						Computed:            true,
						Description:         "Tenant Id",
						MarkdownDescription: "Tenant Id",
					},
				},
				CustomType: ApiKeyType{
					ObjectType: types.ObjectType{
						AttrTypes: ApiKeyValue{}.AttributeTypes(ctx),
					},
				},
				Computed: true,
			},
			"api_key_description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Api Key Description<br>Length(Min/Max): 0/50",
				MarkdownDescription: "Api Key Description<br>Length(Min/Max): 0/50",
			},
			"api_key_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Api Key Id",
				MarkdownDescription: "Api Key Id",
			},
			"api_key_name": schema.StringAttribute{
				Required:            true,
				Description:         "Api Key Name<br>Length(Min/Max): 0/20",
				MarkdownDescription: "Api Key Name<br>Length(Min/Max): 0/20",
			},
			"apikeyid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "api-key-id",
				MarkdownDescription: "api-key-id",
			},
			"domain_code": schema.StringAttribute{
				Computed:            true,
				Description:         "Domain Code",
				MarkdownDescription: "Domain Code",
			},
			"is_enabled": schema.BoolAttribute{
				Required:            true,
				Description:         "Is Enabled",
				MarkdownDescription: "Is Enabled",
			},
			"mod_time": schema.StringAttribute{
				Computed:            true,
				Description:         "Mod Time",
				MarkdownDescription: "Mod Time",
			},
			"modifier": schema.StringAttribute{
				Computed:            true,
				Description:         "Modifier",
				MarkdownDescription: "Modifier",
			},
			"primary_key": schema.StringAttribute{
				Computed:            true,
				Description:         "Primary Key",
				MarkdownDescription: "Primary Key",
			},
			"secondary_key": schema.StringAttribute{
				Computed:            true,
				Description:         "Secondary Key",
				MarkdownDescription: "Secondary Key",
			},
			"tenant_id": schema.StringAttribute{
				Computed:            true,
				Description:         "Tenant Id",
				MarkdownDescription: "Tenant Id",
			},
		},
	}
}

func NewApiKeyResource() resource.Resource {
	return &apiKeyResource{}
}

type apiKeyResource struct {
	config *conn.ProviderConfig
}

func (a *apiKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*conn.ProviderConfig)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ProviderConfig, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	a.config = config
}

func (a *apiKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (a *apiKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan ApikeydtoModel

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "PUT", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/"+plan.ApiKeyId.String(),
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
		)
	}

	response, err := util.Request(execFunc, "")
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("DELETING ERROR", "response invalid")
		return
	}

	err = waitResourceDeleted(ctx)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "DeleteApiKey response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *apiKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_key" //임의로 space 넣어줘야 함 (_)가 필요하기 때문
}

func (a *apiKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan ApikeydtoModel

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *apiKeyResource) Schema(context.Context, resource.SchemaRequest, *resource.SchemaResponse) {
	panic("unimplemented") // WIP
}

func (a *apiKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ApikeydtoModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := fmt.Sprintf(`{
		"isEnabled": %[1]s,
		"apiKeyDescription": %[2]s,
		"apiKeyName": %[3]s
	}`, plan.IsEnabled, plan.ApiKeyDescription, plan.ApiKeyName) // request에만 들어가있는 애들의 경우 dto에 없으므로 추가 수정 필요.

	tflog.Info(ctx, "UpdateApiKey reqParams="+reqBody)

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "PATCH", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/"+plan.ApiKeyId.String(),
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-d", reqBody,
		)
	}

	response, err := util.Request(execFunc, reqBody)
	if err != nil {
		resp.Diagnostics.AddError("UPDATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("UPDATING ERROR", "response invalid")
		return
	}

	tflog.Info(ctx, "UpdateApiKey response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *apiKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApikeydtoModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody := fmt.Sprintf(`{
			"apiKeyDescription": %[1]s,
			"apiKeyName": %[2]s
		}`, plan.ApiKeyDescription, plan.ApiKeyName) // 마찬가지로 수기 필요.

	tflog.Info(ctx, "CreateApiKey reqParams="+reqBody)

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "PUT", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/",
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-d", reqBody,
		)
	}

	response, err := util.Request(execFunc, reqBody)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("CREATING ERROR", "response invalid")
		return
	}

	err = waitResourceCreated(ctx)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "CreateApiKey response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

type ApikeydtoModel struct {
	ApiKeyDescription types.String `json:"api_key_description"`
	DomainCode        types.String `json:"domain_code"`
	IsEnabled         types.Bool   `json:"is_enabled"`
	ModTime           types.String `json:"mod_time"`
	Modifier          types.String `json:"modifier"`
	ApiKeyId          types.String `json:"api_key_id"`
	ApiKeyName        types.String `json:"api_key_name"`
	PrimaryKey        types.String `json:"primary_key"`
	SecondaryKey      types.String `json:"secondary_key"`
	TenantId          types.String `json:"tenant_id"`
}

func ConvertToFrameworkTypes(data map[string]interface{}) (*ApikeydtoModel, error) {
	var dto ApikeydtoModel

	dto.ApiKeyDescription = types.StringValue(data["api_key_description"].(string))
	dto.DomainCode = types.StringValue(data["domain_code"].(string))
	dto.IsEnabled = types.BoolValue(data["is_enabled"].(bool))
	dto.ModTime = types.StringValue(data["mod_time"].(string))
	dto.Modifier = types.StringValue(data["modifier"].(string))
	dto.ApiKeyId = types.StringValue(data["api_key_id"].(string))
	dto.ApiKeyName = types.StringValue(data["api_key_name"].(string))
	dto.PrimaryKey = types.StringValue(data["primary_key"].(string))
	dto.SecondaryKey = types.StringValue(data["secondary_key"].(string))
	dto.TenantId = types.StringValue(data["tenant_id"].(string))

	return &dto, nil
}

func waitResourceCreated(ctx context.Context) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATING"},
		Target:  []string{"CREATED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
				return exec.Command("curl", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/",
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
				)
			}

			response, err := util.Request(getExecFunc, "")
			if err != nil {
				return response, "CREATING", nil
			}
			if response == nil {
				return response, "CREATED", nil
			}

			return response, "CREATING", nil
		},
		Timeout:    conn.DefaultTimeout,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("error occured while waiting for resource to be created: %s", err)
	}
	return nil
}

func waitResourceDeleted(ctx context.Context) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"DELETING"},
		Target:  []string{"DELETED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
				return exec.Command("curl", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/",
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
				)
			}

			response, err := util.Request(getExecFunc, "")
			if response != nil {
				return response, "DELETING", nil
			}

			if err != nil {
				return response, "DELETED", nil
			}

			return response, "DELETED", nil
		},
		Timeout:    conn.DefaultTimeout,
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("error occured while waiting for resource to be deleted: %s", err)
	}
	return nil
}

func diagOff[V, T interface{}](input func(ctx context.Context, elementType T, elements any) (V, diag.Diagnostics), ctx context.Context, elementType T, elements any) V {
	var emptyReturn V

	v, diags := input(ctx, elementType, elements)

	if diags.HasError() {
		diags.AddError("REFRESING ERROR", "invalid diagOff operation")
		return emptyReturn
	}

	return v
}

func getAndRefresh(diagnostics diag.Diagnostics) *ApikeydtoModel {
	getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/",
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
		)
	}

	response, err := util.Request(getExecFunc, "")
	if err != nil {
		diagnostics.AddError("UPDATING ERROR", err.Error())
		return nil
	}
	if response == nil {
		diagnostics.AddError("UPDATING ERROR", "response invalid")
		return nil
	}
	// WIP
	// err = waitBucketCreated(ctx, o.config, plan.BucketName.String())
	// if err != nil {
	// 	resp.Diagnostics.AddError("CREATING ERROR", err.Error())
	// 	return
	// }
	newPlan, err := ConvertToFrameworkTypes(response)
	if err != nil {
		diagnostics.AddError("CREATING ERROR", err.Error())
		return nil
	}

	return newPlan
}

type ApiKeyModel struct {
	ApiKey            ApiKeyValue  `tfsdk:"api_key"`
	ApiKeyDescription types.String `tfsdk:"api_key_description"`
	ApiKeyId          types.String `tfsdk:"api_key_id"`
	ApiKeyName        types.String `tfsdk:"api_key_name"`
	Apikeyid          types.String `tfsdk:"apikeyid"`
	DomainCode        types.String `tfsdk:"domain_code"`
	IsEnabled         types.Bool   `tfsdk:"is_enabled"`
	ModTime           types.String `tfsdk:"mod_time"`
	Modifier          types.String `tfsdk:"modifier"`
	PrimaryKey        types.String `tfsdk:"primary_key"`
	SecondaryKey      types.String `tfsdk:"secondary_key"`
	TenantId          types.String `tfsdk:"tenant_id"`
}

type ApiKeyType struct {
	basetypes.ObjectType
}

func (t ApiKeyType) Equal(o attr.Type) bool {
	other, ok := o.(ApiKeyType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ApiKeyType) String() string {
	return "ApiKeyType"
}

func (t ApiKeyType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	apiKeyDescriptionAttribute, ok := attributes["api_key_description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_key_description is missing from object`)

		return nil, diags
	}

	apiKeyDescriptionVal, ok := apiKeyDescriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_key_description expected to be basetypes.StringValue, was: %T`, apiKeyDescriptionAttribute))
	}

	apiKeyIdAttribute, ok := attributes["api_key_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_key_id is missing from object`)

		return nil, diags
	}

	apiKeyIdVal, ok := apiKeyIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_key_id expected to be basetypes.StringValue, was: %T`, apiKeyIdAttribute))
	}

	apiKeyNameAttribute, ok := attributes["api_key_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_key_name is missing from object`)

		return nil, diags
	}

	apiKeyNameVal, ok := apiKeyNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_key_name expected to be basetypes.StringValue, was: %T`, apiKeyNameAttribute))
	}

	domainCodeAttribute, ok := attributes["domain_code"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`domain_code is missing from object`)

		return nil, diags
	}

	domainCodeVal, ok := domainCodeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`domain_code expected to be basetypes.StringValue, was: %T`, domainCodeAttribute))
	}

	isEnabledAttribute, ok := attributes["is_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`is_enabled is missing from object`)

		return nil, diags
	}

	isEnabledVal, ok := isEnabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`is_enabled expected to be basetypes.BoolValue, was: %T`, isEnabledAttribute))
	}

	modTimeAttribute, ok := attributes["mod_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mod_time is missing from object`)

		return nil, diags
	}

	modTimeVal, ok := modTimeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mod_time expected to be basetypes.StringValue, was: %T`, modTimeAttribute))
	}

	modifierAttribute, ok := attributes["modifier"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`modifier is missing from object`)

		return nil, diags
	}

	modifierVal, ok := modifierAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`modifier expected to be basetypes.StringValue, was: %T`, modifierAttribute))
	}

	primaryKeyAttribute, ok := attributes["primary_key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`primary_key is missing from object`)

		return nil, diags
	}

	primaryKeyVal, ok := primaryKeyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`primary_key expected to be basetypes.StringValue, was: %T`, primaryKeyAttribute))
	}

	secondaryKeyAttribute, ok := attributes["secondary_key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`secondary_key is missing from object`)

		return nil, diags
	}

	secondaryKeyVal, ok := secondaryKeyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`secondary_key expected to be basetypes.StringValue, was: %T`, secondaryKeyAttribute))
	}

	tenantIdAttribute, ok := attributes["tenant_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tenant_id is missing from object`)

		return nil, diags
	}

	tenantIdVal, ok := tenantIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tenant_id expected to be basetypes.StringValue, was: %T`, tenantIdAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ApiKeyValue{
		ApiKeyDescription: apiKeyDescriptionVal,
		ApiKeyId:          apiKeyIdVal,
		ApiKeyName:        apiKeyNameVal,
		DomainCode:        domainCodeVal,
		IsEnabled:         isEnabledVal,
		ModTime:           modTimeVal,
		Modifier:          modifierVal,
		PrimaryKey:        primaryKeyVal,
		SecondaryKey:      secondaryKeyVal,
		TenantId:          tenantIdVal,
		state:             attr.ValueStateKnown,
	}, diags
}

func NewApiKeyValueNull() ApiKeyValue {
	return ApiKeyValue{
		state: attr.ValueStateNull,
	}
}

func NewApiKeyValueUnknown() ApiKeyValue {
	return ApiKeyValue{
		state: attr.ValueStateUnknown,
	}
}

func NewApiKeyValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ApiKeyValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ApiKeyValue Attribute Value",
				"While creating a ApiKeyValue value, a missing attribute value was detected. "+
					"A ApiKeyValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ApiKeyValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ApiKeyValue Attribute Type",
				"While creating a ApiKeyValue value, an invalid attribute value was detected. "+
					"A ApiKeyValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ApiKeyValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ApiKeyValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ApiKeyValue Attribute Value",
				"While creating a ApiKeyValue value, an extra attribute value was detected. "+
					"A ApiKeyValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ApiKeyValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewApiKeyValueUnknown(), diags
	}

	apiKeyDescriptionAttribute, ok := attributes["api_key_description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_key_description is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	apiKeyDescriptionVal, ok := apiKeyDescriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_key_description expected to be basetypes.StringValue, was: %T`, apiKeyDescriptionAttribute))
	}

	apiKeyIdAttribute, ok := attributes["api_key_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_key_id is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	apiKeyIdVal, ok := apiKeyIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_key_id expected to be basetypes.StringValue, was: %T`, apiKeyIdAttribute))
	}

	apiKeyNameAttribute, ok := attributes["api_key_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`api_key_name is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	apiKeyNameVal, ok := apiKeyNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`api_key_name expected to be basetypes.StringValue, was: %T`, apiKeyNameAttribute))
	}

	domainCodeAttribute, ok := attributes["domain_code"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`domain_code is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	domainCodeVal, ok := domainCodeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`domain_code expected to be basetypes.StringValue, was: %T`, domainCodeAttribute))
	}

	isEnabledAttribute, ok := attributes["is_enabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`is_enabled is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	isEnabledVal, ok := isEnabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`is_enabled expected to be basetypes.BoolValue, was: %T`, isEnabledAttribute))
	}

	modTimeAttribute, ok := attributes["mod_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mod_time is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	modTimeVal, ok := modTimeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mod_time expected to be basetypes.StringValue, was: %T`, modTimeAttribute))
	}

	modifierAttribute, ok := attributes["modifier"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`modifier is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	modifierVal, ok := modifierAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`modifier expected to be basetypes.StringValue, was: %T`, modifierAttribute))
	}

	primaryKeyAttribute, ok := attributes["primary_key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`primary_key is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	primaryKeyVal, ok := primaryKeyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`primary_key expected to be basetypes.StringValue, was: %T`, primaryKeyAttribute))
	}

	secondaryKeyAttribute, ok := attributes["secondary_key"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`secondary_key is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	secondaryKeyVal, ok := secondaryKeyAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`secondary_key expected to be basetypes.StringValue, was: %T`, secondaryKeyAttribute))
	}

	tenantIdAttribute, ok := attributes["tenant_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tenant_id is missing from object`)

		return NewApiKeyValueUnknown(), diags
	}

	tenantIdVal, ok := tenantIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tenant_id expected to be basetypes.StringValue, was: %T`, tenantIdAttribute))
	}

	if diags.HasError() {
		return NewApiKeyValueUnknown(), diags
	}

	return ApiKeyValue{
		ApiKeyDescription: apiKeyDescriptionVal,
		ApiKeyId:          apiKeyIdVal,
		ApiKeyName:        apiKeyNameVal,
		DomainCode:        domainCodeVal,
		IsEnabled:         isEnabledVal,
		ModTime:           modTimeVal,
		Modifier:          modifierVal,
		PrimaryKey:        primaryKeyVal,
		SecondaryKey:      secondaryKeyVal,
		TenantId:          tenantIdVal,
		state:             attr.ValueStateKnown,
	}, diags
}

func NewApiKeyValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ApiKeyValue {
	object, diags := NewApiKeyValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewApiKeyValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ApiKeyType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewApiKeyValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewApiKeyValueUnknown(), nil
	}

	if in.IsNull() {
		return NewApiKeyValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewApiKeyValueMust(ApiKeyValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ApiKeyType) ValueType(ctx context.Context) attr.Value {
	return ApiKeyValue{}
}

type ApiKeyValue struct {
	ApiKeyDescription basetypes.StringValue `tfsdk:"api_key_description"`
	ApiKeyId          basetypes.StringValue `tfsdk:"api_key_id"`
	ApiKeyName        basetypes.StringValue `tfsdk:"api_key_name"`
	DomainCode        basetypes.StringValue `tfsdk:"domain_code"`
	IsEnabled         basetypes.BoolValue   `tfsdk:"is_enabled"`
	ModTime           basetypes.StringValue `tfsdk:"mod_time"`
	Modifier          basetypes.StringValue `tfsdk:"modifier"`
	PrimaryKey        basetypes.StringValue `tfsdk:"primary_key"`
	SecondaryKey      basetypes.StringValue `tfsdk:"secondary_key"`
	TenantId          basetypes.StringValue `tfsdk:"tenant_id"`
	state             attr.ValueState
}

func (v ApiKeyValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 10)

	var val tftypes.Value
	var err error

	attrTypes["api_key_description"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["api_key_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["api_key_name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["domain_code"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["is_enabled"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["mod_time"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["modifier"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["primary_key"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["secondary_key"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["tenant_id"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 10)

		val, err = v.ApiKeyDescription.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["api_key_description"] = val

		val, err = v.ApiKeyId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["api_key_id"] = val

		val, err = v.ApiKeyName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["api_key_name"] = val

		val, err = v.DomainCode.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["domain_code"] = val

		val, err = v.IsEnabled.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["is_enabled"] = val

		val, err = v.ModTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["mod_time"] = val

		val, err = v.Modifier.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["modifier"] = val

		val, err = v.PrimaryKey.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["primary_key"] = val

		val, err = v.SecondaryKey.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["secondary_key"] = val

		val, err = v.TenantId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["tenant_id"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ApiKeyValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ApiKeyValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ApiKeyValue) String() string {
	return "ApiKeyValue"
}

func (v ApiKeyValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"api_key_description": basetypes.StringType{},
		"api_key_id":          basetypes.StringType{},
		"api_key_name":        basetypes.StringType{},
		"domain_code":         basetypes.StringType{},
		"is_enabled":          basetypes.BoolType{},
		"mod_time":            basetypes.StringType{},
		"modifier":            basetypes.StringType{},
		"primary_key":         basetypes.StringType{},
		"secondary_key":       basetypes.StringType{},
		"tenant_id":           basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"api_key_description": v.ApiKeyDescription,
			"api_key_id":          v.ApiKeyId,
			"api_key_name":        v.ApiKeyName,
			"domain_code":         v.DomainCode,
			"is_enabled":          v.IsEnabled,
			"mod_time":            v.ModTime,
			"modifier":            v.Modifier,
			"primary_key":         v.PrimaryKey,
			"secondary_key":       v.SecondaryKey,
			"tenant_id":           v.TenantId,
		})

	return objVal, diags
}

func (v ApiKeyValue) Equal(o attr.Value) bool {
	other, ok := o.(ApiKeyValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.ApiKeyDescription.Equal(other.ApiKeyDescription) {
		return false
	}

	if !v.ApiKeyId.Equal(other.ApiKeyId) {
		return false
	}

	if !v.ApiKeyName.Equal(other.ApiKeyName) {
		return false
	}

	if !v.DomainCode.Equal(other.DomainCode) {
		return false
	}

	if !v.IsEnabled.Equal(other.IsEnabled) {
		return false
	}

	if !v.ModTime.Equal(other.ModTime) {
		return false
	}

	if !v.Modifier.Equal(other.Modifier) {
		return false
	}

	if !v.PrimaryKey.Equal(other.PrimaryKey) {
		return false
	}

	if !v.SecondaryKey.Equal(other.SecondaryKey) {
		return false
	}

	if !v.TenantId.Equal(other.TenantId) {
		return false
	}

	return true
}

func (v ApiKeyValue) Type(ctx context.Context) attr.Type {
	return ApiKeyType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ApiKeyValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"api_key_description": basetypes.StringType{},
		"api_key_id":          basetypes.StringType{},
		"api_key_name":        basetypes.StringType{},
		"domain_code":         basetypes.StringType{},
		"is_enabled":          basetypes.BoolType{},
		"mod_time":            basetypes.StringType{},
		"modifier":            basetypes.StringType{},
		"primary_key":         basetypes.StringType{},
		"secondary_key":       basetypes.StringType{},
		"tenant_id":           basetypes.StringType{},
	}
}
