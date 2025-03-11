terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.90.1"
    }
  }
}

variable name {
    type = string
}
variable email {
    type = string
}

resource "aws_organizations_account" "account" {
  name  = var.name
  email = var.email
}
