resource "costradar_aws_account" "this" {
  account_id = "123456789"
  alias = "Alias"
  access_config {
    reader_mode              = "assumeRole"
    assume_role_arn          = "arn:aws:iam::<account>:role/<role-name-with-path>"
    assume_role_external_id  = "1I8d2lI9D"
    assume_role_session_name = "CostradarSession"
  }
  owners = [
    "email1@gmail.com",
    "email2@gmail.com"
  ]
  tags = {
    tenant = "client"
  }
}

resource "costradar_aws_account" "direct_access" {
  account_id = "123456789"
  access_config {
    reader_mode = "direct"
  }
  owners = ["email1@gmail.com"]
}