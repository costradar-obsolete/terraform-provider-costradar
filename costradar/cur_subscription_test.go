package costradar

import (
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccCurSubscription(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_bsyz9nkv2G7l9NFCFepghgo7xrGHtFpZ")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_cur_subscription.test"
	resource.Test(t, resource.TestCase{
		//PreCheck:   func() { testAccPreCheck(t) },
		//ErrorCheck: testAccErrorCheck(t, ec2.EndpointsID),
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCurSubscriptionTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "report_name", "report_name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "bucket_name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_region", "bucket_region"),
					resource.TestCheckResourceAttr(resourceName, "bucket_path_prefix", "xxx!!"),
					resource.TestCheckResourceAttr(resourceName, "time_unit", "hour"),
					//resource.TestCheckResourceAttr(resourceName, "reader_mode", "assumeRole"),
					//resource.TestCheckResourceAttr(resourceName, "assume_role_arn", "assume_role_arn_value"),
					//resource.TestCheckResourceAttr(resourceName, "assume_role_external_id", "assume_role_external_id_value"),
					//resource.TestCheckResourceAttr(resourceName, "assume_role_session_name", "assume_role_session_name_value"),
				),
			},
		},
	})
}

func testAccCurSubscriptionTF() string {
	return `
	  resource "costradar_cur_subscription" "test" {
	  	report_name        = "report_name"
	  	bucket_name        = "bucket_name"
	  	bucket_region      = "bucket_region"
	  	bucket_path_prefix = "xxx!!"
	  	time_unit          = "hour"
	  	access_config {
		  reader_mode              = "assumeRole"
 		  assume_role_arn          = "assume_role_arn_value"
		  assume_role_external_id  = "assume_role_external_id_value"
		  assume_role_session_name = "assume_role_session_name_value"
	    }
	  }`
}