terraform {
  required_version = ">= 1.5.7"
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4.52.0"
    }
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.2.0"
    }
  }
}

data "sops_file" "cloudflare-secret" {
  source_file = "${path.module}/cloudflare.sops.yaml"
}

provider "cloudflare" {
  api_token = data.sops_file.cloudflare-secret.data["cloudflare.api_key"]
}
