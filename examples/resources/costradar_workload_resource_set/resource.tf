resource "costradar_workload" "this" {
  name = "Delivery Service"
}

resource "costradar_workload_resource_set" "test" {
  workload_id = costradar_workload.this.id

  workload_resource {
    service_vendor = "aws"
    resource_id    = "i-04aabf27e32c327d6"
  }

  workload_resource {
    service_vendor = "aws"
    resource_id    = "vol-0428a47e1055177bd"
  }

  workload_resource {
    service_vendor = "aws"
    resource_id    = "arn:aws:lambda:eu-central-1:123456789:function:delivery-processing"
  }
}
