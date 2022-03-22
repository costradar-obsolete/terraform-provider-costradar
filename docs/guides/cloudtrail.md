---
subcategory: "Subscriptions"
page_title: "CloudTrail report subscription"
description: |-
An example of using CloudTrail report subscription.
---


You could create/manage a `costradar_cloudtrail_subscription` for the CloudTrail report subscription with the following config:

```terraform
resource "costradar_cloudtrail_subscription" "this" {
  trail_name         = "report-name"
  bucket_name        = "s3-report-bucket"
  bucket_region      = "eu-central-1"
  bucket_path_prefix = "bucket/prefix"

  source_topic_arn = "arn:aws:sns:<region>:<account-id>:<costradar-topic>"

  access_config {
    reader_mode              = "assumeRole"
    assume_role_arn          = "arn:aws:iam::<account>:role/<role-name-with-path>"
    assume_role_external_id  = "1I8d2lI9D"
    assume_role_session_name = "CostradarSession"
  }
}
```