resource "cloudflare_tunnel" "amrik_co_uk" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name       = "amrik-co-uk"
  secret     = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.tunnel.secret"]
}

resource "cloudflare_access_identity_provider" "github_oauth" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name       = "GitHub OAuth"
  type       = "github"
  config {
    client_id = data.sops_file.cloudflare-secret.data["github.client_id"]
    client_secret = data.sops_file.cloudflare-secret.data["github.client_secret"]
  }
}

resource "cloudflare_tunnel_config" "grafana" {
  account_id        = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  tunnel_id  = cloudflare_tunnel.amrik_co_uk.id

  config {
    ingress_rule {
      hostname = "grafana.amrik.co.uk"
      service  = "http://grafana.monitoring.svc.cluster.local"
    }
    ingress_rule {
      service  = "https://10.0.0.1:8001"
      origin_request {
        access {
          required = true
          team_name = "impossible"
          aud_tag = ["IMPOSSIBLE"]
        }
      }
    }
  }
}
