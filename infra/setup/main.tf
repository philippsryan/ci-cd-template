terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.23.0"
    }
  }
  backend "s3" {
    bucket  = "ci-cd-template-tf-state"
    key     = "tf-state-setup"
    region  = "us-east-2"
    encrypt = true

  }
}

provider "aws" {
  region = "us-east-2"
  default_tags {
    tags = {
      Environemnt = terraform.workspace
      Project     = var.project
      ManageBy    = "Terraform/setup"
    }
  }
}