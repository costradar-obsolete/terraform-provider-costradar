resource "costradar_workload" "this" {
  name = "Delivery Service"
  description = "Core delivery workload"
  owners = [
    "admin@email.com"
  ]
  tags = {
    "unit": "Delivery"
  }
}