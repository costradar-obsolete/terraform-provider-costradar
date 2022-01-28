package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccIntegrationConfig(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "data.costradar_integration_config.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataIntegrationConfigTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "integration_role_arn", "arn:aws:iam::016720764197:role/dev-integration-lambda"),
					resource.TestCheckResourceAttr(resourceName, "cur_sqs_arn", "arn:aws:sqs:eu-central-1:016720764197:dev-integration-cur"),
					resource.TestCheckResourceAttr(resourceName, "cur_sqs_url", "https://sqs.eu-central-1.amazonaws.com/016720764197/dev-integration-cur"),
					resource.TestCheckResourceAttr(resourceName, "cloudtrail_sqs_arn", "arn:aws:sqs:eu-central-1:016720764197:dev-integration-cloudtrail"),
					resource.TestCheckResourceAttr(resourceName, "cloudtrail_sqs_url", "https://sqs.eu-central-1.amazonaws.com/016720764197/dev-integration-cloudtrail"),
				),
			},
		},
	})
}

func testAccDataIntegrationConfigTF() string {
	return `
	  data "costradar_integration_config" "test" {}
	`
}
