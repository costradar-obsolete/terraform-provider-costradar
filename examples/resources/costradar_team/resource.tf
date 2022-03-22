resource "costradar_team" "this" {
  name = "Costradar Team"
  description = "Core Costradar Team"
  tags = {
    tenant = "costradar"
  }
}