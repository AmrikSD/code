terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.90.1"
    }
  }
}

locals {
  tags = merge({
    ManagedBy = "terraform"
    CodeRepo  = "AmrikSD/Code"
    Module    = "aws-account"
  },var.tags)
}

resource "aws_organizations_account" "account" {
  name  = var.name
  email = var.email
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${aws_organizations_account.account.id}:root"]
    }
  }
}

provider "aws" {
  alias  = "new_account"
  region = var.region
  secret_key = var.secret_key
  access_key = var.access_key
  assume_role {
    role_arn = "arn:aws:iam::${aws_organizations_account.account.id}:role/OrganizationAccountAccessRole"
  }
}

resource "aws_iam_user" "child_account" {
    provider = aws.new_account
    name = format("%s-terraform", var.name)
    path = format("/%s/", var.name)
    tags = local.tags
}

data "aws_iam_policy_document" "child_iam_policy_document" {
    provider = aws.new_account
    statement {
        sid = "1"
        actions = [
          "s3:*"
        ]
        resources = [
          "arn:aws:s3:::${var.name}-terraform",
          "arn:aws:s3:::${var.name}-terraform/*"
        ]
    }
}

resource "aws_iam_policy" "child_iam_policy" {
  provider = aws.new_account
  name   = format("%s-terraform-policy", var.name)
  path   = format("/%s/", var.name)
  policy = data.aws_iam_policy_document.child_iam_policy_document.json
}

resource "aws_iam_user_policy_attachment" "child_policy_attach" {
  provider = aws.new_account
  user       = aws_iam_user.child_account.name
  policy_arn = aws_iam_policy.child_iam_policy.arn
}

resource "aws_s3_bucket" "account_bucket" {
  provider = aws.new_account
  bucket   = "${var.name}-terraform-state"
  tags     = local.tags

  depends_on = [aws_organizations_account.account]
}

resource "aws_s3_bucket_versioning" "state_versioning" {
  provider = aws.new_account
  bucket   = aws_s3_bucket.account_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "state_encryption" {
  provider = aws.new_account
  bucket   = aws_s3_bucket.account_bucket.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_iam_access_key" "terraform_user" {
  provider = aws.new_account
  user     = aws_iam_user.child_account.name
}

output "terraform_access_key" {
  value     = aws_iam_access_key.terraform_user.id
  sensitive = true
}

output "terraform_secret_key" {
  value     = aws_iam_access_key.terraform_user.secret
  sensitive = true
}
