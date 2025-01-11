terraform {
  required_version = ">= 1.5.7"
  required_providers {
    truenas = {
      source  = "dariusbakunas/truenas"
      version = "0.11.1"
    }
    sops = {
      source  = "carlpett/sops"
      version = "~> 1.1.0"
    }
  }
}

data "sops_file" "truenas-secret" {
  source_file = "${path.module}/truenas.sops.yaml"
}

provider "truenas" {
  api_key  = data.sops_file.truenas-secret.data["truenas.api_key"]
  base_url = format("https://%s/api/v2.0", data.sops_file.truenas-secret.data["truenas.url"])
}
