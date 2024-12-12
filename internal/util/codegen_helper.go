package util

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// For curl request
func makeSignature(method, url, timestamp, accessKey, secretKey string) string {
	message := fmt.Sprintf("%s %s\n%s\n%s",
		method,
		url,
		timestamp,
		accessKey,
	)

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Convert nested map structured json into terraform object
func ConvertMapToObject(ctx context.Context, data map[string]interface{}) (types.Object, error) {
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

// Convert interface{} into attr.Type attr.Value
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
			// Treat as array list in case of empty
			return types.ListType{ElemType: types.StringType},
				types.ListValueMust(types.StringType, []attr.Value{}),
				nil
		}
		// Determine type based on first element
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
		objValue, err := ConvertMapToObject(ctx, v)
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

func DiagOff[V, T interface{}](input func(ctx context.Context, elementType T, elements any) (V, diag.Diagnostics), ctx context.Context, elementType T, elements any) V {
	var emptyReturn V

	v, diags := input(ctx, elementType, elements)

	if diags.HasError() {
		diags.AddError("REFRESING ERROR", "invalid diagOff operation")
		return emptyReturn
	}

	return v
}

// convertKeys recursively converts all keys in a map from camelCase to snake_case
func ConvertKeys(input interface{}) interface{} {
	switch v := input.(type) {
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, value := range v {
			// Convert the key to snake_case
			newKey := camelToSnake(key)
			// Recursively convert nested values
			newMap[newKey] = ConvertKeys(value)
		}
		return newMap
	case []interface{}:
		newSlice := make([]interface{}, len(v))
		for i, value := range v {
			newSlice[i] = ConvertKeys(value)
		}
		return newSlice
	default:
		return v
	}
}

func camelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

func MakeIdGetter(target string) string {
	s := "response"
	parts := strings.Split(target, ".")

	for idx, val := range parts {
		if idx == len(parts)-1 {
			s = s + fmt.Sprintf(`["%s"].(string)`, ToCamelCase(val))
			continue
		}

		s = s + fmt.Sprintf(`["%s"].(map[string]interface{})`, ToCamelCase(val))
	}

	return s
}

func ClearDoubleQuote(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(s, "\\", "", -1), "\"", "", -1), `"`, "", -1)
}

func MakeRequest(method, path, endpoint, reqBody string) (map[string]interface{}, error) {
	access_key := os.Getenv("NCLOUD_ACCESS_KEY")
	secret_key := os.Getenv("NCLOUD_SECRET_KEY")
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	signature := makeSignature(method, ExtractPath(endpoint), timestamp, access_key, secret_key)
	b := bytes.NewBuffer([]byte(reqBody))

	req, err := http.NewRequest(method, endpoint, b)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-ncp-apigw-timestamp", timestamp)
	req.Header.Add("x-ncp-iam-access-key", access_key)
	req.Header.Add("x-ncp-apigw-signature-v2", signature)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("pragma", "no-cache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}
