resource "cloudflare_pages_project" "amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name              = "amrik-co-uk"
  production_branch = "main"

  build_config {
    build_command   = "bundle exec jekyll build"
    destination_dir = "_site"
  }

  source {
    type = "github"
    config {
      owner                         = "AmrikSD"
      repo_name                     = "Amrik.de"
      production_branch             = "main"
      deployments_enabled           = true
      production_deployment_enabled = true
      preview_deployment_setting    = "custom"
      preview_branch_includes       = ["dev", "preview"]
      preview_branch_excludes       = ["main", "prod"]
    }
  }

  deployment_configs {
    production {

    }
  }
}

resource "cloudflare_pages_domain" "amrik_co_uk" {
  account_id   = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  project_name = "amrik-co-uk"
  domain       = "amrik.co.uk"
}

