terraform {
  required_providers {
    costradar = {
      version = "0.1.4"
      source  = "localhost/local/costradar"
    }
  }
}

provider "costradar" {
  token    = "api_Uu94jTpg7UOw5vDFcHUsqFo0pkYoZiL8"
  endpoint = "http://localhost:8000/graphql"
}

resource "costradar_cur_subscription" "test" {
  report_name        = "report-name"
  bucket_name        = "test-costradar-bucket"
  bucket_region      = "eu-central-1"
  bucket_path_prefix = "prefix"
  time_unit          = "hour"

  source_topic_arn = "arn:aws:sns:eu-central-1:123456789012:cur"

  access_config {
    reader_mode              = "direct"
    assume_role_arn          = "assume_role_arn_value"
    assume_role_external_id  = "assume_role_external_id_value"
    assume_role_session_name = "assume_role_session_name_value"
  }
}

resource "costradar_cloudtrail_subscription" "test" {
  trail_name         = "trail-name"
  bucket_name        = "test-costradar-bucket"
  bucket_region      = "eu-central-1"
  bucket_path_prefix = "prefix"

  source_topic_arn = "arn:aws:sns:eu-central-1:123456789012:cloudtrail"

  access_config {
    reader_mode              = "direct"
    assume_role_arn          = "assume_role_arn_value"
    assume_role_external_id  = "assume_role_external_id_value"
    assume_role_session_name = "assume_role_session_name_value"
  }
}

resource "costradar_identity_resolver" "this" {
  lambda_arn = "123:xxx:yyy"
  access_config {
    reader_mode              = "direct"
    assume_role_arn          = "assume_role_arn_value"
    assume_role_external_id  = "assume_role_external_id_value"
    assume_role_session_name = "assume_role_session_name_value"
  }
}

data "costradar_subscription_meta" "this" {}

//output "test_output" {
//  value = data.costradar_subscription_meta.this
//}

output "costradar_user_identity_resolver_config" {
  value = costradar_identity_resolver.this
}

output "cur_subscription" {
  value = costradar_cur_subscription.test
}

output "cloudtrail_subscription" {
  value = costradar_cloudtrail_subscription.test
}