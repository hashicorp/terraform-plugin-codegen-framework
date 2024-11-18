package resource_product

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/common"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/conn"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

func ProductResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description<br>Length(Min/Max): 0/300",
				MarkdownDescription: "Description<br>Length(Min/Max): 0/300",
			},
			"product": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"action_name": schema.Int64Attribute{
						Computed:            true,
						Description:         "Action Name",
						MarkdownDescription: "Action Name",
					},
					"disabled": schema.BoolAttribute{
						Computed:            true,
						Description:         "Disabled",
						MarkdownDescription: "Disabled",
					},
					"domain_code": schema.StringAttribute{
						Computed:            true,
						Description:         "Domain Code",
						MarkdownDescription: "Domain Code",
					},
					"invoke_id": schema.ListAttribute{
						ElementType:         types.StringType,
						Computed:            true,
						Description:         "Invoke Id",
						MarkdownDescription: "Invoke Id",
					},
					"is_deleted": schema.BoolAttribute{
						Computed:            true,
						Description:         "Is Deleted",
						MarkdownDescription: "Is Deleted",
					},
					"is_published": schema.BoolAttribute{
						Computed:            true,
						Description:         "Is Published",
						MarkdownDescription: "Is Published",
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
					"permission": schema.StringAttribute{
						Computed:            true,
						Description:         "Permission",
						MarkdownDescription: "Permission",
					},
					"product_description": schema.StringAttribute{
						Computed:            true,
						Description:         "Product Description",
						MarkdownDescription: "Product Description",
					},
					"product_id": schema.StringAttribute{
						Computed:            true,
						Description:         "Product Id",
						MarkdownDescription: "Product Id",
					},
					"product_name": schema.StringAttribute{
						Computed:            true,
						Description:         "Product Name",
						MarkdownDescription: "Product Name",
					},
					"subscription_code": schema.StringAttribute{
						Computed:            true,
						Description:         "Subscription Code<br>Allowable values: PROTECTED, PUBLIC",
						MarkdownDescription: "Subscription Code<br>Allowable values: PROTECTED, PUBLIC",
					},
					"tenant_id": schema.StringAttribute{
						Computed:            true,
						Description:         "Tenant Id",
						MarkdownDescription: "Tenant Id",
					},
					"test_nested_array": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kek": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"arrrr": schema.ListAttribute{
											ElementType: types.StringType,
											Computed:    true,
										},
										"kkkkk": schema.StringAttribute{
											Computed:            true,
											Description:         "kkk",
											MarkdownDescription: "kkk",
										},
									},
									CustomType: KekType{
										ObjectType: types.ObjectType{
											AttrTypes: KekValue{}.AttributeTypes(ctx),
										},
									},
									Computed: true,
								},
								"newnew": schema.StringAttribute{
									Computed:            true,
									Description:         "newenw",
									MarkdownDescription: "newenw",
								},
							},
							CustomType: TestNestedArrayType{
								ObjectType: types.ObjectType{
									AttrTypes: TestNestedArrayValue{}.AttributeTypes(ctx),
								},
							},
						},
						Computed:            true,
						Description:         "test",
						MarkdownDescription: "test",
					},
				},
				CustomType: ProductType{
					ObjectType: types.ObjectType{
						AttrTypes: ProductValue{}.AttributeTypes(ctx),
					},
				},
				Computed: true,
			},
			"product_name": schema.StringAttribute{
				Required:            true,
				Description:         "Product Name<br>Length(Min/Max): 0/100",
				MarkdownDescription: "Product Name<br>Length(Min/Max): 0/100",
			},
			"productid": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "product-id",
				MarkdownDescription: "product-id",
			},
			"subscription_code": schema.StringAttribute{
				Required:            true,
				Description:         "Subscription Code<br>Allowable values: PROTECTED, PUBLIC",
				MarkdownDescription: "Subscription Code<br>Allowable values: PROTECTED, PUBLIC",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"PROTECTED",
						"PUBLIC",
					),
				},
			},
		},
	}
}

func NewProductResource() resource.Resource {
	return &productResource{}
}

type productResource struct {
	config *conn.ProviderConfig
}

