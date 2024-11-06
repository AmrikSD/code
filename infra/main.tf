terraform {
  backend "s3" {
    profile = "terraform"
    bucket  = "berries-terraform"
    key     = "terraform.tfstate"
    region  = "eu-west-2"
  }
}

module "cloudflare" {
  source = "./cloudflare/"
}

