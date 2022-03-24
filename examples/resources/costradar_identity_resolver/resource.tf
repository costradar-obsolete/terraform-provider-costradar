resource "costradar_identity_resolver" "this" {
  lambda_arn = "arn:aws:lambda:<region>:<account-id>:function:resolver"
  access_config {
    reader_mode              = "assumeRole"
    assume_role_arn          = "arn:aws:iam::<account>:role/<role-name-with-path>"
    assume_role_external_id  = "1I8d2lI9D"
    assume_role_session_name = "CostradarSession"
  }
}

resource "costradar_identity_resolver" "direct_access" {
  lambda_arn = "arn:aws:lambda:<region>:<account-id>:function:resolver"
  access_config {
    reader_mode              = "direct"
  }
}