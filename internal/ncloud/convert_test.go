package ncloud

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Result struct {
	Description types.String `tfsdk:"description"`
}

// 중첩된 구조의 map을 Terraform Object로 변환하는 함수
func convertMapToObject(ctx context.Context, data map[string]interface{}) (types.Object, error) {
	attrTypes := make(map[string]attr.Type)
	attrValues := make(map[string]attr.Value)

	for key, value := range data {
		attrType, attrValue, err := convertInterfaceToAttr(ctx, value)
		if err != nil {
			return types.Object{}, fmt.Errorf("error converting field %s: %v", key, err)
		}

		attrTypes[key] = attrType
		attrValues[key] = attrValue
	}

	r, _ := types.ObjectValue(attrTypes, attrValues)

	return r, nil
}

// interface{} 값을 Terraform attr.Type과 attr.Value로 변환하는 함수
func convertInterfaceToAttr(ctx context.Context, value interface{}) (attr.Type, attr.Value, error) {
	switch v := value.(type) {
	case string:
		return types.StringType, types.StringValue(v), nil

	case float64:
		return types.Int64Type, types.Int64Value(int64(v)), nil

	case bool:
		return types.BoolType, types.BoolValue(v), nil

	case []interface{}:
		if len(v) == 0 {
			// 빈 배열인 경우 기본적으로 문자열 리스트로 처리
			return types.ListType{ElemType: types.StringType},
				types.ListValueMust(types.StringType, []attr.Value{}),
				nil
		}

		// 첫 번째 요소의 타입을 기준으로 배열 타입 결정
		elemType, _, err := convertInterfaceToAttr(ctx, v[0])
		if err != nil {
			return nil, nil, err
		}

		values := make([]attr.Value, len(v))
		for i, item := range v {
			_, value, err := convertInterfaceToAttr(ctx, item)
			if err != nil {
				return nil, nil, err
			}
			values[i] = value
		}

		listType := types.ListType{ElemType: elemType}
		listValue, diags := types.ListValue(elemType, values)
		if diags.HasError() {
			return nil, nil, err
		}

		return listType, listValue, nil

	case map[string]interface{}:
		objValue, err := convertMapToObject(ctx, v)
		if err != nil {
			return nil, nil, err
		}
		return objValue.Type(ctx), objValue, nil

	case nil:
		return types.StringType, types.StringNull(), nil

	default:
		return nil, nil, fmt.Errorf("unsupported type: %T", value)
	}
}

func TestConverting(t *testing.T) {
	var result map[string]interface{}
	jsonData := `{
	"description": "a",
	"product_name": "n",
	"subscription_code": "c",
	"new": [1,1],
	"productid": "i",
	"product": {
		"actionName": "string",
		"disabled": true,
		"domainCode": "string",
		"invokeId": "string",
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
				"arrrr": "21",
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
	fmt.Println(m)
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
	v, _ := types.ListValueFrom(context.TODO(), types.ListType{ElemType: types.Int64Type}.ElemType, data["new"].([]interface{}))
	dto.New = v
	dto.Productid = types.StringValue(data["productid"].(string))
	tempProduct := data["product"].(map[string]interface{})
	tt, err := convertMapToObject(context.TODO(), tempProduct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tt)
	attrTypes := types.ObjectType{AttrTypes: map[string]attr.Type{
		"actionName":         types.StringType,
		"disabled":           types.BoolType,
		"domainCode":         types.StringType,
		"invokeId":           types.StringType,
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
				"arrrr": types.StringType,
				"kkkkk": types.StringType,
			}},
			"newnew": types.StringType,
		}}},
	}}

	// o := types.ObjectUnknown(attrTypes)

	p, diags := types.ObjectValueFrom(context.TODO(), attrTypes.AttrTypes, tt)
	fmt.Println(diags.Errors())
	dto.Product = p
	// dto.Product = diagOff(types.ObjectValueFrom, context.TODO(), attrTypes, tempProduct)
	// dto.Product = diagOff(types.ObjectValueFrom, context.TODO(), types.ObjectType{AttrTypes: map[string]attr.Type{
	// 	"actionName":         types.StringType,
	// 	"disabled":           types.BoolType,
	// 	"domainCode":         types.StringType,
	// 	"invokeId":           types.StringType,
	// 	"isDeleted":          types.BoolType,
	// 	"isPublished":        types.BoolType,
	// 	"modTime":            types.StringType,
	// 	"modifier":           types.StringType,
	// 	"permission":         types.StringType,
	// 	"productDescription": types.StringType,
	// 	"productId":          types.StringType,
	// 	"productName":        types.StringType,
	// 	"subscriptionCode":   types.StringType,
	// 	"tenantId":           types.StringType,
	// }}.AttrTypes, tempProduct)

	// "testNestedArray": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{

	// 	"kek": types.ObjectType{AttrTypes: map[string]attr.Type{
	// 		"arrrr": types.StringType,
	// 		"kkkkk": types.StringType,
	// 	}},

	// 	"newnew": types.StringType,
	// },
	// },
	// }}}.AttributeTypes()
	// tempProduct)

	return &dto
}

func diagOff[V, T interface{}](input func(ctx context.Context, elementType T, elements any) (V, diag.Diagnostics), ctx context.Context, elementType T, elements any) V {
	var emptyReturn V

	v, diags := input(ctx, elementType, elements)
	fmt.Println(diags)
	if diags.HasError() {
		diags.AddError("REFRESING ERROR", "invalid diagOff operation")
		return emptyReturn
	}

	return v
}
