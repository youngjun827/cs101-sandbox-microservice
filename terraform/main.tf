provider "aws" {
  region = "ca-central-1"
}

variable "app_name" {
  description = "CS101 Sandbox Microservice"
  default     = "sandbox-lambda-api"
}

variable "app_env" {
  description = "Application environment tag"
  default     = "prod"
}

locals {
  app_id = "${lower(var.app_name)}-${lower(var.app_env)}-${random_id.unique_suffix.hex}"
}

data "archive_file" "lambda_zip" {
  type        = "zip"
  source_file = "../build/bin/app"
  output_path = "../build/bin/app.zip"
}

resource "random_id" "unique_suffix" {
  byte_length = 2
}
