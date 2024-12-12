package ncloud

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/NaverCloudPlatform/terraform-plugin-codegen-framework/internal/util"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
)

// To generate actual data, extract data from config.yml and code-spec.json, and render code for each receiver based on that data.
// New(): Extracts the data needed for code generation. Currently, it extracts data from config.yml and code-spec.json, but it is planned to unify everything into code-spec.json in the future.
// RenderInitial(): Generates small code blocks needed initially.
// RenderCreate(): Generates the Create function.
// RenderRead(): Generates the Read function.
// RenderUpdate(): Generates the Update function.
// RenderDelete(): Generates the Delete function.
// Calculates the necessary data during initialization and performs rendering for each method.
type Template struct {
	spec              util.NcloudSpecification
	providerName      string
	resourceName      string
	importStateLogic  string
	refreshObjectName string
	model             string
	refreshLogic      string
	endpoint          string
	deletePathParams  string
	updatePathParams  string
	readPathParams    string
	createPathParams  string
	deleteMethod      string
	updateMethod      string
	readMethod        string
	createMethod      string
	createReqBody     string
	updateReqBody     string
	idGetter          string
	funcMap           template.FuncMap
}

func (t *Template) RenderInitial() []byte {
	var b bytes.Buffer

	initialTemplate, err := template.New("").Funcs(t.funcMap).Parse(InitialTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering initial: %v", err)
	}

	data := struct {
		ProviderName string
		ResourceName string
	}{
		ProviderName: t.providerName,
		ResourceName: t.resourceName,
	}

	err = initialTemplate.ExecuteTemplate(&b, "Initial", data)
	if err != nil {
		log.Fatalf("error occurred with generating Initial template: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderImportState() []byte {
	var b bytes.Buffer

	initialTemplate, err := template.New("").Funcs(t.funcMap).Parse(ImportStateTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering initial: %v", err)
	}

	data := struct {
		ResourceName     string
		ImportStateLogic string
	}{
		ResourceName:     t.resourceName,
		ImportStateLogic: t.importStateLogic,
	}

	err = initialTemplate.ExecuteTemplate(&b, "ImportState", data)
	if err != nil {
		log.Fatalf("error occurred with generating ImportState template: %v", err)
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
		ResourceName      string
		RefreshObjectName string
		CreateReqBody     string
		CreateMethod      string
		Endpoint          string
		CreatePathParams  string
		IdGetter          string
	}{
		ResourceName:      t.resourceName,
		RefreshObjectName: t.refreshObjectName,
		CreateReqBody:     t.createReqBody,
		CreateMethod:      t.createMethod,
		Endpoint:          t.endpoint,
		CreatePathParams:  t.createPathParams,
		IdGetter:          t.idGetter,
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
		ResourceName      string
		RefreshObjectName string
		ReadPathParams    string
	}{
		ResourceName:      t.resourceName,
		RefreshObjectName: t.refreshObjectName,
		ReadPathParams:    t.readPathParams,
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
		ResourceName      string
		RefreshObjectName string
		UpdateReqBody     string
		UpdateMethod      string
		Endpoint          string
		UpdatePathParams  string
		ReadPathParams    string
	}{
		ResourceName:      t.resourceName,
		RefreshObjectName: t.refreshObjectName,
		UpdateReqBody:     t.updateReqBody,
		UpdateMethod:      t.updateMethod,
		Endpoint:          t.endpoint,
		UpdatePathParams:  t.updatePathParams,
		ReadPathParams:    t.readPathParams,
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
		ResourceName      string
		RefreshObjectName string
		DeleteMethod      string
		Endpoint          string
		DeletePathParams  string
		IdGetter          string
	}{
		ResourceName:      t.resourceName,
		RefreshObjectName: t.refreshObjectName,
		DeleteMethod:      t.deleteMethod,
		Endpoint:          t.endpoint,
		DeletePathParams:  t.deletePathParams,
		IdGetter:          t.idGetter,
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
		RefreshObjectName string
		Model             string
	}{
		RefreshObjectName: t.refreshObjectName,
		Model:             t.model,
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
		ResourceName      string
		RefreshObjectName string
		RefreshLogic      string
		ReadMethod        string
		Endpoint          string
		ReadPathParams    string
	}{
		ResourceName:      t.resourceName,
		RefreshObjectName: t.refreshObjectName,
		RefreshLogic:      t.refreshLogic,
		ReadMethod:        t.readMethod,
		Endpoint:          t.endpoint,
		ReadPathParams:    t.readPathParams,
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
		ReadMethod        string
		Endpoint          string
		ReadPathParams    string
		RefreshObjectName string
	}{
		ReadMethod:        t.readMethod,
		Endpoint:          t.endpoint,
		ReadPathParams:    t.readPathParams,
		RefreshObjectName: t.refreshObjectName,
	}

	err = waitTemplate.ExecuteTemplate(&b, "Wait", data)
	if err != nil {
		log.Fatalf("error occurred with Generating wait: %v", err)
	}

	return b.Bytes()
}

func (t *Template) RenderTest() []byte {
	var b bytes.Buffer

	testTemplate, err := template.New("").Funcs(t.funcMap).Parse(TestTemplate)
	if err != nil {
		log.Fatalf("error occurred with baseTemplate at rendering test: %v", err)
	}

	data := struct {
		ProviderName      string
		ResourceName      string
		RefreshObjectName string
		ReadMethod        string
		Endpoint          string
		ReadPathParams    string
	}{
		ProviderName:      t.providerName,
		ResourceName:      t.resourceName,
		RefreshObjectName: t.refreshObjectName,
		ReadMethod:        t.readMethod,
		Endpoint:          t.endpoint,
		ReadPathParams:    t.readPathParams,
	}

	err = testTemplate.ExecuteTemplate(&b, "Test", data)
	if err != nil {
		log.Fatalf("error occurred with Generating test: %v", err)
	}

	return b.Bytes()
}

// 초기화를 통해 필요한 데이터들을 미리 계산한다.
func New(spec util.NcloudSpecification, resourceName string) *Template {
	var refreshObjectName string
	var id string
	var attributes resource.Attributes
	var createReqBody string
	var updateReqBody string
	var importStateOverride string
	var targetResourceRequest util.RequestWithRefreshObjectName

	t := &Template{
		spec:         spec,
		resourceName: resourceName,
	}

	funcMap := util.CreateFuncMap()

	for _, resource := range spec.Resources {
		if resource.Name == resourceName {
			refreshObjectName = resource.RefreshObjectName
			id = resource.Id
			attributes = resource.Schema.Attributes
			importStateOverride = resource.ImportStateOverride
		}
	}

	for _, val := range spec.Requests {
		if val.Name == resourceName {
			targetResourceRequest = val
		}
	}

	refreshLogic, model, err := Gen_ConvertOAStoTFTypes(attributes)
	if err != nil {
		log.Fatalf("error occurred with Gen_ConvertOAStoTFTypes: %v", err)
	}

	for _, val := range targetResourceRequest.Create.RequestBody.Required {
		createReqBody = createReqBody + fmt.Sprintf(`"%[1]s": clearDoubleQuote(plan.%[2]s.String()),`, val, util.FirstAlphabetToUpperCase(val)) + "\n"
	}

	for _, val := range targetResourceRequest.Update[0].RequestBody.Required {
		updateReqBody = updateReqBody + fmt.Sprintf(`"%[1]s": clearDoubleQuote(plan.%[2]s.String()),`, val, util.FirstAlphabetToUpperCase(val)) + "\n"
	}

	t.funcMap = funcMap
	t.providerName = spec.Provider.Name
	t.refreshObjectName = refreshObjectName
	t.importStateLogic = MakeImportStateLogic(importStateOverride)
	t.model = model
	t.refreshLogic = refreshLogic
	t.endpoint = spec.Provider.Endpoint
	t.deletePathParams = extractPathParams(targetResourceRequest.Delete.Path)
	t.updatePathParams = extractPathParams(targetResourceRequest.Update[0].Path)
	t.readPathParams = extractReadPathParams(targetResourceRequest.Read.Path)
	t.createPathParams = extractPathParams(targetResourceRequest.Create.Path)
	t.deleteMethod = targetResourceRequest.Delete.Method
	t.updateMethod = targetResourceRequest.Update[0].Method
	t.readMethod = targetResourceRequest.Read.Method
	t.createMethod = targetResourceRequest.Create.Method
	t.createReqBody = createReqBody
	t.updateReqBody = updateReqBody
	t.idGetter = makeIdGetter(id)

	return t
}

func extractPathParams(path string) string {
	parts := strings.Split(path, "/")
	s := ``

	for idx, val := range parts {

		if len(val) < 1 {
			continue
		}

		s = s + `+"/"+`

		start := strings.Index(val, "{")

		// if val doesn't wrapped with curly brace
		if start == -1 {
			s = s + fmt.Sprintf(`"%s"`, val)
		} else {
			if idx == len(parts)-1 {
				s = s + `clearDoubleQuote(plan.ID.String())`
			} else {
				s = s + fmt.Sprintf(`clearDoubleQuote(plan.%s.String())`, util.PathToPascal(val))
			}
		}
	}

	return s
}

func extractReadPathParams(path string) string {
	parts := strings.Split(path, "/")
	s := ``

	for idx, val := range parts {
		if len(val) < 1 {
			continue
		}

		if idx == len(parts)-1 {
			continue
		}

		s = s + `+"/"+`

		start := strings.Index(val, "{")

		// if val doesn't wrapped with curly brace
		if start == -1 {
			s = s + fmt.Sprintf(`"%s"`, val)
		} else {
			s = s + fmt.Sprintf(`clearDoubleQuote(plan.%s.String())`, util.PathToPascal(val))
		}
	}

	return s
}

func makeIdGetter(target string) string {
	parts := strings.Split(target, ".")
	s := "response"

	for idx, val := range parts {
		if idx == len(parts)-1 {
			s = s + fmt.Sprintf(`["%s"].(string)`, util.ToCamelCase(val))
			continue
		}

		s = s + fmt.Sprintf(`["%s"].(map[string]interface{})`, util.ToCamelCase(val))
	}

	return s
}

func MakeImportStateLogic(target string) string {
	parts := strings.Split(target, ".")

	if len(parts) < 2 {
		return `resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)` + "\n"
	}

	s := `parts := strings.Split(req.ID, ".")` + "\n"
	for idx, val := range parts {
		s = s + fmt.Sprintf(`resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("%s"), parts[%d])...)`, util.ToLowerCase(util.PathToPascal(val)), idx) + "\n"
	}

	return s
}
