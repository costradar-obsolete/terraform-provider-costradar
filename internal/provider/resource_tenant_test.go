package provider

//
//import (
//	//"fmt"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	"os"
//	"testing"
//)
//
//func TestAccTenant(t *testing.T) {
//	os.Setenv("COSTRADAR_TOKEN", "api_xyz_costradar")
//	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
//	resourceName := "costradar_tenant.test"
//	resource.Test(t, resource.TestCase{
//		ProviderFactories: map[string]func() (*schema.Provider, error){
//			"costradar": func() (*schema.Provider, error) {
//				return Provider(), nil
//			},
//		},
//		Steps: []resource.TestStep{
//			{
//				Config: testAccTenantTF(),
//				Check: resource.ComposeTestCheckFunc(
//					resource.TestCheckResourceAttr(resourceName, "alias", "Costradar"),
//					//resource.TestCheckResourceAttr(resourceName, "alias", "Alias"),
//					//resource.TestCheckResourceAttr(resourceName, "owners.0", "email1@gmail.com"),
//					//resource.TestCheckResourceAttr(resourceName, "owners.1", "email2@gmail.com"),
//					//resource.TestCheckResourceAttr(resourceName, "tags.name", "test"),
//					//resource.TestCheckResourceAttr(resourceName, "tags.tenant", "costradar"),
//					//resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
//				),
//			},
//			//{
//			//	ResourceName:      resourceName,
//			//	ImportState:       true,
//			//	ImportStateVerify: true,
//			//},
//			//{
//			//	Config: testAccTenantUpdateTF(),
//			//	Check: resource.ComposeTestCheckFunc(
//			//		resource.TestCheckResourceAttr(resourceName, "account_id", "123:123"),
//			//		resource.TestCheckResourceAttr(resourceName, "alias", "Changed Alias"),
//			//		resource.TestCheckResourceAttr(resourceName, "owners.0", "email3@gmail.com"),
//			//		resource.TestCheckResourceAttr(resourceName, "tags.name", "updated tag"),
//			//		resource.TestCheckResourceAttr(resourceName, "access_config.0.reader_mode", "direct"),
//			//		resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_arn", "assume_role_arn_value"),
//			//		resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_external_id", "assume_role_external_id_value"),
//			//		resource.TestCheckResourceAttr(resourceName, "access_config.0.assume_role_session_name", "assume_role_session_name_value"),
//			//	),
//			//},
//		},
//	})
//}
//
//func testAccTenantTF() string {
//	return `
//	resource "costradar_tenant" "test" {
//	  alias = "Costradar"
//	  auth {
//		client_id = "123"
//		client_secret = "xxx"
//		server_metadata_url = "https://tenant.com"
//		client_kwargs = {
//		  name = "value"
//		}
//		email_domains = ["@gmail.com"]
//	  }
//	}`
//}
//
////func testAccTenantUpdateTF() string {
////	return `
////	  resource "costradar_account" "test" {
////		account_id = "123:123"
////		alias = "Changed Alias"
////		owners = [
////			"email3@gmail.com"
////		]
////		access_config {
////			reader_mode              = "direct"
////			assume_role_arn          = "assume_role_arn_value"
////   			assume_role_external_id  = "assume_role_external_id_value"
////		   	assume_role_session_name = "assume_role_session_name_value"
////		}
////		tags = {
////			name   = "updated tag"
////		}
////	  }`
////}
