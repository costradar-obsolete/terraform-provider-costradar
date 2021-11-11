package costradar

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccIdentityResolverConfig(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_Uu94jTpg7UOw5vDFcHUsqFo0pkYoZiL8")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_identity_resolver.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityResolverConfigTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "lambda_arn", "123:xxx:yyy"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_arn", "assume_role_arn_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_external_id", "assume_role_external_id_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_session_name", "assume_role_session_name_value"),
				),
			},
		},
	})
}

func testAccIdentityResolverConfigTF() string {
	return `
	  resource "costradar_identity_resolver" "test" {
	  	lambda_arn = "123:xxx:yyy"
	  	access_config {
		  reader_mode              = "direct"
 		  assume_role_arn          = "assume_role_arn_value"
		  assume_role_external_id  = "assume_role_external_id_value"
		  assume_role_session_name = "assume_role_session_name_value"
	    }
	  }`
}
