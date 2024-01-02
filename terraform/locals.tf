locals {
  service_name = "lambda"
  src_path     = "../${local.service_name}"
  binary_name  = local.service_name
  binary_path  = "${path.module}/tf_generated/${local.binary_name}"
  archive_path = "${path.module}/tf_generated/${local.service_name}.zip"
}