func (a *productResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (a *productResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (a *productResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "product"
}

func (a *productResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProductResourceSchema(ctx)
}

func (a *productResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PostproductresponseModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody, err := json.Marshal(map[string]string{
		"ProductName":      plan.ProductName.String(),
		"SubscriptionCode": plan.SubscriptionCode.String(),
	})
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "CreateProduct reqParams="+string(reqBody))

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "POST", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/",
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-d", string(reqBody),
		)
	}

	response, err := util.Request(execFunc, string(reqBody))
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("CREATING ERROR", "response invalid")
		return
	}

	err = waitResourceCreated(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "CreateProduct response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *productResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan PostproductresponseModel

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan = *getAndRefresh(resp.Diagnostics, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *productResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PostproductresponseModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody, err := json.Marshal(map[string]string{
		"ProductName":      plan.ProductName.String(),
		"SubscriptionCode": plan.SubscriptionCode.String(),
	})
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "UpdateProduct reqParams="+string(reqBody))

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "PATCH", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/"+plan.ProductId.String(),
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-d", string(reqBody),
		)
	}

	response, err := util.Request(execFunc, string(reqBody))
	if err != nil {
		resp.Diagnostics.AddError("UPDATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("UPDATING ERROR", "response invalid")
		return
	}

	tflog.Info(ctx, "UpdateProduct response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *productResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan PostproductresponseModel

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "DELETE", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/"+plan.ProductId.String(),
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

	err = waitResourceDeleted(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "DeleteProduct response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

type PostproductresponseModel struct {
	Description      types.String `json:"description"`
	ProductName      types.String `json:"product_name"`
	SubscriptionCode types.String `json:"subscription_code"`
	Product          types.Object `json:"product"`
	Productid        types.String `json:"productid"`
}

func ConvertToFrameworkTypes(data map[string]interface{}) (*PostproductresponseModel, error) {
	var dto PostproductresponseModel

	dto.Description = types.StringValue(data["description"].(string))
	dto.ProductName = types.StringValue(data["product_name"].(string))
	dto.SubscriptionCode = types.StringValue(data["subscription_code"].(string))

	tempProduct := data["product"].(map[string]interface{})
	convertedTempProduct, err := util.ConvertMapToObject(context.TODO(), tempProduct)
	if err != nil {
		log.Fatalf("ConvertMapToObject err: product", err)
	}

	dto.Product = diagOff(types.ObjectValueFrom, context.TODO(), types.ObjectType{AttrTypes: map[string]attr.Type{
		"actionName":         types.Int64Type,
		"disabled":           types.BoolType,
		"domainCode":         types.StringType,
		"invokeId":           types.ListType{ElemType: types.StringType},
		"isDeleted":          types.BoolType,
		"isPublished":        types.BoolType,
		"modTime":            types.StringType,
		"modifier":           types.StringType,
		"permission":         types.StringType,
		"productDescription": types.StringType,
		"productId":          types.StringType,
		"productName":        types.StringType,
		"subscriptionCode":   types.StringType,
		"tenantId":           types.StringType,

		"testNestedArray": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{

			"kek": types.ObjectType{AttrTypes: map[string]attr.Type{
				"arrrr": types.ListType{ElemType: types.StringType},
				"kkkkk": types.StringType,
			}},

			"newnew": types.StringType,
		},
		},
		}}}.AttributeTypes(), convertedTempProduct)
	dto.Productid = types.StringValue(data["productid"].(string))

	return &dto, nil
}

func diagOff[V, T interface{}](input func(ctx context.Context, elementType T, elements any) (V, diag.Diagnostics), ctx context.Context, elementType T, elements any) V {
	var emptyReturn V

	v, diags := input(ctx, elementType, elements)

	if diags.HasError() {
		diags.AddError("REFRESHING ERROR", "invalid diagOff operation")
		return emptyReturn
	}

	return v
}

func getAndRefresh(diagnostics diag.Diagnostics, plan PostproductresponseModel) *PostproductresponseModel {
	getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/"+plan.ProductId.String(),
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

	newPlan, err := ConvertToFrameworkTypes(response)
	if err != nil {
		diagnostics.AddError("CREATING ERROR", err.Error())
		return nil
	}

	return newPlan
}

func waitResourceCreated(ctx context.Context, plan PostproductresponseModel) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATING"},
		Target:  []string{"CREATED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
				return exec.Command("curl", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/"+plan.ProductId.String(),
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

func waitResourceDeleted(ctx context.Context, plan PostproductresponseModel) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"DELETING"},
		Target:  []string{"DELETED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
				return exec.Command("curl", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1/api-keys/"+plan.ProductId.String(),
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

type ProductType struct {
	basetypes.ObjectType
}

func (t ProductType) Equal(o attr.Type) bool {
	other, ok := o.(ProductType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ProductType) String() string {
	return "ProductType"
}

func (t ProductType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	actionNameAttribute, ok := attributes["action_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`action_name is missing from object`)

		return nil, diags
	}

	actionNameVal, ok := actionNameAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`action_name expected to be basetypes.Int64Value, was: %T`, actionNameAttribute))
	}

	disabledAttribute, ok := attributes["disabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`disabled is missing from object`)

		return nil, diags
	}

	disabledVal, ok := disabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`disabled expected to be basetypes.BoolValue, was: %T`, disabledAttribute))
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

	invokeIdAttribute, ok := attributes["invoke_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`invoke_id is missing from object`)

		return nil, diags
	}

	invokeIdVal, ok := invokeIdAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`invoke_id expected to be basetypes.ListValue, was: %T`, invokeIdAttribute))
	}

	isDeletedAttribute, ok := attributes["is_deleted"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`is_deleted is missing from object`)

		return nil, diags
	}

	isDeletedVal, ok := isDeletedAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`is_deleted expected to be basetypes.BoolValue, was: %T`, isDeletedAttribute))
	}

	isPublishedAttribute, ok := attributes["is_published"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`is_published is missing from object`)

		return nil, diags
	}

	isPublishedVal, ok := isPublishedAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`is_published expected to be basetypes.BoolValue, was: %T`, isPublishedAttribute))
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

	permissionAttribute, ok := attributes["permission"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`permission is missing from object`)

		return nil, diags
	}

	permissionVal, ok := permissionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`permission expected to be basetypes.StringValue, was: %T`, permissionAttribute))
	}

	productDescriptionAttribute, ok := attributes["product_description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_description is missing from object`)

		return nil, diags
	}

	productDescriptionVal, ok := productDescriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_description expected to be basetypes.StringValue, was: %T`, productDescriptionAttribute))
	}

	productIdAttribute, ok := attributes["product_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_id is missing from object`)

		return nil, diags
	}

	productIdVal, ok := productIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_id expected to be basetypes.StringValue, was: %T`, productIdAttribute))
	}

	productNameAttribute, ok := attributes["product_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_name is missing from object`)

		return nil, diags
	}

	productNameVal, ok := productNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_name expected to be basetypes.StringValue, was: %T`, productNameAttribute))
	}

	subscriptionCodeAttribute, ok := attributes["subscription_code"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`subscription_code is missing from object`)

		return nil, diags
	}

	subscriptionCodeVal, ok := subscriptionCodeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`subscription_code expected to be basetypes.StringValue, was: %T`, subscriptionCodeAttribute))
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

	testNestedArrayAttribute, ok := attributes["test_nested_array"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`test_nested_array is missing from object`)

		return nil, diags
	}

	testNestedArrayVal, ok := testNestedArrayAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`test_nested_array expected to be basetypes.ListValue, was: %T`, testNestedArrayAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ProductValue{
		ActionName:         actionNameVal,
		Disabled:           disabledVal,
		DomainCode:         domainCodeVal,
		InvokeId:           invokeIdVal,
		IsDeleted:          isDeletedVal,
		IsPublished:        isPublishedVal,
		ModTime:            modTimeVal,
		Modifier:           modifierVal,
		Permission:         permissionVal,
		ProductDescription: productDescriptionVal,
		ProductId:          productIdVal,
		ProductName:        productNameVal,
		SubscriptionCode:   subscriptionCodeVal,
		TenantId:           tenantIdVal,
		TestNestedArray:    testNestedArrayVal,
		state:              attr.ValueStateKnown,
	}, diags
}

func NewProductValueNull() ProductValue {
	return ProductValue{
		state: attr.ValueStateNull,
	}
}

func NewProductValueUnknown() ProductValue {
	return ProductValue{
		state: attr.ValueStateUnknown,
	}
}

func NewProductValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ProductValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ProductValue Attribute Value",
				"While creating a ProductValue value, a missing attribute value was detected. "+
					"A ProductValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ProductValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ProductValue Attribute Type",
				"While creating a ProductValue value, an invalid attribute value was detected. "+
					"A ProductValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ProductValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ProductValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ProductValue Attribute Value",
				"While creating a ProductValue value, an extra attribute value was detected. "+
					"A ProductValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ProductValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewProductValueUnknown(), diags
	}

	actionNameAttribute, ok := attributes["action_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`action_name is missing from object`)

		return NewProductValueUnknown(), diags
	}

	actionNameVal, ok := actionNameAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`action_name expected to be basetypes.Int64Value, was: %T`, actionNameAttribute))
	}

	disabledAttribute, ok := attributes["disabled"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`disabled is missing from object`)

		return NewProductValueUnknown(), diags
	}

	disabledVal, ok := disabledAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`disabled expected to be basetypes.BoolValue, was: %T`, disabledAttribute))
	}

	domainCodeAttribute, ok := attributes["domain_code"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`domain_code is missing from object`)

		return NewProductValueUnknown(), diags
	}

	domainCodeVal, ok := domainCodeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`domain_code expected to be basetypes.StringValue, was: %T`, domainCodeAttribute))
	}

	invokeIdAttribute, ok := attributes["invoke_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`invoke_id is missing from object`)

		return NewProductValueUnknown(), diags
	}

	invokeIdVal, ok := invokeIdAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`invoke_id expected to be basetypes.ListValue, was: %T`, invokeIdAttribute))
	}

	isDeletedAttribute, ok := attributes["is_deleted"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`is_deleted is missing from object`)

		return NewProductValueUnknown(), diags
	}

	isDeletedVal, ok := isDeletedAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`is_deleted expected to be basetypes.BoolValue, was: %T`, isDeletedAttribute))
	}

	isPublishedAttribute, ok := attributes["is_published"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`is_published is missing from object`)

		return NewProductValueUnknown(), diags
	}

	isPublishedVal, ok := isPublishedAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`is_published expected to be basetypes.BoolValue, was: %T`, isPublishedAttribute))
	}

	modTimeAttribute, ok := attributes["mod_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mod_time is missing from object`)

		return NewProductValueUnknown(), diags
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

		return NewProductValueUnknown(), diags
	}

	modifierVal, ok := modifierAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`modifier expected to be basetypes.StringValue, was: %T`, modifierAttribute))
	}

	permissionAttribute, ok := attributes["permission"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`permission is missing from object`)

		return NewProductValueUnknown(), diags
	}

	permissionVal, ok := permissionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`permission expected to be basetypes.StringValue, was: %T`, permissionAttribute))
	}

	productDescriptionAttribute, ok := attributes["product_description"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_description is missing from object`)

		return NewProductValueUnknown(), diags
	}

	productDescriptionVal, ok := productDescriptionAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_description expected to be basetypes.StringValue, was: %T`, productDescriptionAttribute))
	}

	productIdAttribute, ok := attributes["product_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_id is missing from object`)

		return NewProductValueUnknown(), diags
	}

	productIdVal, ok := productIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_id expected to be basetypes.StringValue, was: %T`, productIdAttribute))
	}

	productNameAttribute, ok := attributes["product_name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_name is missing from object`)

		return NewProductValueUnknown(), diags
	}

	productNameVal, ok := productNameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_name expected to be basetypes.StringValue, was: %T`, productNameAttribute))
	}

	subscriptionCodeAttribute, ok := attributes["subscription_code"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`subscription_code is missing from object`)

		return NewProductValueUnknown(), diags
	}

	subscriptionCodeVal, ok := subscriptionCodeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`subscription_code expected to be basetypes.StringValue, was: %T`, subscriptionCodeAttribute))
	}

	tenantIdAttribute, ok := attributes["tenant_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tenant_id is missing from object`)

		return NewProductValueUnknown(), diags
	}

	tenantIdVal, ok := tenantIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tenant_id expected to be basetypes.StringValue, was: %T`, tenantIdAttribute))
	}

	testNestedArrayAttribute, ok := attributes["test_nested_array"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`test_nested_array is missing from object`)

		return NewProductValueUnknown(), diags
	}

	testNestedArrayVal, ok := testNestedArrayAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`test_nested_array expected to be basetypes.ListValue, was: %T`, testNestedArrayAttribute))
	}

	if diags.HasError() {
		return NewProductValueUnknown(), diags
	}

	return ProductValue{
		ActionName:         actionNameVal,
		Disabled:           disabledVal,
		DomainCode:         domainCodeVal,
		InvokeId:           invokeIdVal,
		IsDeleted:          isDeletedVal,
		IsPublished:        isPublishedVal,
		ModTime:            modTimeVal,
		Modifier:           modifierVal,
		Permission:         permissionVal,
		ProductDescription: productDescriptionVal,
		ProductId:          productIdVal,
		ProductName:        productNameVal,
		SubscriptionCode:   subscriptionCodeVal,
		TenantId:           tenantIdVal,
		TestNestedArray:    testNestedArrayVal,
		state:              attr.ValueStateKnown,
	}, diags
}

func NewProductValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ProductValue {
	object, diags := NewProductValue(attributeTypes, attributes)

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

		panic("NewProductValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ProductType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewProductValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewProductValueUnknown(), nil
	}

	if in.IsNull() {
		return NewProductValueNull(), nil
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

	return NewProductValueMust(ProductValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ProductType) ValueType(ctx context.Context) attr.Value {
	return ProductValue{}
}

type ProductValue struct {
	ActionName         basetypes.Int64Value  `tfsdk:"action_name"`
	Disabled           basetypes.BoolValue   `tfsdk:"disabled"`
	DomainCode         basetypes.StringValue `tfsdk:"domain_code"`
	InvokeId           basetypes.ListValue   `tfsdk:"invoke_id"`
	IsDeleted          basetypes.BoolValue   `tfsdk:"is_deleted"`
	IsPublished        basetypes.BoolValue   `tfsdk:"is_published"`
	ModTime            basetypes.StringValue `tfsdk:"mod_time"`
	Modifier           basetypes.StringValue `tfsdk:"modifier"`
	Permission         basetypes.StringValue `tfsdk:"permission"`
	ProductDescription basetypes.StringValue `tfsdk:"product_description"`
	ProductId          basetypes.StringValue `tfsdk:"product_id"`
	ProductName        basetypes.StringValue `tfsdk:"product_name"`
	SubscriptionCode   basetypes.StringValue `tfsdk:"subscription_code"`
	TenantId           basetypes.StringValue `tfsdk:"tenant_id"`
	TestNestedArray    basetypes.ListValue   `tfsdk:"test_nested_array"`
	state              attr.ValueState
}

func (v ProductValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 15)

	var val tftypes.Value
	var err error

	attrTypes["action_name"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["disabled"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["domain_code"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["invoke_id"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["is_deleted"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["is_published"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["mod_time"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["modifier"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["permission"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["product_description"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["product_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["product_name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["subscription_code"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["tenant_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["test_nested_array"] = basetypes.ListType{
		ElemType: TestNestedArrayValue{}.Type(ctx),
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 15)

		val, err = v.ActionName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["action_name"] = val

		val, err = v.Disabled.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["disabled"] = val

		val, err = v.DomainCode.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["domain_code"] = val

		val, err = v.InvokeId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["invoke_id"] = val

		val, err = v.IsDeleted.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["is_deleted"] = val

		val, err = v.IsPublished.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["is_published"] = val

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

		val, err = v.Permission.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["permission"] = val

		val, err = v.ProductDescription.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["product_description"] = val

		val, err = v.ProductId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["product_id"] = val

		val, err = v.ProductName.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["product_name"] = val

		val, err = v.SubscriptionCode.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["subscription_code"] = val

		val, err = v.TenantId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["tenant_id"] = val

		val, err = v.TestNestedArray.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["test_nested_array"] = val

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

func (v ProductValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ProductValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ProductValue) String() string {
	return "ProductValue"
}

func (v ProductValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	testNestedArray := types.ListValueMust(
		TestNestedArrayType{
			basetypes.ObjectType{
				AttrTypes: TestNestedArrayValue{}.AttributeTypes(ctx),
			},
		},
		v.TestNestedArray.Elements(),
	)

	if v.TestNestedArray.IsNull() {
		testNestedArray = types.ListNull(
			TestNestedArrayType{
				basetypes.ObjectType{
					AttrTypes: TestNestedArrayValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	if v.TestNestedArray.IsUnknown() {
		testNestedArray = types.ListUnknown(
			TestNestedArrayType{
				basetypes.ObjectType{
					AttrTypes: TestNestedArrayValue{}.AttributeTypes(ctx),
				},
			},
		)
	}

	var invokeIdVal basetypes.ListValue
	switch {
	case v.InvokeId.IsUnknown():
		invokeIdVal = types.ListUnknown(types.StringType)
	case v.InvokeId.IsNull():
		invokeIdVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		invokeIdVal, d = types.ListValue(types.StringType, v.InvokeId.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"action_name": basetypes.Int64Type{},
			"disabled":    basetypes.BoolType{},
			"domain_code": basetypes.StringType{},
			"invoke_id": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_deleted":          basetypes.BoolType{},
			"is_published":        basetypes.BoolType{},
			"mod_time":            basetypes.StringType{},
			"modifier":            basetypes.StringType{},
			"permission":          basetypes.StringType{},
			"product_description": basetypes.StringType{},
			"product_id":          basetypes.StringType{},
			"product_name":        basetypes.StringType{},
			"subscription_code":   basetypes.StringType{},
			"tenant_id":           basetypes.StringType{},
			"test_nested_array": basetypes.ListType{
				ElemType: TestNestedArrayValue{}.Type(ctx),
			},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"action_name": basetypes.Int64Type{},
		"disabled":    basetypes.BoolType{},
		"domain_code": basetypes.StringType{},
		"invoke_id": basetypes.ListType{
			ElemType: types.StringType,
		},
		"is_deleted":          basetypes.BoolType{},
		"is_published":        basetypes.BoolType{},
		"mod_time":            basetypes.StringType{},
		"modifier":            basetypes.StringType{},
		"permission":          basetypes.StringType{},
		"product_description": basetypes.StringType{},
		"product_id":          basetypes.StringType{},
		"product_name":        basetypes.StringType{},
		"subscription_code":   basetypes.StringType{},
		"tenant_id":           basetypes.StringType{},
		"test_nested_array": basetypes.ListType{
			ElemType: TestNestedArrayValue{}.Type(ctx),
		},
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
			"action_name":         v.ActionName,
			"disabled":            v.Disabled,
			"domain_code":         v.DomainCode,
			"invoke_id":           invokeIdVal,
			"is_deleted":          v.IsDeleted,
			"is_published":        v.IsPublished,
			"mod_time":            v.ModTime,
			"modifier":            v.Modifier,
			"permission":          v.Permission,
			"product_description": v.ProductDescription,
			"product_id":          v.ProductId,
			"product_name":        v.ProductName,
			"subscription_code":   v.SubscriptionCode,
			"tenant_id":           v.TenantId,
			"test_nested_array":   testNestedArray,
		})

	return objVal, diags
}

func (v ProductValue) Equal(o attr.Value) bool {
	other, ok := o.(ProductValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.ActionName.Equal(other.ActionName) {
		return false
	}

	if !v.Disabled.Equal(other.Disabled) {
		return false
	}

	if !v.DomainCode.Equal(other.DomainCode) {
		return false
	}

	if !v.InvokeId.Equal(other.InvokeId) {
		return false
	}

	if !v.IsDeleted.Equal(other.IsDeleted) {
		return false
	}

	if !v.IsPublished.Equal(other.IsPublished) {
		return false
	}

	if !v.ModTime.Equal(other.ModTime) {
		return false
	}

	if !v.Modifier.Equal(other.Modifier) {
		return false
	}

	if !v.Permission.Equal(other.Permission) {
		return false
	}

	if !v.ProductDescription.Equal(other.ProductDescription) {
		return false
	}

	if !v.ProductId.Equal(other.ProductId) {
		return false
	}

	if !v.ProductName.Equal(other.ProductName) {
		return false
	}

	if !v.SubscriptionCode.Equal(other.SubscriptionCode) {
		return false
	}

	if !v.TenantId.Equal(other.TenantId) {
		return false
	}

	if !v.TestNestedArray.Equal(other.TestNestedArray) {
		return false
	}

	return true
}

func (v ProductValue) Type(ctx context.Context) attr.Type {
	return ProductType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ProductValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"action_name": basetypes.Int64Type{},
		"disabled":    basetypes.BoolType{},
		"domain_code": basetypes.StringType{},
		"invoke_id": basetypes.ListType{
			ElemType: types.StringType,
		},
		"is_deleted":          basetypes.BoolType{},
		"is_published":        basetypes.BoolType{},
		"mod_time":            basetypes.StringType{},
		"modifier":            basetypes.StringType{},
		"permission":          basetypes.StringType{},
		"product_description": basetypes.StringType{},
		"product_id":          basetypes.StringType{},
		"product_name":        basetypes.StringType{},
		"subscription_code":   basetypes.StringType{},
		"tenant_id":           basetypes.StringType{},
		"test_nested_array": basetypes.ListType{
			ElemType: TestNestedArrayValue{}.Type(ctx),
		},
	}
}

type TestNestedArrayType struct {
	basetypes.ObjectType
}

func (t TestNestedArrayType) Equal(o attr.Type) bool {
	other, ok := o.(TestNestedArrayType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t TestNestedArrayType) String() string {
	return "TestNestedArrayType"
}

func (t TestNestedArrayType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	kekAttribute, ok := attributes["kek"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`kek is missing from object`)

		return nil, diags
	}

	kekVal, ok := kekAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`kek expected to be basetypes.ObjectValue, was: %T`, kekAttribute))
	}

	newnewAttribute, ok := attributes["newnew"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`newnew is missing from object`)

		return nil, diags
	}

	newnewVal, ok := newnewAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`newnew expected to be basetypes.StringValue, was: %T`, newnewAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return TestNestedArrayValue{
		Kek:    kekVal,
		Newnew: newnewVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewTestNestedArrayValueNull() TestNestedArrayValue {
	return TestNestedArrayValue{
		state: attr.ValueStateNull,
	}
}

func NewTestNestedArrayValueUnknown() TestNestedArrayValue {
	return TestNestedArrayValue{
		state: attr.ValueStateUnknown,
	}
}

func NewTestNestedArrayValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (TestNestedArrayValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing TestNestedArrayValue Attribute Value",
				"While creating a TestNestedArrayValue value, a missing attribute value was detected. "+
					"A TestNestedArrayValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("TestNestedArrayValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid TestNestedArrayValue Attribute Type",
				"While creating a TestNestedArrayValue value, an invalid attribute value was detected. "+
					"A TestNestedArrayValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("TestNestedArrayValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("TestNestedArrayValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra TestNestedArrayValue Attribute Value",
				"While creating a TestNestedArrayValue value, an extra attribute value was detected. "+
					"A TestNestedArrayValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra TestNestedArrayValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewTestNestedArrayValueUnknown(), diags
	}

	kekAttribute, ok := attributes["kek"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`kek is missing from object`)

		return NewTestNestedArrayValueUnknown(), diags
	}

	kekVal, ok := kekAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`kek expected to be basetypes.ObjectValue, was: %T`, kekAttribute))
	}

	newnewAttribute, ok := attributes["newnew"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`newnew is missing from object`)

		return NewTestNestedArrayValueUnknown(), diags
	}

	newnewVal, ok := newnewAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`newnew expected to be basetypes.StringValue, was: %T`, newnewAttribute))
	}

	if diags.HasError() {
		return NewTestNestedArrayValueUnknown(), diags
	}

	return TestNestedArrayValue{
		Kek:    kekVal,
		Newnew: newnewVal,
		state:  attr.ValueStateKnown,
	}, diags
}

func NewTestNestedArrayValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) TestNestedArrayValue {
	object, diags := NewTestNestedArrayValue(attributeTypes, attributes)

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

		panic("NewTestNestedArrayValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t TestNestedArrayType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewTestNestedArrayValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewTestNestedArrayValueUnknown(), nil
	}

	if in.IsNull() {
		return NewTestNestedArrayValueNull(), nil
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

	return NewTestNestedArrayValueMust(TestNestedArrayValue{}.AttributeTypes(ctx), attributes), nil
}

func (t TestNestedArrayType) ValueType(ctx context.Context) attr.Value {
	return TestNestedArrayValue{}
}

type TestNestedArrayValue struct {
	Kek    basetypes.ObjectValue `tfsdk:"kek"`
	Newnew basetypes.StringValue `tfsdk:"newnew"`
	state  attr.ValueState
}

func (v TestNestedArrayValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["kek"] = basetypes.ObjectType{
		AttrTypes: KekValue{}.AttributeTypes(ctx),
	}.TerraformType(ctx)
	attrTypes["newnew"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Kek.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["kek"] = val

		val, err = v.Newnew.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["newnew"] = val

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

func (v TestNestedArrayValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v TestNestedArrayValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v TestNestedArrayValue) String() string {
	return "TestNestedArrayValue"
}

func (v TestNestedArrayValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var kek basetypes.ObjectValue

	if v.Kek.IsNull() {
		kek = types.ObjectNull(
			KekValue{}.AttributeTypes(ctx),
		)
	}

	if v.Kek.IsUnknown() {
		kek = types.ObjectUnknown(
			KekValue{}.AttributeTypes(ctx),
		)
	}

	if !v.Kek.IsNull() && !v.Kek.IsUnknown() {
		kek = types.ObjectValueMust(
			KekValue{}.AttributeTypes(ctx),
			v.Kek.Attributes(),
		)
	}

	attributeTypes := map[string]attr.Type{
		"kek": basetypes.ObjectType{
			AttrTypes: KekValue{}.AttributeTypes(ctx),
		},
		"newnew": basetypes.StringType{},
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
			"kek":    kek,
			"newnew": v.Newnew,
		})

	return objVal, diags
}

func (v TestNestedArrayValue) Equal(o attr.Value) bool {
	other, ok := o.(TestNestedArrayValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Kek.Equal(other.Kek) {
		return false
	}

	if !v.Newnew.Equal(other.Newnew) {
		return false
	}

	return true
}

func (v TestNestedArrayValue) Type(ctx context.Context) attr.Type {
	return TestNestedArrayType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v TestNestedArrayValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"kek": basetypes.ObjectType{
			AttrTypes: KekValue{}.AttributeTypes(ctx),
		},
		"newnew": basetypes.StringType{},
	}
}

type KekType struct {
	basetypes.ObjectType
}

func (t KekType) Equal(o attr.Type) bool {
	other, ok := o.(KekType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t KekType) String() string {
	return "KekType"
}

func (t KekType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	arrrrAttribute, ok := attributes["arrrr"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`arrrr is missing from object`)

		return nil, diags
	}

	arrrrVal, ok := arrrrAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`arrrr expected to be basetypes.ListValue, was: %T`, arrrrAttribute))
	}

	kkkkkAttribute, ok := attributes["kkkkk"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`kkkkk is missing from object`)

		return nil, diags
	}

	kkkkkVal, ok := kkkkkAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`kkkkk expected to be basetypes.StringValue, was: %T`, kkkkkAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return KekValue{
		Arrrr: arrrrVal,
		Kkkkk: kkkkkVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewKekValueNull() KekValue {
	return KekValue{
		state: attr.ValueStateNull,
	}
}

func NewKekValueUnknown() KekValue {
	return KekValue{
		state: attr.ValueStateUnknown,
	}
}

func NewKekValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (KekValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing KekValue Attribute Value",
				"While creating a KekValue value, a missing attribute value was detected. "+
					"A KekValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("KekValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid KekValue Attribute Type",
				"While creating a KekValue value, an invalid attribute value was detected. "+
					"A KekValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("KekValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("KekValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra KekValue Attribute Value",
				"While creating a KekValue value, an extra attribute value was detected. "+
					"A KekValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra KekValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewKekValueUnknown(), diags
	}

	arrrrAttribute, ok := attributes["arrrr"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`arrrr is missing from object`)

		return NewKekValueUnknown(), diags
	}

	arrrrVal, ok := arrrrAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`arrrr expected to be basetypes.ListValue, was: %T`, arrrrAttribute))
	}

	kkkkkAttribute, ok := attributes["kkkkk"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`kkkkk is missing from object`)

		return NewKekValueUnknown(), diags
	}

	kkkkkVal, ok := kkkkkAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`kkkkk expected to be basetypes.StringValue, was: %T`, kkkkkAttribute))
	}

	if diags.HasError() {
		return NewKekValueUnknown(), diags
	}

	return KekValue{
		Arrrr: arrrrVal,
		Kkkkk: kkkkkVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewKekValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) KekValue {
	object, diags := NewKekValue(attributeTypes, attributes)

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

		panic("NewKekValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t KekType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewKekValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewKekValueUnknown(), nil
	}

	if in.IsNull() {
		return NewKekValueNull(), nil
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

	return NewKekValueMust(KekValue{}.AttributeTypes(ctx), attributes), nil
}

func (t KekType) ValueType(ctx context.Context) attr.Value {
	return KekValue{}
}

type KekValue struct {
	Arrrr basetypes.ListValue   `tfsdk:"arrrr"`
	Kkkkk basetypes.StringValue `tfsdk:"kkkkk"`
	state attr.ValueState
}

func (v KekValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["arrrr"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["kkkkk"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Arrrr.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["arrrr"] = val

		val, err = v.Kkkkk.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["kkkkk"] = val

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

func (v KekValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v KekValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v KekValue) String() string {
	return "KekValue"
}

func (v KekValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var arrrrVal basetypes.ListValue
	switch {
	case v.Arrrr.IsUnknown():
		arrrrVal = types.ListUnknown(types.StringType)
	case v.Arrrr.IsNull():
		arrrrVal = types.ListNull(types.StringType)
	default:
		var d diag.Diagnostics
		arrrrVal, d = types.ListValue(types.StringType, v.Arrrr.Elements())
		diags.Append(d...)
	}

	if diags.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"arrrr": basetypes.ListType{
				ElemType: types.StringType,
			},
			"kkkkk": basetypes.StringType{},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"arrrr": basetypes.ListType{
			ElemType: types.StringType,
		},
		"kkkkk": basetypes.StringType{},
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
			"arrrr": arrrrVal,
			"kkkkk": v.Kkkkk,
		})

	return objVal, diags
}

func (v KekValue) Equal(o attr.Value) bool {
	other, ok := o.(KekValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Arrrr.Equal(other.Arrrr) {
		return false
	}

	if !v.Kkkkk.Equal(other.Kkkkk) {
		return false
	}

	return true
}

func (v KekValue) Type(ctx context.Context) attr.Type {
	return KekType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v KekValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"arrrr": basetypes.ListType{
			ElemType: types.StringType,
		},
		"kkkkk": basetypes.StringType{},
	}
}
