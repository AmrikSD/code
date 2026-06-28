terraform {
  backend "s3" {
    bucket       = "berries-terraform"
    key          = "terraform.tfstate"
    region       = "eu-west-2"
    use_lockfile = true
  }
  required_version = "~> 1.15.0"
  required_providers {
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.4.0"
    }
  }
}

data "sops_file" "cloudflare-secret" {
  source_file = "${path.module}/cloudflare/cloudflare.sops.yaml"
}

resource "random_bytes" "tunnel_secret" {
  length = 32
}

module "cloudflare" {
  source = "./cloudflare/"
}


module "aws" {
  source = "./aws"
}
