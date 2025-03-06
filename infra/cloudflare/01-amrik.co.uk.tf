resource "cloudflare_pages_project" "amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name              = "amrik-co-uk"
  production_branch = "main"

  build_config {
    build_command   = "bundle exec jekyll build"
    root_dir        = "projects/amrik.co.uk"
    destination_dir = "_site"
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

resource "cloudflare_pages_domain" "amrik_co_uk" {
  account_id   = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  project_name = "amrik-co-uk"
  domain       = "amrik.co.uk"
}

resource "cloudflare_record" "amrik_co_uk" {
  zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.zone_id"]
  name    = "@"
  type    = "CNAME"
  content = "amrik-co-uk.pages.dev"
}
