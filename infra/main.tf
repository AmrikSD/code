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
  source        = "./cloudflare/"
  tunnel_secret = random_bytes.tunnel_secret.base64
}

module "gcp" {
  source           = "./gcp/"
  cf_account_id    = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  cf_tunnel_id     = module.cloudflare.tunnel_id
  cf_tunnel_secret = random_bytes.tunnel_secret.base64
}

module "aws" {
  source = "./aws"
}
