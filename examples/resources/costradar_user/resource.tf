resource "costradar_user" "this" {
  email = "user@gmail.com"
  name = "User"
  initials = "U"
  icon_url = "url://"
  tags = {
    tenant = "costradar"
  }
}