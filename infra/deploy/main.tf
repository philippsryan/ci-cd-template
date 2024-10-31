terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.23.0"
    }
  }
  backend "s3" {
    bucket               = "ci-cd-template-tf-state"
    key                  = "tf-state-deploy"
    region               = "us-east-2"
    encrypt              = true
    workspace_key_prefix = "tf-state-deploy-env"
  }
}

provider "aws" {
  region = "us-east-2"
  default_tags {
    tags = {
      Environemnt = terraform.workspace
      Project     = var.project
      ManageBy    = "Terraform/Deploy"
    }
  }
}

locals {
  prefix = "${var.prefix}-${terraform.workspace}"
}

data "aws_region" "current" {

}
