terraform {
  required_version = ">= 1.5.7"
  required_providers {
    unifi = {
      source  = "paultyng/unifi"
      version = "0.41.0"
    }
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.0.0"
    }
  }
}

data "sops_file" "unifi-secret" {
  source_file = "${path.module}/unifi.sops.yaml"
}

variable "api_url" {
  type = string
}

variable "insecure" {
  type = string
}

provider "unifi" {
  username       = data.sops_file.unifi-secret.data["unifi.username"]
  password       = data.sops_file.unifi-secret.data["unifi.password"]
  api_url        = var.api_url
  allow_insecure = var.insecure # optionally use UNIFI_INSECURE env var
}
