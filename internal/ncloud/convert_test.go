package ncloud

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestConverting(t *testing.T) {
	var result map[string]interface{}
	jsonData := `{
	"description": "a",
	"product_name": "n",
	"subscription_code": "c",
	"new": [1,1],
	"productid": "i",
	"product": {
		"actionName": 1,
		"disabled": true,
		"domainCode": "string",
		"invokeId": ["s"],
		"isDeleted": true,
		"isPublished": true,
		"modTime": "2024-11-15T04:33:36.570Z",
		"modifier": "string",
		"permission": "string",
		"productDescription": "string",
		"productId": "string",
		"productName": "string",
		"subscriptionCode": "PROTECTED",
		"tenantId": "string",
		"testNestedArray": [{
			"kek": {
				"arrrr": ["21"],
				"kkkkk": "hello"
				},
				"newnew": "1"
			}]
		}
	}`
	if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
		panic(err)
	}

	m := convertToFrameworkTypes(result)
	fmt.Println("this is m: ", m)
}

type Model struct {
	Description       types.String `json:"description"`
	Product_name      types.String `json:"product_name"`
	Subscription_code types.String `json:"subscription_code"`
	Product           types.Object `json:"product"`
	New               types.List   `json:"new"`
	Productid         types.String `json:"productid"`
}

func convertToFrameworkTypes(data map[string]interface{}) *Model {
	var dto Model

	dto.Description = types.StringValue(data["description"].(string))
	dto.Product_name = types.StringValue(data["product_name"].(string))
	dto.Subscription_code = types.StringValue(data["subscription_code"].(string))

	tempProduct := data["product"].(map[string]interface{})
	convertedTempProduct, err := util.ConvertMapToObject(context.TODO(), tempProduct)
	if err != nil {
		panic(err)
	}

	dto.Product = util.DiagOff(types.ObjectValueFrom, context.TODO(), types.ObjectType{AttrTypes: map[string]attr.Type{
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

	return &dto
}
