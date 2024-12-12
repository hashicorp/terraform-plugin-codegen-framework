{{ define "Test" }}
// Template for generating Terraform provider test code
// Needed data is as follows.
// ProviderName string
// ResourceName string
// RefreshObjectName string
// ReadMethod string
// Endpoint string
// ReadPathParams string, optional

package {{.ResourceName}}_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	. "github.com/terraform-providers/terraform-provider-ncloud/internal/acctest"
	"github.com/terraform-providers/terraform-provider-ncloud/internal/conn"
)

func TestAccResourceNcloud{{.ProviderName | ToPascalCase}}_{{.ResourceName | ToLowerCase}}_basic(t *testing.T) {
	{{.ResourceName | ToCamelCase}}Name := fmt.Sprintf("tf-{{.ResourceName | ToCamelCase}}-%s", acctest.RandString(5))

	resourceName := "ncloud_{{.ProviderName | ToLowerCase}}_{{.ResourceName | ToLowerCase}}.testing_{{.ResourceName | ToLowerCase}}"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheck{{.ResourceName | ToPascalCase}}Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAcc{{.ResourceName | ToLowerCase}}Config({{.ResourceName | ToCamelCase}}Name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheck{{.ResourceName | ToLowerCase}}Exists(resourceName, GetTestProvider(true)),
					resource.TestMatchResourceAttr(resourceName, "{{.ResourceName | ToCamelCase}}_name", {{.ResourceName | ToCamelCase}}Name),
				),
			},
		},
	})
}

func testAccCheck{{.ResourceName | ToLowerCase}}Exists(n string, provider *schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		resource, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found %s", n)
		}

		if resource.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

        response, err := util.MakeRequest("{{.ReadMethod}}", "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+clearDoubleQuote(resource.Primary.ID){{end}}, "")
        if response == nil {
            return err
        }
		if err != nil {
			return err
		}

		return fmt.Errorf("{{.ResourceName | ToCamelCase}} not found")
	}
}

func testAccCheck{{.ResourceName | ToPascalCase}}Destroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ncloud_{{.ProviderName | ToLowerCase}}_{{.ResourceName | ToLowerCase}}.testing_{{.ResourceName | ToLowerCase}}" {
			continue
		}

        response, _ := util.MakeRequest("{{.ReadMethod}}", "{{.Endpoint | ExtractPath}}", "{{.Endpoint}}"{{if .ReadPathParams}}{{.ReadPathParams}}+"/"+clearDoubleQuote(rs.Primary.ID){{end}}, "")
        if response["error"] != nil {
            return nil
        }
	}

	return nil
}

func testAcc{{.ResourceName | ToLowerCase}}Config({{.ResourceName | ToCamelCase}}Name string) string {
	return fmt.Sprintf(`
	resource "ncloud_{{.ProviderName | ToLowerCase}}_{{.ResourceName | ToLowerCase}}" "testing_{{.ResourceName | ToLowerCase}}" {
		{{.ResourceName | ToCamelCase}}_name			= "%[1]s"
	}`, {{.ResourceName | ToCamelCase}}Name)
}

func clearDoubleQuote(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(s, "\\", "", -1), "\"", "", -1), `"`, "", -1)
}

{{ end }}