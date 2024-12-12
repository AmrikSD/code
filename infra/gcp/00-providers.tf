terraform {
    required_version = ">= 1.5.7"
    required_providers {
        google = {
          source = "hashicorp/google"
          version = "6.13.0"
        }
        sops = {
          source  = "carlpett/sops"
          version = "~> 1.0.0"
        }
    }
}

data "sops_file" "gcp-secret" {
  source_file = "${path.module}/gcp.sops.yaml"
}

provider "google" {
    project = "milestone-medical"
    credentials = data.sops_file.gcp-secret.data["google.credentials"]
}
