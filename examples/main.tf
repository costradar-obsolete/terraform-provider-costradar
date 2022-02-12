terraform {
  required_providers {
    costradar = {
      version = "0.1.11"
      source  = "localhost/local/costradar"
    }
  }
}

provider "costradar" {
  token    = "api_xyz_costradar"
  endpoint = "http://localhost:8000/graphql"
}

resource "costradar_cur_subscription" "test" {
  report_name        = "report-name"
  bucket_name        = "test-costradar-bucket"
  bucket_region      = "eu-central-1"
  bucket_path_prefix = null
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
  bucket_path_prefix = costradar_cur_subscription.test.bucket_path_prefix
//  bucket_path_prefix = null

  source_topic_arn = "arn:aws:sns:eu-central-1:123456789012:cloudtrail"

  access_config {
    reader_mode              = "direct"
    assume_role_arn          = "assume_role_arn_value"
    assume_role_external_id  = "assume_role_external_id_value"
    assume_role_session_name = "assume_role_session_name_value"
  }
  depends_on = [
    costradar_cur_subscription.test
  ]
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

data "costradar_integration_config" "this" {}

output "test_output" {
  value = data.costradar_integration_config.this
}

output "costradar_user_identity_resolver_config" {
  value = costradar_identity_resolver.this
}

output "cur_subscription" {
  value = costradar_cur_subscription.test
}

output "cloudtrail_subscription" {
  value = costradar_cloudtrail_subscription.test
}

output "test" {
  value = coalesce("", null, "123")
}
resource "costradar_workload" "test" {
  name = "Terraform Test Workload"
  owners = []
  tags = {}
}

resource "costradar_workload_resource_set" "test" {
  workload_id = costradar_workload.test.id

  workload_resource {
    service_vendor = "aws"
    resource_id    = "xxx:yyy:zzz"
  }

  workload_resource {
    service_vendor = "aws"
    resource_id    = "xxx:yyy:zzz1"
  }
}

resource "costradar_team" "test" {
  name = "Terraform Test Team"
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

resource "costradar_user" "test" {
  email = "terraform@gmail.com"
  name = "Daniel"
  initials = "D"

  tags = {
    name   = "test"
    tenant = "costradar"
  }
}

resource "costradar_user_identity_set" "test" {
  user_id = costradar_user.test.id

  user_identity {
    service_vendor = "aws"
    identity    = "xxx:yyy:zzz"
  }

  user_identity {
    service_vendor = "aws"
    identity    = "yyy:zzz:xxx"
  }
}
//resource "costradar_tenant" "test" {
//  alias = "Costradar"
//  auth {
//    client_id = "123"
//    client_secret = "xxx"
//    server_metadata_url = "https://tenant.com"
//    client_kwargs = {
//      name = "value"
//    }
//    email_domains = ["@gmail.com"]
//  }
//
//}