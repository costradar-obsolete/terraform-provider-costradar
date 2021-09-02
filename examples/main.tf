terraform {
  required_providers {
    costradar = {
      version = ">= 0.1"
      source  = "localhost/local/costradar"
    }
  }
}

provider "costradar" {
//  token   = "api_bsyz9nkv2G7l9NFCFepghgo7xrGHtFpZ"
  endpoint = "http://localhost:8000/graphql"
}

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
}


output "cur_subscription" {
  value = costradar_cur_subscription.test
}
