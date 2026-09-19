resource "cloudflare_pages_project" "osrs_amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name              = "osrs-amrik-co-uk"
  production_branch = "main"

  build_config {
    build_command   = ""
    root_dir        = "projects/osrs.amrik.co.uk"
    destination_dir = "dist"
  }

  source {
    type = "github"
    config {
      owner                         = "AmrikSD"
      repo_name                     = "code"
      production_branch             = "main"
      deployments_enabled           = true
      production_deployment_enabled = true
      preview_deployment_setting    = "custom"
      preview_branch_includes       = ["dev", "preview"]
      preview_branch_excludes       = ["main", "prod"]
      pr_comments_enabled           = true
    }
  }

  deployment_configs {
    production {

    }
  }
}

resource "cloudflare_pages_domain" "osrs_subdomain" {
  account_id   = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  project_name = "osrs-amrik-co-uk"
  domain       = "osrs.amrik.co.uk"
}

resource "cloudflare_record" "osrs_amrik_co_uk" {
  zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.zone_id"]
  name    = "osrs"
  type    = "CNAME"
  content = "osrs-amrik-co-uk.pages.dev"
}
