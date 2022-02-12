package provider

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccWorkloadSet(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_workload_resource_set.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWorkloadSetTF(),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckTypeSetElemNestedAttrs(resourceName, "workload_resource.0.*", map[string]string{
					//	"service_vendor": "aws",
					//	"resource_id": "xxx:yyy:zzz",
					//}),
					resource.TestCheckResourceAttr(resourceName, "workload_resource.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "workload_resource.0.service_vendor", "aws"),
					resource.TestCheckResourceAttr(resourceName, "workload_resource.0.resource_id", "xxx:yyy:zzz"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccWorkloadSetUpdateTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workload_resource.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "workload_resource.0.service_vendor", "aws"),
					resource.TestCheckResourceAttr(resourceName, "workload_resource.0.resource_id", "xxx:yyy:zzz"),
					resource.TestCheckResourceAttr(resourceName, "workload_resource.1.service_vendor", "aws"),
					resource.TestCheckResourceAttr(resourceName, "workload_resource.1.resource_id", "yyy:zzz:xxx"),
				),
			},
		},
	})
}

func testAccWorkloadSetTF() string {
	return `
	  resource "costradar_workload" "test" {
	  	name = "Terraform Test Workload"
	  }
   	  
	  resource "costradar_workload_resource_set" "test" {
	  	workload_id = costradar_workload.test.id
	  	
		workload_resource {
			service_vendor = "aws"
			resource_id    = "xxx:yyy:zzz"
		}
	  }
`
}

func testAccWorkloadSetUpdateTF() string {
	return `
	  resource "costradar_workload" "test" {
		name = "Terraform Updated Workload"
	  }

	  resource "costradar_workload_resource_set" "test" {
	  	workload_id = costradar_workload.test.id

		workload_resource {
			service_vendor = "aws"
			resource_id    = "xxx:yyy:zzz"
		}

		workload_resource {
			service_vendor = "aws"
			resource_id    = "yyy:zzz:xxx"
		}
	  }
`
}
