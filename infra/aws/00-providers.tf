terraform {
  required_version = ">= 1.5.7"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.90.1"
    }
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.1.0"
    }
  }
}

data "sops_file" "aws-secret" {
  source_file = "${path.module}/aws.sops.yaml"
}

provider "aws" {
  region     = "eu-west-2"
  secret_key = data.sops_file.aws-secret.data["aws.secret_access_key"]
  access_key = data.sops_file.aws-secret.data["aws.access_key"]
}
