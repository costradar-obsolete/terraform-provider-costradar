package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccSubscriptionMeta(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_Uu94jTpg7UOw5vDFcHUsqFo0pkYoZiL8")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "data.costradar_subscription_meta.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"provider": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSubscriptionMetaTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cur_sqs_arn", "arn:aws:sqs:eu-central-1:123456789012:cur_queue"),
				),
			},
		},
	})
}

func testAccDataSubscriptionMetaTF() string {
	return `
	  data "costradar_subscription_meta" "test" {}
	`
}
