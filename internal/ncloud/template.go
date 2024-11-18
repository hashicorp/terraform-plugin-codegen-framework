package ncloud

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
)

type Template struct {
	configPath       string
	codeSpecPath     string
	resourceName     string
	dtoName          string
	model            string
	refreshLogic     string
	endpoint         string
	deletePathParams string
	updatePathParams string
	readPathParams   string
	createPathParams string
	deleteMethod     string
	updateMethod     string
	readMethod       string
	createMethod     string
	createReqBody    string
	updateReqBody    string
	funcMap          template.FuncMap
}

func (t *Template) RenderInitial() []byte {
	var b bytes.Buffer

	initialTemplate, err := template.New("").Funcs(t.funcMap).Parse(InitialTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering initial: %v", err)
	}

	data := struct {
		ResourceName string
	}{
		ResourceName: t.resourceName,
	}

	err = initialTemplate.ExecuteTemplate(&b, "Initial", data)
	if err != nil {
		log.Fatalf("error occurred with generating Initial template: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderCreate() []byte {
	var b bytes.Buffer

	createTemplate, err := template.New("").Funcs(t.funcMap).Parse(CreateTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering create: %v", err)
	}

	data := struct {
		ResourceName     string
		DtoName          string
		CreateReqBody    string
		CreateMethod     string
		Endpoint         string
		CreatePathParams string
	}{
		ResourceName:     t.resourceName,
		DtoName:          t.dtoName,
		CreateReqBody:    t.createReqBody,
		CreateMethod:     t.createMethod,
		Endpoint:         t.endpoint,
		CreatePathParams: t.createPathParams,
	}

	err = createTemplate.ExecuteTemplate(&b, "Create", data)
	if err != nil {
		log.Fatalf("error occurred with Generating Create: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderRead() []byte {
	var b bytes.Buffer

	readTemplate, err := template.New("").Funcs(t.funcMap).Parse(ReadTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering read: %v", err)
	}

	data := struct {
		ResourceName string
		DtoName      string
	}{
		ResourceName: t.resourceName,
		DtoName:      t.dtoName,
	}

	err = readTemplate.ExecuteTemplate(&b, "Read", data)
	if err != nil {
		log.Fatalf("error occurred with Generating Read: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderUpdate() []byte {
	var b bytes.Buffer

	updateTemplate, err := template.New("").Funcs(t.funcMap).Parse(UpdateTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering update: %v", err)
	}

	data := struct {
		ResourceName     string
		DtoName          string
		UpdateReqBody    string
		UpdateMethod     string
		Endpoint         string
		UpdatePathParams string
	}{
		ResourceName:     t.resourceName,
		DtoName:          t.dtoName,
		UpdateReqBody:    t.updateReqBody,
		UpdateMethod:     t.updateMethod,
		Endpoint:         t.endpoint,
		UpdatePathParams: t.updatePathParams,
	}

	err = updateTemplate.ExecuteTemplate(&b, "Update", data)
	if err != nil {
		log.Fatalf("error occurred with Generating Update: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderDelete() []byte {
	var b bytes.Buffer

	deleteTemplate, err := template.New("").Funcs(t.funcMap).Parse(DeleteTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering delete: %v", err)
	}

	data := struct {
		ResourceName     string
		DtoName          string
		DeleteMethod     string
		Endpoint         string
		DeletePathParams string
	}{
		ResourceName:     t.resourceName,
		DtoName:          t.dtoName,
		DeleteMethod:     t.deleteMethod,
		Endpoint:         t.endpoint,
		DeletePathParams: t.deletePathParams,
	}

	err = deleteTemplate.ExecuteTemplate(&b, "Delete", data)
	if err != nil {
		log.Fatalf("error occurred with Generating delete: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderModel() []byte {
	var b bytes.Buffer

	modelTemplate, err := template.New("").Funcs(t.funcMap).Parse(ModelTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering model: %v", err)
	}

	data := struct {
		DtoName string
		Model   string
	}{
		DtoName: t.dtoName,
		Model:   t.model,
	}

	err = modelTemplate.ExecuteTemplate(&b, "Model", data)
	if err != nil {
		log.Fatalf("error occurred with Generating Model: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderRefresh() []byte {
	var b bytes.Buffer

	refreshTemplate, err := template.New("").Funcs(t.funcMap).Parse(RefreshTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering refresh: %v", err)
	}

	data := struct {
		DtoName        string
		RefreshLogic   string
		ReadMethod     string
		Endpoint       string
		ReadPathParams string
	}{
		DtoName:        t.dtoName,
		RefreshLogic:   t.refreshLogic,
		ReadMethod:     t.readMethod,
		Endpoint:       t.endpoint,
		ReadPathParams: t.readPathParams,
	}

	err = refreshTemplate.ExecuteTemplate(&b, "Refresh", data)
	if err != nil {
		log.Fatalf("error occurred with Generating Refresh: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderWait() []byte {
	var b bytes.Buffer

	waitTemplate, err := template.New("").Funcs(t.funcMap).Parse(WaitTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering wait: %v", err)
	}

	data := struct {
		ReadMethod     string
		Endpoint       string
		ReadPathParams string
		DtoName        string
	}{
		ReadMethod:     t.readMethod,
		Endpoint:       t.endpoint,
		ReadPathParams: t.readPathParams,
		DtoName:        t.dtoName,
	}

	err = waitTemplate.ExecuteTemplate(&b, "Wait", data)
	if err != nil {
		log.Fatalf("error occurred with Generating wait: %v", err)
	}

	return b.Bytes()
}

func New(configPath, codeSpecPath, resourceName string) *Template {
	t := &Template{
		configPath:   configPath,
		codeSpecPath: codeSpecPath,
		resourceName: resourceName,
	}

	funcMap := util.CreateFuncMap()

	t.funcMap = funcMap

	attributes, _, _, err := util.ExtractAttribute(codeSpecPath)
	if err != nil {
		log.Fatalf("error occurred with ExtractAttribute: %v", err)
	}

	// Extract needed information
	APIConfig, _, endpoint, err := util.ExtractConfig(configPath, resourceName)
	if err != nil {
		log.Fatalf("error occurred with ExtractConfig: %v", err)
	}

	refreshLogic, model, err := Gen_ConvertOAStoTFTypes(attributes)
	if err != nil {
		log.Fatalf("error occurred with Gen_ConvertOAStoTFTypes: %v", err)
	}

	targetResource := util.ExtractRequest(codeSpecPath, resourceName)

	var createReqBody string
	for _, val := range targetResource.Create.RequestBody.Required {
		createReqBody = createReqBody + fmt.Sprintf(`"%[1]s": plan.%[2]s.String(),`, util.FirstAlphabetToUpperCase(val), util.FirstAlphabetToUpperCase(val)) + "\n"
	}

	var updateReqBody string
	for _, val := range targetResource.Update[0].RequestBody.Required {
		updateReqBody = updateReqBody + fmt.Sprintf(`"%[1]s": plan.%[2]s.String(),`, util.FirstAlphabetToUpperCase(val), util.FirstAlphabetToUpperCase(val)) + "\n"
	}

	t.dtoName = APIConfig.DtoName
	t.model = model
	t.refreshLogic = refreshLogic
	t.endpoint = endpoint
	t.deletePathParams = extractPathParams(APIConfig.Delete.Path)
	t.updatePathParams = extractPathParams(APIConfig.Update[0].Path)
	t.readPathParams = extractPathParams(APIConfig.Read.Path)
	t.deleteMethod = APIConfig.Delete.Method
	t.updateMethod = APIConfig.Update[0].Method
	t.readMethod = APIConfig.Read.Method
	t.createMethod = APIConfig.Create.Method
	t.createReqBody = createReqBody
	t.updateReqBody = updateReqBody

	return t
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
