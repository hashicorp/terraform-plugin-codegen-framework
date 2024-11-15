package ncloud

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
)

func Gen_Template() []byte {
	var b bytes.Buffer
	funcMap := util.CreateFuncMap()

	// Call go template and be ready for codegen
	schemaTemplate, err := template.New("").Funcs(funcMap).Parse(MainTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate: %v", err)
	}

	// Get absolute path of config files
	// jsonPath := util.MustAbs("./internal/apigw_v1.json")
	configPath := util.MustAbs("./internal/generator_config_apigw.yml")
	codeSpecPath := util.MustAbs("./internal/example-code-spec.json")

	attr, resourceName, _ := util.ExtractAttribute(codeSpecPath)
	if err != nil {
		log.Fatalf("error occurred with ExtractDto: %v", err)
	}

	// Extract needed information
	APIConfig, _, endpoint, err := util.ExtractConfig(configPath, resourceName)
	if err != nil {
		log.Fatalf("error occurred with ExtractConfig: %v", err)
	}

	// dtoByte, err := util.ExtractDto(jsonPath, APIConfig.DtoName)
	// if err != nil {
	// 	log.Fatalf("error occurred with ExtractDto: %v", err)
	// }
	refreshLogic, model, err := Gen_ConvertOAStoTFTypes(attr)
	if err != nil {
		log.Fatalf("error occurred with Gen_ConvertOAStoTFTypes: %v", err)
	}

	// Prepare all datas to execute template
	// Core information to make code-spec.json
	data := struct {
		ResourceName     string
		DtoName          string
		Model            string
		RefreshLogic     string
		Endpoint         string
		DeletePathParams string
		UpdatePathParams string
		ReadPathParams   string
		CreatePathParams string
		DeleteMethod     string
		UpdateMethod     string
		ReadMethod       string
		CreateMethod     string
	}{
		ResourceName:     resourceName,
		DtoName:          APIConfig.Read.Response,
		Model:            model,
		RefreshLogic:     refreshLogic,
		Endpoint:         endpoint,
		DeletePathParams: extractPathParams(APIConfig.Delete.Path),
		UpdatePathParams: extractPathParams(APIConfig.Update[0].Path),
		ReadPathParams:   extractPathParams(APIConfig.Read.Path),
		DeleteMethod:     APIConfig.Create.Method,
		UpdateMethod:     APIConfig.Update[0].Method,
		ReadMethod:       APIConfig.Read.Method,
		CreateMethod:     APIConfig.Create.Method,
	}

	// Execute go template
	err = schemaTemplate.ExecuteTemplate(&b, "Main", data)
	if err != nil {
		log.Fatalf("error occurred with Converting: %v", err)
	}

	return b.Bytes()
}

func extractPathParams(path string) string {
	start := strings.Index(path, "{")
	if start == -1 {
		return ""
	}
	end := strings.Index(path, "}")
	if end == -1 {
		return ""
	}

	param := path[start+1 : end]

	param = strings.ReplaceAll(param, "-", "_")

	return param
}
