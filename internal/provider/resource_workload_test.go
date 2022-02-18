package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccWorkload(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_workload.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWorkloadTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform Test Workload"),
					//resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "owners.0", "email1@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "owners.1", "email2@gmail.com"),
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
				Config: testAccWorkloadUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Terraform Updated Workload"),
					resource.TestCheckResourceAttr(resourceName, "description", "Some description"),
					resource.TestCheckResourceAttr(resourceName, "owners.0", "email3@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "tags.name", "updated tag"),
				),
			},
		},
	})
}

func testAccWorkloadTF() string {
	return `
	  resource "costradar_workload" "test" {
	  	name = "Terraform Test Workload"
	  	owners = [
			"email1@gmail.com",
			"email2@gmail.com"
		]
		tags = {
			name   = "test"
			tenant = "costradar"
		}
	  }`
}

func testAccWorkloadUpdateTF() string {
	return `
		resource "costradar_workload" "test" {
		name = "Terraform Updated Workload"
		description = "Some description"
		owners = [
			"email3@gmail.com"
		]
		tags = {
			name   = "updated tag"
		}
	  }`
}
