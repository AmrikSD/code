terraform {
  required_version = ">= 1.5.7"
  required_providers {
    flux = {
      source  = "fluxcd/flux"
      version = "1.6.4"
    }
    github = {
      source  = "integrations/github"
      version = "6.6"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "4.1"
    }
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.2.0"
    }
  }
}

data "sops_file" "kubernetes-secret" {
  source_file = "${path.module}/k8s.sops.yaml"
}

provider "github" {
  owner = data.sops_file.kubernetes-secret.data["github.owner"]
  token = data.sops_file.kubernetes-secret.data["github.token"]
}

provider "flux" {
  kubernetes = {
    host                   = data.sops_file.kubernetes-secret.data["kubernetes.host"]
    client_certificate     = base64decode(data.sops_file.kubernetes-secret.data["kubernetes.client-certificate-data"])
    client_key             = base64decode(data.sops_file.kubernetes-secret.data["kubernetes.client-key-data"])
    cluster_ca_certificate = base64decode(data.sops_file.kubernetes-secret.data["kubernetes.certificate-authority-data"])
  }
  git = {
    url = ("ssh://git@github.com/${var.github_org}/${var.github_repository}.git")
    ssh = {
      username    = "git"
      private_key = tls_private_key.flux.private_key_pem
    }
  }
}
