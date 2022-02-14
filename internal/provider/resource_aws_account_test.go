package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccAccount(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_aws_account.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccAccountTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", "123:123"),
					resource.TestCheckResourceAttr(resourceName, "alias", "Alias"),
					resource.TestCheckResourceAttr(resourceName, "owners.0", "email1@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "owners.1", "email2@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "tags.name", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.tenant", "costradar"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAccountUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_id", "123:123"),
					resource.TestCheckResourceAttr(resourceName, "alias", "Changed Alias"),
					resource.TestCheckResourceAttr(resourceName, "owners.0", "email3@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "tags.name", "updated tag"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_arn", "assume_role_arn_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_external_id", "assume_role_external_id_value"),
					resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_session_name", "assume_role_session_name_value"),
				),
			},
		},
	})
}

func testAccAccountTF() string {
	return `
	  resource "costradar_aws_account" "test" {
	  	account_id = "123:123"
		alias = "Alias"
	  	owners = [
			"email1@gmail.com",
			"email2@gmail.com"
		]
		access_config {
			reader_mode = "direct"
		}
		tags = {
			name   = "test"
			tenant = "costradar"
		}
	  }`
}

func testAccAccountUpdateTF() string {
	return `
	  resource "costradar_aws_account" "test" {
		account_id = "123:123"
		alias = "Changed Alias"
		owners = [
			"email3@gmail.com"
		]
		access_config {
			reader_mode              = "direct"
			assume_role_arn          = "assume_role_arn_value"
   			assume_role_external_id  = "assume_role_external_id_value"
		   	assume_role_session_name = "assume_role_session_name_value"
		}
		tags = {
			name   = "updated tag"
		}
	  }`
}
