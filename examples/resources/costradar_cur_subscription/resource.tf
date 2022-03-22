resource "costradar_cur_subscription" "this" {
  report_name        = "cur-report-name"
  bucket_name        = "s3-report-bucket"
  bucket_region      = "eu-central-1"
  bucket_path_prefix = "bucket/prefix"
  time_unit          = "hour"

  source_topic_arn = "arn:aws:sns:<region>:<account-id>:<costradar-topic>"

  access_config {
    reader_mode              = "assumeRole"
    assume_role_arn          = "arn:aws:iam::<account>:role/<role-name-with-path>"
    assume_role_external_id  = "1I8d2lI9D"
    assume_role_session_name = "CostradarSession"
  }
}