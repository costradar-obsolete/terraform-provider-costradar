package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccTeam(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_team.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTeamTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform Test Team"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
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
				Config: testAccTeamUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform Updated Team"),
					resource.TestCheckResourceAttr(resourceName, "description", "Some description"),
					resource.TestCheckResourceAttr(resourceName, "tags.name", "updated tag"),
				),
			},
		},
	})
}

func testAccTeamTF() string {
	return `
	  resource "costradar_team" "test" {
	  	name = "Terraform Test Team"
		tags = {
			name   = "test"
			tenant = "costradar"
		}
	  }`
}

func testAccTeamUpdateTF() string {
	return `
		resource "costradar_team" "test" {
		name = "Terraform Updated Team"
		description = "Some description"
		tags = {
			name   = "updated tag"
		}
	  }`
}
