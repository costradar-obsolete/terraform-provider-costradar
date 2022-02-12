package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccUser(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_user.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Daniel"),
					resource.TestCheckResourceAttr(resourceName, "email", "terraform@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "initials", "D"),
					resource.TestCheckResourceAttr(resourceName, "tags.name", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.tenant", "costradar"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccUserUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Daniel van der Wel"),
					resource.TestCheckResourceAttr(resourceName, "email", "terraform@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "initials", "DV"),
					resource.TestCheckResourceAttr(resourceName, "icon_url", "https://test.com"),
					resource.TestCheckResourceAttr(resourceName, "tags.name", "updated tag"),
				),
			},
		},
	})
}

func testAccUserTF() string {
	return `
	  resource "costradar_user" "test" {
	  	email = "terraform@gmail.com"
		name = "Daniel"
		initials = "D"
		
		tags = {
			name   = "test"
			tenant = "costradar"
		}
	  }`
}

func testAccUserUpdateTF() string {
	return `
	  resource "costradar_user" "test" {
		email = "terraform@gmail.com"
		name = "Daniel van der Wel"
		initials = "DV"
		icon_url = "https://test.com"

		tags = {
			name   = "updated tag"
		}
	  }`
}
