package costradar

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccCloudTrailSubscription(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_Uu94jTpg7UOw5vDFcHUsqFo0pkYoZiL8")
	os.Setenv("COSTRADAR_ENDPOINT", "http://localhost:8000/graphql")
	resourceName := "costradar_cloudtrail_subscription.test"
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"costradar": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCloudTrailSubscriptionTF(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "trail_name", "trail-name"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "bucket"),
					resource.TestCheckResourceAttr(resourceName, "source_topic_arn", "topic-arn"),
				),
			},
		},
	})
}

func testAccCloudTrailSubscriptionTF() string {
	return `
	  resource "costradar_cloudtrail_subscription" "test" {
		  trail_name         = "trail-name"
		  bucket_name        = "bucket"
		  bucket_region      = "region"
		  bucket_path_prefix = "prefix"
          source_topic_arn   = "topic-arn"
          include_global_service_events = true
          is_multi_region_trail = true
          is_organization_trail = false
		  access_config {
			reader_mode              = "assumeRole"
			assume_role_arn          = "assume_role_arn_value"
			assume_role_external_id  = "assume_role_external_id_value"
			assume_role_session_name = "assume_role_session_name_value"
		  }
		}
	`
}
