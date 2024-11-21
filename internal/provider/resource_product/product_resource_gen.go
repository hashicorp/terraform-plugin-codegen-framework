package resource_product

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"os/exec"
	"time"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/common"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/conn"
	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func ProductResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Description<br>Length(Min/Max): 0/300",
				MarkdownDescription: "Description<br>Length(Min/Max): 0/300",
			},
			"product": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"action_name": schema.StringAttribute{
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
					"invoke_id": schema.StringAttribute{
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
				},
				CustomType: ProductType{
					ObjectType: types.ObjectType{
						AttrTypes: ProductValue{}.AttributeTypes(ctx),
					},
				},
				Computed: true,
			},
			"product_name": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
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
				Computed:            true,
				Optional:            true,
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
	resp.TypeName = req.ProviderTypeName + "_apigw_product"
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
		"productName": util.ClearDoubleQuote(plan.ProductName.String()),
"subscriptionCode": util.ClearDoubleQuote(plan.SubscriptionCode.String()),

	})
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "CreateProduct reqParams="+strings.Replace(string(reqBody), `\"`, "", -1))

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-s", "-X", "POST", "https://apigateway.apigw.ntruss.com/api/v1"+"/"+"products",
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-H", "cache-control: no-cache",
			"-H", "pragma: no-cache",
			"-d", strings.Replace(string(reqBody), `\"`, "", -1),
		)
	}

	response, err := util.Request(execFunc, "POST", "/api/v1"+"/"+"products", os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), strings.Replace(string(reqBody), `\"`, "", -1))
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("CREATING ERROR", "response invalid")
		return
	}

	err = waitResourceCreated(ctx, response["product"].(map[string]interface{})["productId"].(string))
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "CreateProduct response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, response["product"].(map[string]interface{})["productId"].(string))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *productResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var plan PostproductresponseModel

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan = *getAndRefresh(resp.Diagnostics, plan.ID.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *productResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PostproductresponseModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	reqBody, err := json.Marshal(map[string]string{
		"productName": util.ClearDoubleQuote(plan.ProductName.String()),
"subscriptionCode": util.ClearDoubleQuote(plan.SubscriptionCode.String()),

	})
	if err != nil {
		resp.Diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	tflog.Info(ctx, "UpdateProduct reqParams="+strings.Replace(string(reqBody), `\"`, "", -1))

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-s", "-X", "PATCH", "https://apigateway.apigw.ntruss.com/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(plan.Productid.String()),
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-H", "cache-control: no-cache",
			"-H", "pragma: no-cache",
			"-d", strings.Replace(string(reqBody), `\"`, "", -1),
		)
	}

	response, err := util.Request(execFunc, "PATCH", "/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(plan.Productid.String()), os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), strings.Replace(string(reqBody), `\"`, "", -1))
	if err != nil {
		resp.Diagnostics.AddError("UPDATING ERROR", err.Error())
		return
	}
	if response == nil {
		resp.Diagnostics.AddError("UPDATING ERROR", "response invalid")
		return
	}

	tflog.Info(ctx, "UpdateProduct response="+common.MarshalUncheckedString(response))

	plan = *getAndRefresh(resp.Diagnostics, plan.ID.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (a *productResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan PostproductresponseModel

	resp.Diagnostics.Append(req.State.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	execFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-s", "-X", "DELETE", "https://apigateway.apigw.ntruss.com/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(plan.Productid.String()),
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-H", "cache-control: no-cache",
			"-H", "pragma: no-cache",
		)
	}

	_, err := util.Request(execFunc, "DELETE", "/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(plan.Productid.String()), os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}

	err = waitResourceDeleted(ctx, util.ClearDoubleQuote(plan.ID.String()))
	if err != nil {
		resp.Diagnostics.AddError("DELETING ERROR", err.Error())
		return
	}
}

type PostproductresponseModel struct {
    ID types.String `tfsdk:"id"`
    ApiKeyDescription         types.String `tfsdk:"api_key_description"`
ApiKeyName         types.String `tfsdk:"api_key_name"`
Api_key         types.Object `tfsdk:"api_key"`
ApiKeyId         types.String `tfsdk:"api_key_id"`
DomainCode         types.String `tfsdk:"domain_code"`
IsEnabled         types.Bool `tfsdk:"is_enabled"`
ModTime         types.String `tfsdk:"mod_time"`
Modifier         types.String `tfsdk:"modifier"`
PrimaryKey         types.String `tfsdk:"primary_key"`
SecondaryKey         types.String `tfsdk:"secondary_key"`
TenantId         types.String `tfsdk:"tenant_id"`
Apikeyid         types.String `tfsdk:"apikeyid"`

}

func ConvertToFrameworkTypes(data map[string]interface{}, id string, rest []interface{}) (*PostproductresponseModel, error) {
	var dto PostproductresponseModel

	dto.ID = types.StringValue(id)

    dto.ApiKeyDescription = types.StringValue(data["api_key_description"].(string))
dto.ApiKeyName = types.StringValue(data["api_key_name"].(string))

			tempApi_key := data["api_key"].(map[string]interface{})
			convertedTempApi_key, err := util.ConvertMapToObject(context.TODO(), tempApi_key)
			if err != nil {
				fmt.Println("ConvertMapToObject Error")
			}

			dto.Api_key = diagOff(types.ObjectValueFrom, context.TODO(), types.ObjectType{AttrTypes: map[string]attr.Type{
				"api_key_description": types.StringType,
"api_key_id": types.StringType,
"api_key_name": types.StringType,
"domain_code": types.StringType,
"is_enabled": types.BoolType,
"mod_time": types.StringType,
"modifier": types.StringType,
"primary_key": types.StringType,
"secondary_key": types.StringType,
"tenant_id": types.StringType,

			}}.AttributeTypes(), convertedTempApi_key)
dto.ApiKeyId = types.StringValue(data["api_key_id"].(string))
dto.DomainCode = types.StringValue(data["domain_code"].(string))
dto.IsEnabled = types.BoolValue(data["is_enabled"].(bool))
dto.ModTime = types.StringValue(data["mod_time"].(string))
dto.Modifier = types.StringValue(data["modifier"].(string))
dto.PrimaryKey = types.StringValue(data["primary_key"].(string))
dto.SecondaryKey = types.StringValue(data["secondary_key"].(string))
dto.TenantId = types.StringValue(data["tenant_id"].(string))
dto.Apikeyid = types.StringValue(data["apikeyid"].(string))


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

func getAndRefresh(diagnostics diag.Diagnostics, id string, rest ...interface{}) *PostproductresponseModel {
	getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
		return exec.Command("curl", "-s", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(id),
			"-H", "Content-Type: application/json",
			"-H", "x-ncp-apigw-timestamp: "+timestamp,
			"-H", "x-ncp-iam-access-key: "+accessKey,
			"-H", "x-ncp-apigw-signature-v2: "+signature,
			"-H", "cache-control: no-cache",
			"-H", "pragma: no-cache",
		)
	}

	response, _ := util.Request(getExecFunc, "GET", "/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(id), os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
	if response == nil {
		diagnostics.AddError("UPDATING ERROR", "response invalid")
		return nil
	}

	newPlan, err := ConvertToFrameworkTypes(util.ConvertKeys(response).(map[string]interface{}), id, rest)
	if err != nil {
		diagnostics.AddError("CREATING ERROR", err.Error())
		return nil
	}

	return newPlan
}

func waitResourceCreated(ctx context.Context, id string) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"CREATING"},
		Target:  []string{"CREATED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
			return exec.Command("curl", "-s", "-X", "GET",  "https://apigateway.apigw.ntruss.com/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(id),
					"-H", "accept: application/json;charset=UTF-8",
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
					"-H", "cache-control: no-cache",
					"-H", "pragma: no-cache",
				)
			}

			response, err := util.Request(getExecFunc, "GET","/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(id), os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
			if err != nil {
				return response, "CREATING", nil
			}
			if response != nil {
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

func waitResourceDeleted(ctx context.Context, id string) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"DELETING"},
		Target:  []string{"DELETED"},
		Refresh: func() (interface{}, string, error) {
			getExecFunc := func(timestamp, accessKey, signature string) *exec.Cmd {
			return exec.Command("curl", "-s", "-X", "GET", "https://apigateway.apigw.ntruss.com/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(id),
					"-H", "accept: application/json;charset=UTF-8",
					"-H", "Content-Type: application/json",
					"-H", "x-ncp-apigw-timestamp: "+timestamp,
					"-H", "x-ncp-iam-access-key: "+accessKey,
					"-H", "x-ncp-apigw-signature-v2: "+signature,
					"-H", "cache-control: no-cache",
					"-H", "pragma: no-cache",
				)
			}

			response, _ := util.Request(getExecFunc, "GET", "/api/v1"+"/"+"products"+"/"+util.ClearDoubleQuote(id), os.Getenv("NCLOUD_ACCESS_KEY"), os.Getenv("NCLOUD_SECRET_KEY"), "")
			if response["error"] != nil {
				return response, "DELETED", nil
			}

			return response, "DELETING", nil
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

