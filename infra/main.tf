terraform {
  backend "s3" {
    bucket       = "berries-terraform"
    key          = "terraform.tfstate"
    region       = "eu-west-2"
    use_lockfile = true
  }
  required_version = "~> 1.16.0"
  # Provider versions are pinned here, next to .terraform.lock.hcl, so
  # Renovate updates the lock file in the same commit as the bump (it
  # only refreshes lock files in the directory of the changed pin).
  required_providers {
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.4.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "6.66.0"
    }
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4.52.5"
    }
    unifi = {
      source  = "ubiquiti-community/unifi"
      version = "0.56.0"
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
