package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccTeamMemberSet(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_team_member_set.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTeamMemberSetTF(),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckTypeSetElemNestedAttrs(resourceName, "team_member.0.*", map[string]string{
					//	"service_vendor": "aws",
					//	"resource_id": "xxx:yyy:zzz",
					//}),
					resource.TestCheckResourceAttr(resourceName, "team_member.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "team_member.0.email", "daniel@gmail.com"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTeamMemberSetUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "team_member.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "team_member.0.email", "bogdan@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "team_member.1.email", "daniel@gmail.com"),
				),
			},
		},
	})
}

func testAccTeamMemberSetTF() string {
	return `
	  resource "costradar_team" "test" {
	  	name = "Terraform Test Team"
	  }
   	  
	  resource "costradar_team_member_set" "test" {
	  	team_id = costradar_team.test.id
	  	
		team_member {
			email = "daniel@gmail.com"
		}
	  }
`
}

func testAccTeamMemberSetUpdateTF() string {
	return `
	  resource "costradar_team" "test" {
		name = "Terraform Updated Team"
	  }

	  resource "costradar_team_member_set" "test" {
	  	team_id = costradar_team.test.id
		
		team_member {
			email = "daniel@gmail.com"
		}

		team_member {
			email = "bogdan@gmail.com"
		}
	  }
`
}
