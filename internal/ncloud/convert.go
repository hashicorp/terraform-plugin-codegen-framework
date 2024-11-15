package ncloud

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-spec/resource"
)

// generate converter that convert openapi.json schema to terraform type
func Gen_ConvertOAStoTFTypes(data resource.Attributes) (string, string, error) {
	// var data resource.Attributes
	var s string
	var m string
	// err := json.Unmarshal(jsonData, &data)
	// if err != nil {
	// return "", "", fmt.Errorf("error numb 1: %v", err)
	// }

	for _, val := range data {
		// field := fieldValue.(map[string]interface{})
		// fieldType := field["type"].(string)
		n := val.Name
		fmt.Println(n)

		if val.String != nil {
			s = s + fmt.Sprintf(`dto.%[1]s = types.StringValue(data["%[2]s"].(string))`, CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
			m = m + fmt.Sprintf("%[1]s         types.String `json:\"%[2]s\"`", CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
		} else if val.Bool != nil {
			s = s + fmt.Sprintf(`dto.%[1]s = types.BoolValue(data["%[2]s"].(bool))`, CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
			m = m + fmt.Sprintf("%[1]s         types.Bool `json:\"%[2]s\"`", CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
		} else if val.Int64 != nil {
			s = s + fmt.Sprintf(`dto.%[1]s = types.Int64Value(data["%[2]s"].(bool))`, CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
			m = m + fmt.Sprintf("%[1]s         types.Int64 `json:\"%[2]s\"`", CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
		} else if val.List != nil {
			if val.List.ElementType.String != nil {
				s = s + fmt.Sprintf(`"%[1]s": types.ListType{ElemType: types.StringType},`, n) + "\n"
			} else if val.List.ElementType.Bool != nil {
				s = s + fmt.Sprintf(`"%[1]s": types.ListType{ElemType: types.BoolType},`, n) + "\n"
			}
		} else if val.ListNested != nil {
			s = s + fmt.Sprintf(`
			temp%[1]s := data["%[2]s"].([]interface{})
			dto.%[1]s = diagOff(types.ListValueFrom, context.TODO(), types.ListType{ElemType:
				%[3]s
			}.ElementType(), temp%[1]s)`, CamelToPascalCase(n), PascalToSnakeCase(n), GenArray(val.ListNested.NestedObject.Attributes, n)) + "\n"
			m = m + fmt.Sprintf("%[1]s         types.List `json:\"%[2]s\"`", CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
		} else if val.SingleNested != nil {
			s = s + fmt.Sprintf(`
			temp%[1]s := data["%[2]s"].([]interface{})
			dto.%[1]s = diagOff(types.ObjectValueFrom, context.TODO(), types.ObjectType{AttrTypes: map[string]attr.Type{
				%[3]s
			}}}.AttributeTypes(), temp%[1]s)`, CamelToPascalCase(n), PascalToSnakeCase(n), GenObject(val.SingleNested.Attributes, n)) + "\n"
			m = m + fmt.Sprintf("%[1]s         types.Object `json:\"%[2]s\"`", CamelToPascalCase(n), PascalToSnakeCase(n)) + "\n"
		}

	}

	return s, m, nil
}

func PascalToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func CamelToPascalCase(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func GenArray(d resource.Attributes, pName string) string {
	var r string
	var s string
	var t string

	for _, val := range d {
		n := val.Name

		if val.String != nil {
			t = t + fmt.Sprintf(`"%[1]s": types.StringType,`, n) + "\n"
		} else if val.Bool != nil {
			t = t + fmt.Sprintf(`"%[1]s": types.BoolType,`, n) + "\n"
		} else if val.Int64 != nil {
			t = t + fmt.Sprintf(`"%[1]s": types.Int64Type,`, n) + "\n"
		} else if val.SingleNested != nil {
			s = s + fmt.Sprintf(`
			"%[1]s": types.ObjectType{AttrTypes: map[string]attr.Type{
				%[2]s
			}},`, n, GenObject(val.SingleNested.Attributes, n)) + "\n"
		}
	}

	r = r + fmt.Sprintf(`
	types.ObjectType{AttrTypes: map[string]attr.Type{
		%[1]s
		%[2]s
	},`, s, t)

	return r
}

func GenObject(d resource.Attributes, pName string) string {
	var s string

	for _, val := range d {
		n := val.Name

		if val.String != nil {
			s = s + fmt.Sprintf(`"%[1]s": types.StringType,`, n) + "\n"
		} else if val.Bool != nil {
			s = s + fmt.Sprintf(`"%[1]s": types.BoolType,`, n) + "\n"
		} else if val.Int64 != nil {
			s = s + fmt.Sprintf(`"%[1]s": types.Int64Type,`, n) + "\n"
		} else if val.List != nil {
			if val.List.ElementType.String != nil {
				s = s + fmt.Sprintf(`"%[1]s": types.ListType{ElemType: types.StringType},`, n) + "\n"
			} else if val.List.ElementType.Bool != nil {
				s = s + fmt.Sprintf(`"%[1]s": types.ListType{ElemType: types.BoolType},`, n) + "\n"
			}
		} else if val.ListNested != nil {
			s = s + fmt.Sprintf(`
			"%[1]s": types.ListType{ElemType:
				%[2]s
			},`, n, GenArray(val.ListNested.NestedObject.Attributes, n)) + "\n"
		} else if val.SingleNested != nil {
			s = s + fmt.Sprintf(`
			"%[1]s": types.ObjectType{AttrTypes: map[string]attr.Type{
				%[2]s
			}},`, n, GenObject(val.SingleNested.Attributes, n)) + "\n"
		}
	}

	// switch p.(type) {
	// case string:
	// 	for _, fieldValue := range p.(map[string]interface{}) {
	// 		field := fieldValue.(map[string]interface{})
	// 		fieldType := field["type"].(string)

	// 		switch fieldType {
	// 		case "string":
	// 			s = s + fmt.Sprintf(`"%[1]s": types.StringType,`, pName) + "\n"
	// 		case "boolean":
	// 			s = s + fmt.Sprintf(`"%[1]s": types.BoolType,`, pName) + "\n"
	// 		}
	// 	}
	// default:
	// 	for n, fieldValue := range p.(map[string]interface{}) {
	// 		field := fieldValue.(map[string]interface{})
	// 		fieldType := field["type"].(string)

	// 		switch fieldType {
	// 		case "array":
	// 			s = s + fmt.Sprintf(`
	// 				"%[1]s": types.ListType{ElemType:
	// 					%[2]s
	// 				},`, pName, GenArray(fieldValue.(map[string]interface{})["items"].(map[string]interface{}), n)) + "\n"
	// 		case "object":
	// 			s = s + fmt.Sprintf(`
	// 			"%[1]s": types.ObjectType{AttrTypes: map[string]attr.Type{
	// 				%[2]s
	// 			}},`, n, GenObject(fieldValue.(map[string]interface{})["properties"].(map[string]interface{}), n)) + "\n"
	// 		}
	// 	}
	// }

	return s
}
