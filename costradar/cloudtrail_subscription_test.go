package costradar

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

func TestAccCloudTrailSubscription(t *testing.T) {
	os.Setenv("COSTRADAR_TOKEN", "api_bsyz9nkv2G7l9NFCFepghgo7xrGHtFpZ")
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
					resource.TestCheckResourceAttr(resourceName, "source_arn", "xxx:yyy:zzz:test"),
					resource.TestCheckResourceAttr(resourceName, "subscription_arn", "aaa:bbb:ccc:test"),
					resource.TestCheckResourceAttr(resourceName, "bucket_name", "bucket_name_test"),
					resource.TestCheckResourceAttr(resourceName, "account_id", "123123_test"),
				),
			},
		},
	})
}

func testAccCloudTrailSubscriptionTF() string {
	return `
	  resource "costradar_cloudtrail_subscription" "test" {
	  source_arn         = "xxx:yyy:zzz:test"
	  subscription_arn   = "aaa:bbb:ccc:test"
	  bucket_name        = "bucket_name_test"
	  account_id         = "123123_test"
	  access_config {
		reader_mode              = "assumeRole"
		assume_role_arn          = "assume_role_arn_value"
		assume_role_external_id  = "assume_role_external_id_value"
		assume_role_session_name = "assume_role_session_name_value"
	  }
	}`
}
