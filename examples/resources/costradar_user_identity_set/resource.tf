resource "costradar_user" "this" {
  email = "user@gmail.com"
  name = "User"
}

resource "costradar_user_identity_set" "this" {
  user_id = costradar_user.this.id

  user_identity {
    service_vendor = "aws"
    identity    = "arn:aws:iam::123456789:assumed-role/DevOpsEngineer/DevOpsSchmitz"
  }

  user_identity {
    service_vendor = "aws"
    identity    = "arn:aws:iam::account:user/Schmitz"
  }
}