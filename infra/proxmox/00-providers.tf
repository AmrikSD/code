terraform {
  required_providers {
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.0.0"
    }
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.55.0"
    }
  }
}

data "sops_file" "proxmox-secret" {
  source_file = "${path.module}/proxmox.sops.yaml"
}

provider "proxmox" {
  endpoint = data.sops_file.proxmox-secret.data["proxmox.endpoint"]
  username = data.sops_file.proxmox-secret.data["proxmox.username"]
  password = data.sops_file.proxmox-secret.data["proxmox.password"]
  insecure = true
}
