resource "cloudflare_pages_project" "amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name              = "amrik-co-uk"
  production_branch = "main"

  build_config {
    build_command   = "bundle exec jekyll build"
    root_dir = "projects/amrik.co.uk"
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
    name = "@"
    type = "CNAME"
    value = "amrik-co-uk.pages.dev"
}

//Kate
resource "cloudflare_pages_project" "kate_amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name              = "kate-amrik-co-uk"
  production_branch = "main"
}

resource "cloudflare_pages_domain" "kate_amrik_co_uk" {
  account_id   = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  project_name = "kate-amrik-co-uk"
  domain       = "kate.amrik.co.uk"
}

resource "cloudflare_record" "kate_amrik_co_uk" {
    zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.zone_id"]
    name = "kate"
    type = "CNAME"
    value = "kate-amrik-co-uk.pages.dev"
}

//APC/ Project steel
resource "cloudflare_pages_project" "steel_amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name              = "steel-amrik-co-uk"
  production_branch = "main"
}

resource "cloudflare_pages_domain" "steel_amrik_co_uk" {
  account_id   = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  project_name = "steel-amrik-co-uk"
  domain       = "steel.amrik.co.uk"
}

resource "cloudflare_record" "steel_amrik_co_uk" {
    zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.zone_id"]
    name = "steel"
    type = "CNAME"
    value = "steel-amrik-co-uk.pages.dev"
}


