module "archive" {
  source     = "./modules/aws-account"
  name       = "archive"
  email      = "archive-admin@amrik.co.uk"
  region     = "eu-west-2"
  secret_key = data.sops_file.aws-secret.data["aws.secret_access_key"]
  access_key = data.sops_file.aws-secret.data["aws.access_key"]
}


provider "aws" {
  alias      = "archive"
  region     = "eu-west-2"
  access_key = module.archive.terraform_access_key
  secret_key = module.archive.terraform_secret_key
}


resource "aws_s3_bucket" "archive_data" {
  provider = aws.archive
  bucket   = "${module.archive.account_name}-${module.archive.account_id}-wedding-photos"
  tags = {
    ManagedBy = "terraform"
    CodeRepo  = "AmrikSD/Code"
    Account   = "archive"
  }
}

