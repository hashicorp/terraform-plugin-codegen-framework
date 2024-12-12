{{ define "Refresh" }}
// Template for generating Terraform provider Refresh operation code
// Needed data is as follows.
// RefreshObjectName string
// RefreshLogic string
// ReadMethod string
// Endpoint string
// ReadPathParams string, optional

func ConvertToFrameworkTypes(data map[string]interface{}, id string, rest []interface{}) (*{{.RefreshObjectName | ToPascalCase}}Model, error) {
	var dto {{.RefreshObjectName | ToPascalCase}}Model

	dto.ID = types.StringValue(id)

    {{.RefreshLogic}}

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

func (a *{{.RefreshObjectName | ToPascalCase}}Model) refreshFromOutput(diagnostics diag.Diagnostics, plan {{.RefreshObjectName | ToPascalCase}}Model, id string, rest ...interface{}) {
	response, _ := util.MakeRequest("{{.ReadMethod}}", "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+clearDoubleQuote(id){{end}}, "")
	if response == nil {
		diagnostics.AddError("UPDATING ERROR", "response invalid")
		return
	}

	newPlan, err := ConvertToFrameworkTypes(convertKeys(response).(map[string]interface{}), id, rest)
	if err != nil {
		diagnostics.AddError("CREATING ERROR", err.Error())
		return
	}

	*a = *newPlan
}

// convertKeys recursively converts all keys in a map from camelCase to snake_case
func convertKeys(input interface{}) interface{} {
	switch v := input.(type) {
	case map[string]interface{}:
		newMap := make(map[string]interface{})
		for key, value := range v {
			// Convert the key to snake_case
			newKey := camelToSnake(key)
			// Recursively convert nested values
			newMap[newKey] = convertKeys(value)
		}
		return newMap
	case []interface{}:
		newSlice := make([]interface{}, len(v))
		for i, value := range v {
			newSlice[i] = convertKeys(value)
		}
		return newSlice
	default:
		return v
	}
}

// Convert nested map structured json into terraform object
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

func clearDoubleQuote(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(s, "\\", "", -1), "\"", "", -1), `"`, "", -1)
}

{{ end }}