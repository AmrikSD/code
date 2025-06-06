variable "ip_address" {
  description = "The IP address to use for the Cloudflare DNS record"
  type        = string
}

variable "tunnel_secret" {
  description = "The secret used to seed the cloudflare tunnel"
  type        = string
}

resource "cloudflare_record" "amrik_xyz" {
  zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.xyz.zone_id"]
  name    = "@"
  type    = "A"
  content = var.ip_address
}

resource "cloudflare_record" "test_amrik_xyz" {
  zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.xyz.zone_id"]
  name    = "test"
  type    = "A"
  content = var.ip_address
}

resource "cloudflare_zero_trust_tunnel_cloudflared" "gcp_tunnel" {
  account_id = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  name       = "Frappe tunnel - terraform"
  secret     = var.tunnel_secret
}

resource "cloudflare_record" "frappe_app" {
  zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.xyz.zone_id"]
  name    = "new"
  content = cloudflare_zero_trust_tunnel_cloudflared.gcp_tunnel.cname
  type    = "CNAME"
  proxied = true
}

resource "cloudflare_zero_trust_tunnel_cloudflared_config" "gcp_tunnel_config" {
  account_id = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  tunnel_id  = cloudflare_zero_trust_tunnel_cloudflared.gcp_tunnel.id
  config {
    ingress_rule {
      hostname = cloudflare_record.frappe_app.hostname
      service  = "http://frontend:8080"
    }
    ingress_rule {
      service = "http_status:404"
    }
  }
}


output "tunnel_id" {
  value = cloudflare_zero_trust_tunnel_cloudflared.gcp_tunnel.id
}
