### Catch-all

resource "cloudflare_email_routing_address" "catchall" {
  account_id = data.sops_file.cloudflare-secret.data["cloudflare.account_id"]
  email      = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.catchall_email"]
  lifecycle {
    ignore_changes = [
      account_id
    ]
  }
}

resource "cloudflare_email_routing_catch_all" "catchall" {
  zone_id = data.sops_file.cloudflare-secret.data["cloudflare.amrik.co.uk.zone_id"]
  name    = "Catchall - Managed By Terraform"
  enabled = true

  matcher {
    type = "all"
  }

  action {
    type  = "forward"
    value = [cloudflare_email_routing_address.catchall.email]
  }

}
