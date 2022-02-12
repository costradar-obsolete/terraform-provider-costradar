package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccUserIdentitySet(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_user_identity_set.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserIdentitySetTF(),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckTypeSetElemNestedAttrs(resourceName, "workload_resource.0.*", map[string]string{
					//	"service_vendor": "aws",
					//	"resource_id": "xxx:yyy:zzz",
					//}),
					resource.TestCheckResourceAttr(resourceName, "user_identity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_identity.0.service_vendor", "aws"),
					resource.TestCheckResourceAttr(resourceName, "user_identity.0.identity", "xxx:yyy:zzz"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccUserIdentitySetUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "user_identity.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "user_identity.0.service_vendor", "aws"),
					resource.TestCheckResourceAttr(resourceName, "user_identity.0.identity", "xxx:yyy:zzz"),
					resource.TestCheckResourceAttr(resourceName, "user_identity.1.service_vendor", "aws"),
					resource.TestCheckResourceAttr(resourceName, "user_identity.1.identity", "yyy:zzz:xxx"),
				),
			},
		},
	})
}

func testAccUserIdentitySetTF() string {
	return `
	  resource "costradar_user" "test" {
		email = "terraform@gmail.com"
		name = "Daniel van der Wel"
	  }
   	  
	  resource "costradar_user_identity_set" "test" {
	  	user_id = costradar_user.test.id
	  	
		user_identity {
			service_vendor = "aws"
			identity    = "xxx:yyy:zzz"
		}
	  }
`
}

func testAccUserIdentitySetUpdateTF() string {
	return `
	  resource "costradar_user" "test" {
		email = "terraform@gmail.com"
		name = "Daniel van der Wel"
	  }

	  resource "costradar_user_identity_set" "test" {
	  	user_id = costradar_user.test.id

		user_identity {
			service_vendor = "aws"
			identity    = "xxx:yyy:zzz"
		}

		user_identity {
			service_vendor = "aws"
			identity    = "yyy:zzz:xxx"
		}
	  }
`
}
